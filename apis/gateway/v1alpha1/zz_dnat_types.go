/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionPolicyObservation struct {

	// This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True. Available as of provider version R2.19.2+.
	ApplyRouteEntry *bool `json:"applyRouteEntry,omitempty" tf:"apply_route_entry,omitempty"`

	// This is a qualifier condition that specifies output connection where the rule applies. Default value: "None".
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	DnatIps *string `json:"dnatIps,omitempty" tf:"dnat_ips,omitempty"`

	// This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	DnatPort *string `json:"dnatPort,omitempty" tf:"dnat_port,omitempty"`

	// This is a qualifier condition that specifies a destination IP address range where the rule applies. When not specified, this field is not used.
	DstCidr *string `json:"dstCidr,omitempty" tf:"dst_cidr,omitempty"`

	// This is a qualifier condition that specifies a destination port where the rule applies. When not specified, this field is not used.
	DstPort *string `json:"dstPort,omitempty" tf:"dst_port,omitempty"`

	// This field specifies which VPC private route table will not be programmed with the default route entry.
	ExcludeRtb *string `json:"excludeRtb,omitempty" tf:"exclude_rtb,omitempty"`

	// based connection in a policy.
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When not specified, this field is not used.
	Mark *string `json:"mark,omitempty" tf:"mark,omitempty"`

	// This is a qualifier condition that specifies a destination port protocol where the rule applies. When not specified, this field is not used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This is a qualifier condition that specifies a source IP address range where the rule applies. When not specified, this field is not used.
	SrcCidr *string `json:"srcCidr,omitempty" tf:"src_cidr,omitempty"`

	// This is a qualifier condition that specifies a source port that the rule applies. When not specified, this field is not used.
	SrcPort *string `json:"srcPort,omitempty" tf:"src_port,omitempty"`
}

type ConnectionPolicyParameters struct {
}

