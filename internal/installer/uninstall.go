package installer

import (
	"context"

	"github.com/SilvanAbuse/guardllama/internal/build"
	"github.com/SilvanAbuse/guardllama/internal/installer/k3d"
	"github.com/SilvanAbuse/guardllama/internal/installer/k3s"
	"github.com/SilvanAbuse/guardllama/internal/log"
	apiv1 "github.com/SilvanAbuse/guardllama/proto/gen/api/v1"
)

func Uninstall(ctx context.Context, cfg *apiv1.ServerConfig) error {
	logger := log.FromContext(ctx)

	cfg, err := fillDefaultServerConfig(cfg)
	if err != nil {
		return err
	}
	if err := cfg.ValidateAll(); err != nil {
		return err
	}

	if k3dCluster := cfg.Cluster.GetK3D(); k3dCluster != nil {
		k := k3d.NewWithOpts(
			k3d.WithName(k3dCluster.Name),
			k3d.WithKubeConfig(cfg.Cluster.KubeConfig),
			k3d.WithK3sVersion(build.K3sVersion),
			k3d.WithLogger(logger),
		)

		return k.Destroy(ctx)
	}

	k := k3s.NewWithOpts(
		k3s.WithKubeConfig(cfg.Cluster.KubeConfig),
		k3s.WithK3sVersion(build.K3sVersion),
		k3s.WithLogger(logger),
	)

	return k.Destroy(ctx)
}
