package controller

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/yaml"

	"github.com/go-logr/logr"
	glmv1 "github.com/guardllamanet/guardllama/api/v1"
	apiv1 "github.com/guardllamanet/guardllama/proto/gen/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	reasonConfigReady  = "ConfigReady"
	reasonConfigError  = "ConfigError"
	reasonDeployReady  = "DeployReady"
	reasonDeployError  = "DeployError"
	reasonServiceReady = "ServiceReady"
	reasonServiceError = "ServiceError"
	reasonPodRunning   = "PodRunning"
	reasonPodPending   = "PodPending"

	tunnelHTTPAddr int32 = 8080

	labelTunnelName               = SystemLabelPrefix + "/tunnel-name"
	annotationTunnelConfigVersion = SystemLabelPrefix + "/tunnel-config-version"
	annotationDNSConfigVersion    = SystemLabelPrefix + "/dns-config-version"
)

type TunnelReconciler struct {
	client.Client
	logr.Logger

	TunnelImage           string
	TunnelImagePullPolicy string
	TunnelImagePullSecret string
}

// +kubebuilder:rbac:groups=tunnel.guardllama.net,resources=tunnels,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=tunnel.guardllama.net,resources=tunnels/status,verbs=get;update;patch

// +kubebuilder:rbac:groups="coordination.k8s.io",resources=leases,verbs=get;list;create;update

// +kubebuilder:rbac:groups="",resources=namespaces,verbs=get;list;watch;update;patch;create;delete
// +kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch;update;patch;create
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;update;patch;create
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete

func (r *TunnelReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	ll := r.Logger.WithValues("tunnel", req.NamespacedName)

	tun := &glmv1.Tunnel{}
	if err := r.Get(ctx, req.NamespacedName, tun); err != nil {
		ll.Error(err, "error getting tunnel")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if IsDeleted(tun) {
		return ctrl.Result{}, nil
	}

	steps := []Step{
		{
			Name: "ReconcileTunnel",
			Run: func() (controllerutil.OperationResult, error) {
				return r.reconcileTunnel(ctx, tun)
			},
		},
	}

	return ctrl.Result{}, RunSteps(ll, steps...)
}

func (r *TunnelReconciler) refreshTunnelStatus(ctx context.Context, tun *glmv1.Tunnel) (controllerutil.OperationResult, error) {
	existing := &glmv1.Tunnel{}
	if err := r.Get(ctx, client.ObjectKeyFromObject(tun), existing); err != nil {
		return controllerutil.OperationResultNone, fmt.Errorf("error getting tunnel: %w", err)
	}

	if equality.Semantic.DeepEqual(existing.Status, tun.Status) {
		return controllerutil.OperationResultNone, nil
	}

	now := metav1.Now()
	tun.Status.UpdatedAt = now.DeepCopy()

	if err := r.Status().Update(ctx, tun); err != nil {
		return controllerutil.OperationResultNone, err
	}

	return controllerutil.OperationResultUpdatedStatusOnly, nil
}

func (r *TunnelReconciler) reconcileTunnel(ctx context.Context, tun *glmv1.Tunnel) (controllerutil.OperationResult, error) {
	var returnErr error

	tunCfg, err := r.upsertTunnelConfig(ctx, tun)
	if err != nil {
		returnErr = errors.Join(returnErr, fmt.Errorf("error upserting tunnel secret: %w", err))
	}

	if tunCfg != nil {
		if _, err := r.upsertTunnelDeploy(ctx, tun, tunCfg); err != nil {
			returnErr = errors.Join(returnErr, fmt.Errorf("error upserting tunnel deploy: %w", err))
		}

		if _, err := r.upsertTunnelService(ctx, tun); err != nil {
			returnErr = errors.Join(returnErr, fmt.Errorf("error upserting tunnel service: %w", err))
		}
	}

	dnsCfg, err := r.upsertDNSServerConfig(ctx, tun)
	if err != nil {
		returnErr = errors.Join(returnErr, fmt.Errorf("error upserting dns server config: %w", err))
	}

	if dnsCfg != nil {
		if _, err := r.upsertDNSServerDeploy(ctx, tun, dnsCfg); err != nil {
			returnErr = errors.Join(returnErr, fmt.Errorf("error upserting dns server deploy: %w", err))
		}

		if _, err := r.upsertDNSServerService(ctx, tun); err != nil {
			returnErr = errors.Join(returnErr, fmt.Errorf("error upserting dns server service: %w", err))
		}
	}

	if _, err := r.refreshTunnelStatus(ctx, tun); err != nil {
		returnErr = errors.Join(returnErr, fmt.Errorf("error refreshing tunnel status: %w", err))
	}

	if returnErr != nil {
		return controllerutil.OperationResultNone, returnErr
	}

	return controllerutil.OperationResultUpdated, nil
}

func (r *TunnelReconciler) upsertTunnelConfig(ctx context.Context, tun *glmv1.Tunnel) (*corev1.Secret, error) {
	wgc, err := makeWireGuardConfig(ctx, r, tun)
	if err != nil {
		return nil, err
	}

	data, err := yaml.Marshal(wgc)
	if err != nil {
		return nil, err
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.WireGuardTypedName(),
			Namespace: tun.Namespace,
		},
	}
	secretFn := func() error {
		secret.Data = map[string][]byte{"wg.yaml": data}

		return controllerutil.SetControllerReference(tun, secret, r.Scheme())
	}

	_, err = controllerutil.CreateOrUpdate(ctx, r, secret, secretFn)
	if err != nil {
		tun.Status.SetCondition(
			glmv1.ConditionTunnelConfigReady,
			corev1.ConditionFalse,
			reasonConfigError,
			err.Error(),
		)

		return nil, err
	}

	tun.Status.SetCondition(
		glmv1.ConditionTunnelConfigReady,
		corev1.ConditionTrue,
		reasonConfigReady,
		"",
	)

	return secret, nil
}