type DnatObservation struct {

	// Computed attribute to store the previous connection policy.
	ConnectionPolicy []ConnectionPolicyObservation `json:"connectionPolicy,omitempty" tf:"connection_policy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Computed attribute to store the previous interface policy.
	InterfacePolicy []InterfacePolicyObservation `json:"interfacePolicy,omitempty" tf:"interface_policy,omitempty"`
}

type DnatParameters struct {

	// Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and Azure(8).
	// Policy rule to be applied to gateway.
	// +kubebuilder:validation:Required
	DnatPolicy []DnatPolicyParameters `json:"dnatPolicy" tf:"dnat_policy,omitempty"`
}

type DnatPolicyObservation struct {
}

type DnatPolicyParameters struct {

	// This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True. Available as of provider version R2.19.2+.
	// This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True.
	// +kubebuilder:validation:Optional
	ApplyRouteEntry *bool `json:"applyRouteEntry,omitempty" tf:"apply_route_entry,omitempty"`

	// This is a qualifier condition that specifies output connection where the rule applies. Default value: "None".
	// This is a qualifier condition that specifies output connection where the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	// This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule field must be specified for this rule to take effect.
	// +kubebuilder:validation:Optional
	DnatIps *string `json:"dnatIps,omitempty" tf:"dnat_ips,omitempty"`

	// This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	// This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When left blank, this field is not used. One of the rule field must be specified for this rule to take effect.
	// +kubebuilder:validation:Optional
	DnatPort *string `json:"dnatPort,omitempty" tf:"dnat_port,omitempty"`

	// This is a qualifier condition that specifies a destination IP address range where the rule applies. When not specified, this field is not used.
	// This is a qualifier condition that specifies a destination IP address range where the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	DstCidr *string `json:"dstCidr,omitempty" tf:"dst_cidr,omitempty"`

	// This is a qualifier condition that specifies a destination port where the rule applies. When not specified, this field is not used.
	// This is a qualifier condition that specifies a destination port where the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	DstPort *string `json:"dstPort,omitempty" tf:"dst_port,omitempty"`

	// This field specifies which VPC private route table will not be programmed with the default route entry.
	// This field specifies which VPC private route table will not be programmed with the default route entry.
	// +kubebuilder:validation:Optional
	ExcludeRtb *string `json:"excludeRtb,omitempty" tf:"exclude_rtb,omitempty"`

	// based connection in a policy.
	// This is a qualifier condition that specifies output interface where the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When not specified, this field is not used.
	// This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	Mark *string `json:"mark,omitempty" tf:"mark,omitempty"`

	// This is a qualifier condition that specifies a destination port protocol where the rule applies. When not specified, this field is not used.
	// This is a qualifier condition that specifies a destination port protocol where the rule applies. Default: all.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This is a qualifier condition that specifies a source IP address range where the rule applies. When not specified, this field is not used.
	// This is a qualifier condition that specifies a source IP address range where the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	SrcCidr *string `json:"srcCidr,omitempty" tf:"src_cidr,omitempty"`

	// This is a qualifier condition that specifies a source port that the rule applies. When not specified, this field is not used.
	// This is a qualifier condition that specifies a source port that the rule applies. When left blank, this field is not used.
	// +kubebuilder:validation:Optional
	SrcPort *string `json:"srcPort,omitempty" tf:"src_port,omitempty"`
}

type InterfacePolicyObservation struct {

	// This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True. Available as of provider version R2.19.2+.
	ApplyRouteEntry *bool `json:"applyRouteEntry,omitempty" tf:"apply_route_entry,omitempty"`

	// This is a qualifier condition that specifies output connection where the rule applies. Default value: "None".
	Connection *string `json:"connection,omitempty" tf:"connection,omitempty"`

	// This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	DnatIps *string `json:"dnatIps,omitempty" tf:"dnat_ips,omitempty"`

	// This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.
	DnatPort *string `json:"dnatPort,omitempty" tf:"dnat_port,omitempty"`

	// This is a qualifier condition that specifies a destination IP address range where the rule applies. When not specified, this field is not used.
	DstCidr *string `json:"dstCidr,omitempty" tf:"dst_cidr,omitempty"`

	// This is a qualifier condition that specifies a destination port where the rule applies. When not specified, this field is not used.
	DstPort *string `json:"dstPort,omitempty" tf:"dst_port,omitempty"`

	// This field specifies which VPC private route table will not be programmed with the default route entry.
	ExcludeRtb *string `json:"excludeRtb,omitempty" tf:"exclude_rtb,omitempty"`

	// based connection in a policy.
	Interface *string `json:"interface,omitempty" tf:"interface,omitempty"`

	// This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When not specified, this field is not used.
	Mark *string `json:"mark,omitempty" tf:"mark,omitempty"`

	// This is a qualifier condition that specifies a destination port protocol where the rule applies. When not specified, this field is not used.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// This is a qualifier condition that specifies a source IP address range where the rule applies. When not specified, this field is not used.
	SrcCidr *string `json:"srcCidr,omitempty" tf:"src_cidr,omitempty"`

	// This is a qualifier condition that specifies a source port that the rule applies. When not specified, this field is not used.
	SrcPort *string `json:"srcPort,omitempty" tf:"src_port,omitempty"`
}

type InterfacePolicyParameters struct {
}

// DnatSpec defines the desired state of Dnat
type DnatSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DnatParameters `json:"forProvider"`
}

// DnatStatus defines the observed state of Dnat.
type DnatStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DnatObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Dnat is the Schema for the Dnats API. Configure policies for destination NAT for an Aviatrix gateway
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aviatrix}
type Dnat struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnatSpec   `json:"spec"`
	Status            DnatStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DnatList contains a list of Dnats
type DnatList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dnat `json:"items"`
}

// Repository type metadata.
var (
	Dnat_Kind             = "Dnat"
	Dnat_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Dnat_Kind}.String()
	Dnat_KindAPIVersion   = Dnat_Kind + "." + CRDGroupVersion.String()
	Dnat_GroupVersionKind = CRDGroupVersion.WithKind(Dnat_Kind)
)

func init() {
	SchemeBuilder.Register(&Dnat{}, &DnatList{})
}
