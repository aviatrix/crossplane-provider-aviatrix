//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPolicyObservation) DeepCopyInto(out *ConnectionPolicyObservation) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DnatIps != nil {
		in, out := &in.DnatIps, &out.DnatIps
		*out = new(string)
		**out = **in
	}
	if in.DnatPort != nil {
		in, out := &in.DnatPort, &out.DnatPort
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPolicyObservation.
func (in *ConnectionPolicyObservation) DeepCopy() *ConnectionPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ConnectionPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPolicyParameters) DeepCopyInto(out *ConnectionPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPolicyParameters.
func (in *ConnectionPolicyParameters) DeepCopy() *ConnectionPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectionPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dnat) DeepCopyInto(out *Dnat) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dnat.
func (in *Dnat) DeepCopy() *Dnat {
	if in == nil {
		return nil
	}
	out := new(Dnat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dnat) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatList) DeepCopyInto(out *DnatList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dnat, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatList.
func (in *DnatList) DeepCopy() *DnatList {
	if in == nil {
		return nil
	}
	out := new(DnatList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DnatList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatObservation) DeepCopyInto(out *DnatObservation) {
	*out = *in
	if in.ConnectionPolicy != nil {
		in, out := &in.ConnectionPolicy, &out.ConnectionPolicy
		*out = make([]ConnectionPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InterfacePolicy != nil {
		in, out := &in.InterfacePolicy, &out.InterfacePolicy
		*out = make([]InterfacePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatObservation.
func (in *DnatObservation) DeepCopy() *DnatObservation {
	if in == nil {
		return nil
	}
	out := new(DnatObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatParameters) DeepCopyInto(out *DnatParameters) {
	*out = *in
	if in.DnatPolicy != nil {
		in, out := &in.DnatPolicy, &out.DnatPolicy
		*out = make([]DnatPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatParameters.
func (in *DnatParameters) DeepCopy() *DnatParameters {
	if in == nil {
		return nil
	}
	out := new(DnatParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatPolicyObservation) DeepCopyInto(out *DnatPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatPolicyObservation.
func (in *DnatPolicyObservation) DeepCopy() *DnatPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(DnatPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatPolicyParameters) DeepCopyInto(out *DnatPolicyParameters) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DnatIps != nil {
		in, out := &in.DnatIps, &out.DnatIps
		*out = new(string)
		**out = **in
	}
	if in.DnatPort != nil {
		in, out := &in.DnatPort, &out.DnatPort
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatPolicyParameters.
func (in *DnatPolicyParameters) DeepCopy() *DnatPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(DnatPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatSpec) DeepCopyInto(out *DnatSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatSpec.
func (in *DnatSpec) DeepCopy() *DnatSpec {
	if in == nil {
		return nil
	}
	out := new(DnatSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnatStatus) DeepCopyInto(out *DnatStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnatStatus.
func (in *DnatStatus) DeepCopy() *DnatStatus {
	if in == nil {
		return nil
	}
	out := new(DnatStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacePolicyObservation) DeepCopyInto(out *InterfacePolicyObservation) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DnatIps != nil {
		in, out := &in.DnatIps, &out.DnatIps
		*out = new(string)
		**out = **in
	}
	if in.DnatPort != nil {
		in, out := &in.DnatPort, &out.DnatPort
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacePolicyObservation.
func (in *InterfacePolicyObservation) DeepCopy() *InterfacePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(InterfacePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfacePolicyParameters) DeepCopyInto(out *InterfacePolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfacePolicyParameters.
func (in *InterfacePolicyParameters) DeepCopy() *InterfacePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(InterfacePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Snat) DeepCopyInto(out *Snat) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Snat.
func (in *Snat) DeepCopy() *Snat {
	if in == nil {
		return nil
	}
	out := new(Snat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Snat) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatConnectionPolicyObservation) DeepCopyInto(out *SnatConnectionPolicyObservation) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SnatIps != nil {
		in, out := &in.SnatIps, &out.SnatIps
		*out = new(string)
		**out = **in
	}
	if in.SnatPort != nil {
		in, out := &in.SnatPort, &out.SnatPort
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatConnectionPolicyObservation.
func (in *SnatConnectionPolicyObservation) DeepCopy() *SnatConnectionPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SnatConnectionPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatConnectionPolicyParameters) DeepCopyInto(out *SnatConnectionPolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatConnectionPolicyParameters.
func (in *SnatConnectionPolicyParameters) DeepCopy() *SnatConnectionPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SnatConnectionPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatInterfacePolicyObservation) DeepCopyInto(out *SnatInterfacePolicyObservation) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SnatIps != nil {
		in, out := &in.SnatIps, &out.SnatIps
		*out = new(string)
		**out = **in
	}
	if in.SnatPort != nil {
		in, out := &in.SnatPort, &out.SnatPort
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatInterfacePolicyObservation.
func (in *SnatInterfacePolicyObservation) DeepCopy() *SnatInterfacePolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SnatInterfacePolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatInterfacePolicyParameters) DeepCopyInto(out *SnatInterfacePolicyParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatInterfacePolicyParameters.
func (in *SnatInterfacePolicyParameters) DeepCopy() *SnatInterfacePolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SnatInterfacePolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatList) DeepCopyInto(out *SnatList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Snat, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatList.
func (in *SnatList) DeepCopy() *SnatList {
	if in == nil {
		return nil
	}
	out := new(SnatList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SnatList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatObservation) DeepCopyInto(out *SnatObservation) {
	*out = *in
	if in.ConnectionPolicy != nil {
		in, out := &in.ConnectionPolicy, &out.ConnectionPolicy
		*out = make([]SnatConnectionPolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InterfacePolicy != nil {
		in, out := &in.InterfacePolicy, &out.InterfacePolicy
		*out = make([]SnatInterfacePolicyObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatObservation.
func (in *SnatObservation) DeepCopy() *SnatObservation {
	if in == nil {
		return nil
	}
	out := new(SnatObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatParameters) DeepCopyInto(out *SnatParameters) {
	*out = *in
	if in.SnatMode != nil {
		in, out := &in.SnatMode, &out.SnatMode
		*out = new(string)
		**out = **in
	}
	if in.SnatPolicy != nil {
		in, out := &in.SnatPolicy, &out.SnatPolicy
		*out = make([]SnatPolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SyncToHa != nil {
		in, out := &in.SyncToHa, &out.SyncToHa
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatParameters.
func (in *SnatParameters) DeepCopy() *SnatParameters {
	if in == nil {
		return nil
	}
	out := new(SnatParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicyObservation) DeepCopyInto(out *SnatPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicyObservation.
func (in *SnatPolicyObservation) DeepCopy() *SnatPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(SnatPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatPolicyParameters) DeepCopyInto(out *SnatPolicyParameters) {
	*out = *in
	if in.ApplyRouteEntry != nil {
		in, out := &in.ApplyRouteEntry, &out.ApplyRouteEntry
		*out = new(bool)
		**out = **in
	}
	if in.Connection != nil {
		in, out := &in.Connection, &out.Connection
		*out = new(string)
		**out = **in
	}
	if in.DstCidr != nil {
		in, out := &in.DstCidr, &out.DstCidr
		*out = new(string)
		**out = **in
	}
	if in.DstPort != nil {
		in, out := &in.DstPort, &out.DstPort
		*out = new(string)
		**out = **in
	}
	if in.ExcludeRtb != nil {
		in, out := &in.ExcludeRtb, &out.ExcludeRtb
		*out = new(string)
		**out = **in
	}
	if in.Interface != nil {
		in, out := &in.Interface, &out.Interface
		*out = new(string)
		**out = **in
	}
	if in.Mark != nil {
		in, out := &in.Mark, &out.Mark
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.SnatIps != nil {
		in, out := &in.SnatIps, &out.SnatIps
		*out = new(string)
		**out = **in
	}
	if in.SnatPort != nil {
		in, out := &in.SnatPort, &out.SnatPort
		*out = new(string)
		**out = **in
	}
	if in.SrcCidr != nil {
		in, out := &in.SrcCidr, &out.SrcCidr
		*out = new(string)
		**out = **in
	}
	if in.SrcPort != nil {
		in, out := &in.SrcPort, &out.SrcPort
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatPolicyParameters.
func (in *SnatPolicyParameters) DeepCopy() *SnatPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(SnatPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatSpec) DeepCopyInto(out *SnatSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatSpec.
func (in *SnatSpec) DeepCopy() *SnatSpec {
	if in == nil {
		return nil
	}
	out := new(SnatSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatStatus) DeepCopyInto(out *SnatStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatStatus.
func (in *SnatStatus) DeepCopy() *SnatStatus {
	if in == nil {
		return nil
	}
	out := new(SnatStatus)
	in.DeepCopyInto(out)
	return out
}