func (r *TunnelReconciler) upsertTunnelDeploy(ctx context.Context, tun *glmv1.Tunnel, cfg *corev1.Secret) (controllerutil.OperationResult, error) {
	var (
		labels = labelsTunnelDeploy(tun)
	)

	var secret *corev1.Secret
	if ips := r.TunnelImagePullSecret; ips != "" {
		secret = &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-imagepullsecret", tun.WireGuardTypedName()),
				Namespace: tun.Namespace,
			},
		}
		secretFn := func() error {
			secret.Type = corev1.SecretTypeDockerConfigJson
			secret.Data = map[string][]byte{
				corev1.DockerConfigJsonKey: []byte(ips),
			}

			return controllerutil.SetControllerReference(tun, secret, r.Scheme())
		}

		if _, err := controllerutil.CreateOrUpdate(ctx, r, secret, secretFn); err != nil {
			return controllerutil.OperationResultNone, fmt.Errorf("Error upserting tunnel image pull secret: %w", err)
		}
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.WireGuardTypedName(),
			Namespace: tun.Namespace,
		},
	}
	deployFn := func() error {
		deploy.Spec = appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: annosTunnelWithAssociatedObject(tun, cfg, annotationTunnelConfigVersion),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "tunnel",
							Image:           r.TunnelImage,
							ImagePullPolicy: corev1.PullPolicy(r.TunnelImagePullPolicy),
							SecurityContext: &corev1.SecurityContext{
								Capabilities: &corev1.Capabilities{
									Add: []corev1.Capability{
										"NET_ADMIN",
										"SYS_MODULE",
									},
								},
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: tunnelHTTPAddr,
									Protocol:      corev1.ProtocolTCP,
								},
								{
									Name:          "wg",
									ContainerPort: tun.Spec.Protocol.WireGuard.Interface.ListenPort,
									Protocol:      corev1.ProtocolUDP,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "LOG_LEVEL",
									Value: "debug",
								},
								{
									Name:  "GLMMGR_HTTP_ADDR",
									Value: fmt.Sprintf(":%d", tunnelHTTPAddr),
								},
								{
									Name:  "GLMMGR_TUNNEL_CONFIG",
									Value: "/etc/tunnel/wg.yaml",
								},
								{
									Name:  "WG_QUICK_USERSPACE_IMPLEMENTATION",
									Value: tun.Spec.Protocol.WireGuard.UserspaceImpl,
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/health",
										Port: intstr.FromString("http"),
									},
								},
							},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/health",
										Port: intstr.FromString("http"),
									},
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "tunnel-secret",
									MountPath: "/etc/tunnel",
									ReadOnly:  true,
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "tunnel-secret",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName: tun.WireGuardTypedName(), // secret is tunnel typed name
								},
							},
						},
					},
				},
			},
		}

		if secret != nil {
			deploy.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{
				{
					Name: secret.Name,
				},
			}
		}

		// Update tunnel pod status
		setDeployStatus(tun, glmv1.ConditionTunnelPodReady, deploy.Status)

		return controllerutil.SetControllerReference(tun, deploy, r.Scheme())
	}

	op, err := controllerutil.CreateOrUpdate(ctx, r, deploy, deployFn)
	if err == nil {
		tun.Status.SetCondition(
			glmv1.ConditionTunnelDeployReady,
			corev1.ConditionTrue,
			reasonDeployReady,
			"",
		)
	} else {
		tun.Status.SetCondition(
			glmv1.ConditionTunnelDeployReady,
			corev1.ConditionFalse,
			reasonDeployError,
			err.Error(),
		)
	}

	return op, err
}

