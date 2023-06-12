//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Gateway) DeepCopyInto(out *Gateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Gateway.
func (in *Gateway) DeepCopy() *Gateway {
	if in == nil {
		return nil
	}
	out := new(Gateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Gateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayList) DeepCopyInto(out *GatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Gateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayList.
func (in *GatewayList) DeepCopy() *GatewayList {
	if in == nil {
		return nil
	}
	out := new(GatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayObservation) DeepCopyInto(out *GatewayObservation) {
	*out = *in
	if in.BGPLanIPList != nil {
		in, out := &in.BGPLanIPList, &out.BGPLanIPList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CloudInstanceID != nil {
		in, out := &in.CloudInstanceID, &out.CloudInstanceID
		*out = new(string)
		**out = **in
	}
	if in.HaBGPLanIPList != nil {
		in, out := &in.HaBGPLanIPList, &out.HaBGPLanIPList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.HaCloudInstanceID != nil {
		in, out := &in.HaCloudInstanceID, &out.HaCloudInstanceID
		*out = new(string)
		**out = **in
	}
	if in.HaGwName != nil {
		in, out := &in.HaGwName, &out.HaGwName
		*out = new(string)
		**out = **in
	}
	if in.HaPrivateIP != nil {
		in, out := &in.HaPrivateIP, &out.HaPrivateIP
		*out = new(string)
		**out = **in
	}
	if in.HaPublicIP != nil {
		in, out := &in.HaPublicIP, &out.HaPublicIP
		*out = new(string)
		**out = **in
	}
	if in.HaSecurityGroupID != nil {
		in, out := &in.HaSecurityGroupID, &out.HaSecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIP != nil {
		in, out := &in.PrivateIP, &out.PrivateIP
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayObservation.
func (in *GatewayObservation) DeepCopy() *GatewayObservation {
	if in == nil {
		return nil
	}
	out := new(GatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayParameters) DeepCopyInto(out *GatewayParameters) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.AllocateNewEIP != nil {
		in, out := &in.AllocateNewEIP, &out.AllocateNewEIP
		*out = new(bool)
		**out = **in
	}
	if in.ApprovedLearnedCidrs != nil {
		in, out := &in.ApprovedLearnedCidrs, &out.ApprovedLearnedCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AvailabilityDomain != nil {
		in, out := &in.AvailabilityDomain, &out.AvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.AzureEIPNameResourceGroup != nil {
		in, out := &in.AzureEIPNameResourceGroup, &out.AzureEIPNameResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.BGPEcmp != nil {
		in, out := &in.BGPEcmp, &out.BGPEcmp
		*out = new(bool)
		**out = **in
	}
	if in.BGPHoldTime != nil {
		in, out := &in.BGPHoldTime, &out.BGPHoldTime
		*out = new(float64)
		**out = **in
	}
	if in.BGPLanInterfacesCount != nil {
		in, out := &in.BGPLanInterfacesCount, &out.BGPLanInterfacesCount
		*out = new(float64)
		**out = **in
	}
	if in.BGPPollingTime != nil {
		in, out := &in.BGPPollingTime, &out.BGPPollingTime
		*out = new(float64)
		**out = **in
	}
	if in.CloudType != nil {
		in, out := &in.CloudType, &out.CloudType
		*out = new(float64)
		**out = **in
	}
	if in.CustomerManagedKeysSecretRef != nil {
		in, out := &in.CustomerManagedKeysSecretRef, &out.CustomerManagedKeysSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.CustomizedSpokeVPCRoutes != nil {
		in, out := &in.CustomizedSpokeVPCRoutes, &out.CustomizedSpokeVPCRoutes
		*out = new(string)
		**out = **in
	}
	if in.DisableRoutePropagation != nil {
		in, out := &in.DisableRoutePropagation, &out.DisableRoutePropagation
		*out = new(bool)
		**out = **in
	}
	if in.EIP != nil {
		in, out := &in.EIP, &out.EIP
		*out = new(string)
		**out = **in
	}
	if in.EnableActiveStandby != nil {
		in, out := &in.EnableActiveStandby, &out.EnableActiveStandby
		*out = new(bool)
		**out = **in
	}
	if in.EnableActiveStandbyPreemptive != nil {
		in, out := &in.EnableActiveStandbyPreemptive, &out.EnableActiveStandbyPreemptive
		*out = new(bool)
		**out = **in
	}
	if in.EnableAutoAdvertiseS2CCidrs != nil {
		in, out := &in.EnableAutoAdvertiseS2CCidrs, &out.EnableAutoAdvertiseS2CCidrs
		*out = new(bool)
		**out = **in
	}
	if in.EnableBGP != nil {
		in, out := &in.EnableBGP, &out.EnableBGP
		*out = new(bool)
		**out = **in
	}
	if in.EnableBGPOverLan != nil {
		in, out := &in.EnableBGPOverLan, &out.EnableBGPOverLan
		*out = new(bool)
		**out = **in
	}
	if in.EnableEncryptVolume != nil {
		in, out := &in.EnableEncryptVolume, &out.EnableEncryptVolume
		*out = new(bool)
		**out = **in
	}
	if in.EnableGlobalVPC != nil {
		in, out := &in.EnableGlobalVPC, &out.EnableGlobalVPC
		*out = new(bool)
		**out = **in
	}
	if in.EnableGroGso != nil {
		in, out := &in.EnableGroGso, &out.EnableGroGso
		*out = new(bool)
		**out = **in
	}
	if in.EnableJumboFrame != nil {
		in, out := &in.EnableJumboFrame, &out.EnableJumboFrame
		*out = new(bool)
		**out = **in
	}
	if in.EnableLearnedCidrsApproval != nil {
		in, out := &in.EnableLearnedCidrsApproval, &out.EnableLearnedCidrsApproval
		*out = new(bool)
		**out = **in
	}
	if in.EnableMonitorGatewaySubnets != nil {
		in, out := &in.EnableMonitorGatewaySubnets, &out.EnableMonitorGatewaySubnets
		*out = new(bool)
		**out = **in
	}
	if in.EnablePreserveAsPath != nil {
		in, out := &in.EnablePreserveAsPath, &out.EnablePreserveAsPath
		*out = new(bool)
		**out = **in
	}
	if in.EnablePrivateOob != nil {
		in, out := &in.EnablePrivateOob, &out.EnablePrivateOob
		*out = new(bool)
		**out = **in
	}
	if in.EnablePrivateVPCDefaultRoute != nil {
		in, out := &in.EnablePrivateVPCDefaultRoute, &out.EnablePrivateVPCDefaultRoute
		*out = new(bool)
		**out = **in
	}
	if in.EnableSkipPublicRouteTableUpdate != nil {
		in, out := &in.EnableSkipPublicRouteTableUpdate, &out.EnableSkipPublicRouteTableUpdate
		*out = new(bool)
		**out = **in
	}
	if in.EnableSpotInstance != nil {
		in, out := &in.EnableSpotInstance, &out.EnableSpotInstance
		*out = new(bool)
		**out = **in
	}
	if in.EnableVPCDNSServer != nil {
		in, out := &in.EnableVPCDNSServer, &out.EnableVPCDNSServer
		*out = new(bool)
		**out = **in
	}
	if in.FaultDomain != nil {
		in, out := &in.FaultDomain, &out.FaultDomain
		*out = new(string)
		**out = **in
	}
	if in.FilteredSpokeVPCRoutes != nil {
		in, out := &in.FilteredSpokeVPCRoutes, &out.FilteredSpokeVPCRoutes
		*out = new(string)
		**out = **in
	}
	if in.GwSize != nil {
		in, out := &in.GwSize, &out.GwSize
		*out = new(string)
		**out = **in
	}
	if in.HaAvailabilityDomain != nil {
		in, out := &in.HaAvailabilityDomain, &out.HaAvailabilityDomain
		*out = new(string)
		**out = **in
	}
	if in.HaAzureEIPNameResourceGroup != nil {
		in, out := &in.HaAzureEIPNameResourceGroup, &out.HaAzureEIPNameResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.HaEIP != nil {
		in, out := &in.HaEIP, &out.HaEIP
		*out = new(string)
		**out = **in
	}
	if in.HaFaultDomain != nil {
		in, out := &in.HaFaultDomain, &out.HaFaultDomain
		*out = new(string)
		**out = **in
	}
	if in.HaGwSize != nil {
		in, out := &in.HaGwSize, &out.HaGwSize
		*out = new(string)
		**out = **in
	}
	if in.HaImageVersion != nil {
		in, out := &in.HaImageVersion, &out.HaImageVersion
		*out = new(string)
		**out = **in
	}
	if in.HaInsaneModeAz != nil {
		in, out := &in.HaInsaneModeAz, &out.HaInsaneModeAz
		*out = new(string)
		**out = **in
	}
	if in.HaOobAvailabilityZone != nil {
		in, out := &in.HaOobAvailabilityZone, &out.HaOobAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.HaOobManagementSubnet != nil {
		in, out := &in.HaOobManagementSubnet, &out.HaOobManagementSubnet
		*out = new(string)
		**out = **in
	}
	if in.HaPrivateModeSubnetZone != nil {
		in, out := &in.HaPrivateModeSubnetZone, &out.HaPrivateModeSubnetZone
		*out = new(string)
		**out = **in
	}
	if in.HaSoftwareVersion != nil {
		in, out := &in.HaSoftwareVersion, &out.HaSoftwareVersion
		*out = new(string)
		**out = **in
	}
	if in.HaSubnet != nil {
		in, out := &in.HaSubnet, &out.HaSubnet
		*out = new(string)
		**out = **in
	}
	if in.HaZone != nil {
		in, out := &in.HaZone, &out.HaZone
		*out = new(string)
		**out = **in
	}
	if in.ImageVersion != nil {
		in, out := &in.ImageVersion, &out.ImageVersion
		*out = new(string)
		**out = **in
	}
	if in.IncludedAdvertisedSpokeRoutes != nil {
		in, out := &in.IncludedAdvertisedSpokeRoutes, &out.IncludedAdvertisedSpokeRoutes
		*out = new(string)
		**out = **in
	}
	if in.InsaneMode != nil {
		in, out := &in.InsaneMode, &out.InsaneMode
		*out = new(bool)
		**out = **in
	}
	if in.InsaneModeAz != nil {
		in, out := &in.InsaneModeAz, &out.InsaneModeAz
		*out = new(string)
		**out = **in
	}
	if in.LearnedCidrsApprovalMode != nil {
		in, out := &in.LearnedCidrsApprovalMode, &out.LearnedCidrsApprovalMode
		*out = new(string)
		**out = **in
	}
	if in.LocalAsNumber != nil {
		in, out := &in.LocalAsNumber, &out.LocalAsNumber
		*out = new(string)
		**out = **in
	}
	if in.ManageHaGateway != nil {
		in, out := &in.ManageHaGateway, &out.ManageHaGateway
		*out = new(bool)
		**out = **in
	}
	if in.MonitorExcludeList != nil {
		in, out := &in.MonitorExcludeList, &out.MonitorExcludeList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OobAvailabilityZone != nil {
		in, out := &in.OobAvailabilityZone, &out.OobAvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.OobManagementSubnet != nil {
		in, out := &in.OobManagementSubnet, &out.OobManagementSubnet
		*out = new(string)
		**out = **in
	}
	if in.PrependAsPath != nil {
		in, out := &in.PrependAsPath, &out.PrependAsPath
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PrivateModeLBVPCID != nil {
		in, out := &in.PrivateModeLBVPCID, &out.PrivateModeLBVPCID
		*out = new(string)
		**out = **in
	}
	if in.PrivateModeSubnetZone != nil {
		in, out := &in.PrivateModeSubnetZone, &out.PrivateModeSubnetZone
		*out = new(string)
		**out = **in
	}
	if in.RxQueueSize != nil {
		in, out := &in.RxQueueSize, &out.RxQueueSize
		*out = new(string)
		**out = **in
	}
	if in.SingleAzHa != nil {
		in, out := &in.SingleAzHa, &out.SingleAzHa
		*out = new(bool)
		**out = **in
	}
	if in.SingleIPSnat != nil {
		in, out := &in.SingleIPSnat, &out.SingleIPSnat
		*out = new(bool)
		**out = **in
	}
	if in.SoftwareVersion != nil {
		in, out := &in.SoftwareVersion, &out.SoftwareVersion
		*out = new(string)
		**out = **in
	}
	if in.SpokeBGPManualAdvertiseCidrs != nil {
		in, out := &in.SpokeBGPManualAdvertiseCidrs, &out.SpokeBGPManualAdvertiseCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SpotPrice != nil {
		in, out := &in.SpotPrice, &out.SpotPrice
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TunnelDetectionTime != nil {
		in, out := &in.TunnelDetectionTime, &out.TunnelDetectionTime
		*out = new(float64)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCReg != nil {
		in, out := &in.VPCReg, &out.VPCReg
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayParameters.
func (in *GatewayParameters) DeepCopy() *GatewayParameters {
	if in == nil {
		return nil
	}
	out := new(GatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewaySpec) DeepCopyInto(out *GatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewaySpec.
func (in *GatewaySpec) DeepCopy() *GatewaySpec {
	if in == nil {
		return nil
	}
	out := new(GatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayStatus) DeepCopyInto(out *GatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayStatus.
func (in *GatewayStatus) DeepCopy() *GatewayStatus {
	if in == nil {
		return nil
	}
	out := new(GatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachment) DeepCopyInto(out *TransitAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachment.
func (in *TransitAttachment) DeepCopy() *TransitAttachment {
	if in == nil {
		return nil
	}
	out := new(TransitAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransitAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachmentList) DeepCopyInto(out *TransitAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TransitAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachmentList.
func (in *TransitAttachmentList) DeepCopy() *TransitAttachmentList {
	if in == nil {
		return nil
	}
	out := new(TransitAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TransitAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachmentObservation) DeepCopyInto(out *TransitAttachmentObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.SpokeBGPEnabled != nil {
		in, out := &in.SpokeBGPEnabled, &out.SpokeBGPEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachmentObservation.
func (in *TransitAttachmentObservation) DeepCopy() *TransitAttachmentObservation {
	if in == nil {
		return nil
	}
	out := new(TransitAttachmentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachmentParameters) DeepCopyInto(out *TransitAttachmentParameters) {
	*out = *in
	if in.EnableMaxPerformance != nil {
		in, out := &in.EnableMaxPerformance, &out.EnableMaxPerformance
		*out = new(bool)
		**out = **in
	}
	if in.RouteTables != nil {
		in, out := &in.RouteTables, &out.RouteTables
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SpokeGwName != nil {
		in, out := &in.SpokeGwName, &out.SpokeGwName
		*out = new(string)
		**out = **in
	}
	if in.SpokePrependAsPath != nil {
		in, out := &in.SpokePrependAsPath, &out.SpokePrependAsPath
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.TransitGwName != nil {
		in, out := &in.TransitGwName, &out.TransitGwName
		*out = new(string)
		**out = **in
	}
	if in.TransitPrependAsPath != nil {
		in, out := &in.TransitPrependAsPath, &out.TransitPrependAsPath
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachmentParameters.
func (in *TransitAttachmentParameters) DeepCopy() *TransitAttachmentParameters {
	if in == nil {
		return nil
	}
	out := new(TransitAttachmentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachmentSpec) DeepCopyInto(out *TransitAttachmentSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachmentSpec.
func (in *TransitAttachmentSpec) DeepCopy() *TransitAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(TransitAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransitAttachmentStatus) DeepCopyInto(out *TransitAttachmentStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransitAttachmentStatus.
func (in *TransitAttachmentStatus) DeepCopy() *TransitAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(TransitAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}
