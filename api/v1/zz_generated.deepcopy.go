//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 GuardLlama.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdGuardHomeBlockList) DeepCopyInto(out *AdGuardHomeBlockList) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdGuardHomeBlockList.
func (in *AdGuardHomeBlockList) DeepCopy() *AdGuardHomeBlockList {
	if in == nil {
		return nil
	}
	out := new(AdGuardHomeBlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdGuardHomeSpec) DeepCopyInto(out *AdGuardHomeSpec) {
	*out = *in
	if in.FilteringEnabled != nil {
		in, out := &in.FilteringEnabled, &out.FilteringEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BlockLists != nil {
		in, out := &in.BlockLists, &out.BlockLists
		*out = make([]AdGuardHomeBlockList, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdGuardHomeSpec.
func (in *AdGuardHomeSpec) DeepCopy() *AdGuardHomeSpec {
	if in == nil {
		return nil
	}
	out := new(AdGuardHomeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.LastTransitionTime != nil {
		in, out := &in.LastTransitionTime, &out.LastTransitionTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionSet) DeepCopyInto(out *ConditionSet) {
	*out = *in
	if in.deps != nil {
		in, out := &in.deps, &out.deps
		*out = make([]ConditionType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionSet.
func (in *ConditionSet) DeepCopy() *ConditionSet {
	if in == nil {
		return nil
	}
	out := new(ConditionSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tunnel) DeepCopyInto(out *Tunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tunnel.
func (in *Tunnel) DeepCopy() *Tunnel {
	if in == nil {
		return nil
	}
	out := new(Tunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelDNS) DeepCopyInto(out *TunnelDNS) {
	*out = *in
	if in.AdGuardHome != nil {
		in, out := &in.AdGuardHome, &out.AdGuardHome
		*out = new(AdGuardHomeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelDNS.
func (in *TunnelDNS) DeepCopy() *TunnelDNS {
	if in == nil {
		return nil
	}
	out := new(TunnelDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelList) DeepCopyInto(out *TunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelList.
func (in *TunnelList) DeepCopy() *TunnelList {
	if in == nil {
		return nil
	}
	out := new(TunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelProtocol) DeepCopyInto(out *TunnelProtocol) {
	*out = *in
	if in.WireGuard != nil {
		in, out := &in.WireGuard, &out.WireGuard
		*out = new(WireGuardSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelProtocol.
func (in *TunnelProtocol) DeepCopy() *TunnelProtocol {
	if in == nil {
		return nil
	}
	out := new(TunnelProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelSpec) DeepCopyInto(out *TunnelSpec) {
	*out = *in
	in.Protocol.DeepCopyInto(&out.Protocol)
	in.DNS.DeepCopyInto(&out.DNS)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelSpec.
func (in *TunnelSpec) DeepCopy() *TunnelSpec {
	if in == nil {
		return nil
	}
	out := new(TunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TunnelStatus) DeepCopyInto(out *TunnelStatus) {
	*out = *in
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TunnelStatus.
func (in *TunnelStatus) DeepCopy() *TunnelStatus {
	if in == nil {
		return nil
	}
	out := new(TunnelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFromSource) DeepCopyInto(out *ValueFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFromSource.
func (in *ValueFromSource) DeepCopy() *ValueFromSource {
	if in == nil {
		return nil
	}
	out := new(ValueFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueOrValueFrom) DeepCopyInto(out *ValueOrValueFrom) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(ValueFromSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueOrValueFrom.
func (in *ValueOrValueFrom) DeepCopy() *ValueOrValueFrom {
	if in == nil {
		return nil
	}
	out := new(ValueOrValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireGuardInterface) DeepCopyInto(out *WireGuardInterface) {
	*out = *in
	in.PrivateKey.DeepCopyInto(&out.PrivateKey)
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireGuardInterface.
func (in *WireGuardInterface) DeepCopy() *WireGuardInterface {
	if in == nil {
		return nil
	}
	out := new(WireGuardInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireGuardPeer) DeepCopyInto(out *WireGuardPeer) {
	*out = *in
	in.PublicKey.DeepCopyInto(&out.PublicKey)
	in.PresharedKey.DeepCopyInto(&out.PresharedKey)
	if in.AllowedIPs != nil {
		in, out := &in.AllowedIPs, &out.AllowedIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireGuardPeer.
func (in *WireGuardPeer) DeepCopy() *WireGuardPeer {
	if in == nil {
		return nil
	}
	out := new(WireGuardPeer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireGuardSpec) DeepCopyInto(out *WireGuardSpec) {
	*out = *in
	in.Interface.DeepCopyInto(&out.Interface)
	if in.Peers != nil {
		in, out := &in.Peers, &out.Peers
		*out = make([]WireGuardPeer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireGuardSpec.
func (in *WireGuardSpec) DeepCopy() *WireGuardSpec {
	if in == nil {
		return nil
	}
	out := new(WireGuardSpec)
	in.DeepCopyInto(out)
	return out
}