func (r *TunnelReconciler) upsertTunnelService(ctx context.Context, tun *glmv1.Tunnel) (controllerutil.OperationResult, error) {
	wgSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.WireGuardServiceName(),
			Namespace: tun.Namespace,
		},
	}
	wgSvcFn := func() error {
		// find and update tunnel ingress port
		var nodePort int32
		for _, port := range wgSvc.Spec.Ports {
			if port.Name == "wg" {
				nodePort = port.NodePort
			}
		}
		tun.Status.IngressPort = nodePort

		wgSvc.Spec.Type = corev1.ServiceTypeNodePort
		wgSvc.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "wg",
				Protocol:   corev1.ProtocolUDP,
				Port:       tun.Spec.Protocol.WireGuard.Interface.ListenPort,
				TargetPort: intstr.FromInt(int(tun.Spec.Protocol.WireGuard.Interface.ListenPort)),
				NodePort:   tun.Spec.Protocol.WireGuard.Interface.IngressPort,
			},
		}
		wgSvc.Spec.Selector = labelsTunnelDeploy(tun)

		return controllerutil.SetControllerReference(tun, wgSvc, r.Scheme())
	}

	op, err := controllerutil.CreateOrUpdate(ctx, r, wgSvc, wgSvcFn)
	if err != nil {
		tun.Status.SetCondition(
			glmv1.ConditionTunnelServiceReady,
			corev1.ConditionFalse,
			reasonServiceError,
			err.Error(),
		)
		return op, err
	}

	adminSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.WireGuardAdminServiceName(),
			Namespace: tun.Namespace,
		},
	}
	adminSvcFn := func() error {
		adminSvc.Spec.Type = corev1.ServiceTypeClusterIP
		adminSvc.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "http",
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromString("http"),
			},
		}
		adminSvc.Spec.Selector = labelsTunnelDeploy(tun)

		return controllerutil.SetControllerReference(tun, adminSvc, r.Scheme())
	}
	op, err = controllerutil.CreateOrUpdate(ctx, r, adminSvc, adminSvcFn)
	if err != nil {
		tun.Status.SetCondition(
			glmv1.ConditionTunnelServiceReady,
			corev1.ConditionFalse,
			reasonServiceError,
			err.Error(),
		)
		return op, err
	}

	tun.Status.SetCondition(
		glmv1.ConditionTunnelServiceReady,
		corev1.ConditionTrue,
		reasonServiceReady,
		"",
	)

	return controllerutil.OperationResultUpdated, nil
}

func (r *TunnelReconciler) upsertDNSServerConfig(ctx context.Context, tun *glmv1.Tunnel) (*corev1.ConfigMap, error) {
	cfg, err := adGuardHomeConfig(tun.Spec.DNS.AdGuard)
	if err != nil {
		return nil, err
	}

	cm := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.AdGuardTypedName(),
			Namespace: tun.Namespace,
		},
	}
	cmFn := func() error {
		cm.Data = map[string]string{"AdGuardHome.yaml": cfg}

		return controllerutil.SetControllerReference(tun, cm, r.Scheme())
	}

	_, err = controllerutil.CreateOrUpdate(ctx, r, cm, cmFn)
	if err != nil {
		tun.Status.SetCondition(
			glmv1.ConditionDNSConfigReady,
			corev1.ConditionFalse,
			reasonConfigError,
			err.Error(),
		)

		return nil, err
	}

	tun.Status.SetCondition(
		glmv1.ConditionDNSConfigReady,
		corev1.ConditionTrue,
		reasonConfigReady,
		"",
	)

	return cm, nil
}

func (r *TunnelReconciler) upsertDNSServerDeploy(ctx context.Context, tun *glmv1.Tunnel, cfg *corev1.ConfigMap) (controllerutil.OperationResult, error) {
	var (
		labels = labelsDNSDeploy(tun)
	)

	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.AdGuardTypedName(),
			Namespace: tun.Namespace,
		},
	}
	deployFn := func() error {
		deploy.Spec = appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels:      labels,
					Annotations: annosTunnelWithAssociatedObject(tun, cfg, annotationDNSConfigVersion),
				},
				Spec: corev1.PodSpec{
					InitContainers: []corev1.Container{
						{
							Name:  "copy-configmap",
							Image: "busybox",
							Command: []string{
								"sh",
								"-c",
								"mkdir -p /opt/adguardhome/conf && cp /tmp/AdGuardHome.yaml /opt/adguardhome/conf/AdGuardHome.yaml",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "adguardhome-ro-config",
									SubPath:   "AdGuardHome.yaml",
									MountPath: "/tmp/AdGuardHome.yaml",
									ReadOnly:  true,
								},
								{
									Name:      "adguardhome-config",
									MountPath: "/opt/adguardhome/conf",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "adguardhome",
							Image: adGuardHomeImage,
							Args: []string{
								"--config",
								"/opt/adguardhome/conf/AdGuardHome.yaml",
								"--work-dir",
								"/opt/adguardhome/work",
								"--no-check-update",
							},
							Ports: []corev1.ContainerPort{
								{
									Name:          "dns-tcp",
									ContainerPort: 5353,
									Protocol:      corev1.ProtocolTCP,
								},
								{
									Name:          "dns-udp",
									ContainerPort: 5353,
									Protocol:      corev1.ProtocolUDP,
								},
								{
									Name:          "http",
									ContainerPort: 3000,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Env: []corev1.EnvVar{
								{
									Name:  "TZ",
									Value: "UTC",
								},
							},
							LivenessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/",
										Port: intstr.FromString("http"),
									},
								},
							},
							ReadinessProbe: &corev1.Probe{
								ProbeHandler: corev1.ProbeHandler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/",
										Port: intstr.FromString("http"),
									},
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "adguardhome-config",
									MountPath: "/opt/adguardhome/conf",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "adguardhome-ro-config",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: tun.AdGuardTypedName(),
									},
								},
							},
						},
						{
							Name: "adguardhome-config",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{
									Medium: corev1.StorageMediumMemory,
								},
							},
						},
					},
				},
			},
		}

		// Update dns pod status
		setDeployStatus(tun, glmv1.ConditionDNSPodReady, deploy.Status)

		return controllerutil.SetControllerReference(tun, deploy, r.Scheme())
	}

	op, err := controllerutil.CreateOrUpdate(ctx, r, deploy, deployFn)
	if err == nil {
		tun.Status.SetCondition(
			glmv1.ConditionDNSDeployReady,
			corev1.ConditionTrue,
			reasonDeployReady,
			"",
		)
	} else {
		tun.Status.SetCondition(
			glmv1.ConditionDNSDeployReady,
			corev1.ConditionFalse,
			reasonDeployError,
			err.Error(),
		)
	}

	return op, err
}

func (r *TunnelReconciler) upsertDNSServerService(ctx context.Context, tun *glmv1.Tunnel) (controllerutil.OperationResult, error) {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tun.AdGuardAdminServiceName(),
			Namespace: tun.Namespace,
		},
	}
	svcFn := func() error {
		svc.Spec.Type = corev1.ServiceTypeClusterIP
		svc.Spec.Ports = []corev1.ServicePort{
			{
				Name:       "dns-tcp",
				Protocol:   corev1.ProtocolTCP,
				Port:       53,
				TargetPort: intstr.FromString("dns-tcp"),
			},
			{
				Name:       "dns-udp",
				Protocol:   corev1.ProtocolUDP,
				Port:       53,
				TargetPort: intstr.FromString("dns-udp"),
			},
			{
				Name:       "http",
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromString("http"),
			},
		}
		svc.Spec.Selector = labelsDNSDeploy(tun)

		tun.Status.DNS = svc.Spec.ClusterIPs

		return controllerutil.SetControllerReference(tun, svc, r.Scheme())
	}

	op, err := controllerutil.CreateOrUpdate(ctx, r, svc, svcFn)
	if err == nil {
		tun.Status.SetCondition(
			glmv1.ConditionDNSServiceReady,
			corev1.ConditionTrue,
			reasonServiceReady,
			"",
		)
	} else {
		tun.Status.SetCondition(
			glmv1.ConditionDNSServiceReady,
			corev1.ConditionFalse,
			reasonServiceError,
			err.Error(),
		)
	}

	return op, err
}

func (r *TunnelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&glmv1.Tunnel{}).
		Owns(&corev1.Secret{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ConfigMap{}).
		Complete(r)
}

func labelsTunnelDeploy(tun *glmv1.Tunnel) map[string]string {
	return MergeLabels(
		filterSystemLabels(tun.Labels),
		map[string]string{
			labelTunnelName: tun.Name,
		},
	)
}

func labelsDNSDeploy(tun *glmv1.Tunnel) map[string]string {
	return MergeLabels(
		filterSystemLabels(tun.Labels),
		map[string]string{
			labelTunnelName: tun.AdGuardTypedName(),
		},
	)
}

func annosTunnelWithAssociatedObject(tun *glmv1.Tunnel, obj metav1.Object, key string) map[string]string {
	return MergeLabels(
		filterSystemLabels(tun.Annotations),
		map[string]string{
			key: obj.GetResourceVersion(),
		},
	)
}

func setDeployStatus(tun *glmv1.Tunnel, ct glmv1.ConditionType, status appsv1.DeploymentStatus) {
	if status.UnavailableReplicas == 0 {
		tun.Status.SetCondition(
			ct,
			corev1.ConditionTrue,
			reasonPodRunning,
			fmt.Sprintf("%d/%d", status.ReadyReplicas, status.Replicas),
		)
	} else {
		var (
			deployCond   = corev1.ConditionUnknown
			deployReason = reasonPodPending
			deployMsg    = fmt.Sprintf("%d/%d", status.ReadyReplicas, status.Replicas)
		)
		for _, cond := range status.Conditions {
			if cond.Status == corev1.ConditionFalse {
				deployReason = cond.Reason
				deployMsg = cond.Message
				break
			}
		}

		tun.Status.SetCondition(
			ct,
			deployCond,
			deployReason,
			deployMsg,
		)
	}
}

func makeWireGuardConfig(ctx context.Context, c client.Client, tun *glmv1.Tunnel) (*apiv1.WireGuardConfig, error) {
	var peers []*apiv1.WireGuardPeer
	for _, p := range tun.Spec.Protocol.WireGuard.Peers {
		var i *durationpb.Duration
		if p.PersistentKeepaliveIntervalSeconds != 0 {
			d, _ := time.ParseDuration(fmt.Sprintf("%ds", p.PersistentKeepaliveIntervalSeconds))
			i = durationpb.New(d)
		}

		publicKey, err := getValueOrValueFrom(ctx, c, tun.Namespace, p.PublicKey)
		if err != nil {
			return nil, err
		}

		presharedKey, err := getValueOrValueFrom(ctx, c, tun.Namespace, p.PresharedKey)
		if err != nil {
			return nil, err
		}

		peers = append(peers, &apiv1.WireGuardPeer{
			PublicKey:                   publicKey,
			PresharedKey:                presharedKey,
			Endpoint:                    p.Endpoint,
			PersistentKeepaliveInterval: i,
			AllowedIps:                  p.AllowedIPs,
		})
	}

	privateKey, err := getValueOrValueFrom(ctx, c, tun.Namespace, tun.Spec.Protocol.WireGuard.Interface.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &apiv1.WireGuardConfig{
		Interface: &apiv1.WireGuardInterface{
			PrivateKey:   privateKey,
			ListenPort:   tun.Spec.Protocol.WireGuard.Interface.ListenPort,
			IngressPort:  tun.Spec.Protocol.WireGuard.Interface.IngressPort,
			FirewallMark: tun.Spec.Protocol.WireGuard.Interface.FirewallMark,
			Address:      tun.Spec.Protocol.WireGuard.Interface.Address,
			Mtu:          tun.Spec.Protocol.WireGuard.Interface.MTU,
			Table:        tun.Spec.Protocol.WireGuard.Interface.Table,
			PreUp:        tun.Spec.Protocol.WireGuard.Interface.PreUp,
			PostUp:       tun.Spec.Protocol.WireGuard.Interface.PostUp,
			PreDown:      tun.Spec.Protocol.WireGuard.Interface.PreDown,
			PostDown:     tun.Spec.Protocol.WireGuard.Interface.PostDown,
		},
		Peers: peers,
	}, nil

}

func getValueOrValueFrom(ctx context.Context, c client.Reader, ns string, i glmv1.ValueOrValueFrom) (string, error) {
	vf := i.ValueFrom
	if vf == nil {
		return i.Value, nil
	}

	secretKeyRef := vf.SecretKeyRef
	name := secretKeyRef.Name
	key := secretKeyRef.Key
	optional := secretKeyRef.Optional != nil && *secretKeyRef.Optional

	var secret corev1.Secret
	if err := c.Get(ctx,
		types.NamespacedName{
			Name:      name,
			Namespace: ns,
		},
		&secret,
	); err != nil {
		return "", fmt.Errorf("Error getting secret %s/%s: %w", ns, name, err)
	}

	valBytes, ok := secret.Data[key]
	if !ok {
		if optional {
			return "", nil
		}

		return "", fmt.Errorf("Couldn't find key %s in secret %s/%s", key, ns, name)
	}

	return string(valBytes), nil
}
