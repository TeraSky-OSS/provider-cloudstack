// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NetworkInitParameters struct {

	// The ACL ID that should be attached to the network or
	// none if you do not want to attach an ACL. You can dynamically attach and
	// swap ACL's, but if you want to detach an attached ACL and revert to using
	// none, this will force a new resource to be created. (defaults none)
	ACLID *string `json:"aclId,omitempty" tf:"acl_id,omitempty"`

	// The CIDR block for the network. Changing this forces a new
	// resource to be created.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The display text of the network.
	DisplayText *string `json:"displayText,omitempty" tf:"display_text,omitempty"`

	// End of the IP block that will be available on the
	// network. Defaults to the last available IP in the range.
	Endip *string `json:"endip,omitempty" tf:"endip,omitempty"`

	// Gateway that will be provided to the instances in this
	// network. Defaults to the first usable IP in the range.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The name of the network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// DNS domain for the network.
	NetworkDomain *string `json:"networkDomain,omitempty" tf:"network_domain,omitempty"`

	// The name or ID of the network offering to use
	// for this network.
	NetworkOffering *string `json:"networkOffering,omitempty" tf:"network_offering,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If set to true a public IP will be associated
	// with the network. This is mainly used when the network supports the source
	// NAT service which claims the first associated IP address. This prevents the
	// ability to manage the IP address as an independent entity.
	SourceNATIP *bool `json:"sourceNatIp,omitempty" tf:"source_nat_ip,omitempty"`

	// Start of the IP block that will be available on the
	// network. Defaults to the second available IP in the range.
	Startip *string `json:"startip,omitempty" tf:"startip,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID in which to create this network. Changing
	// this forces a new resource to be created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VLAN number (1-4095) the network will use. This might be
	// required by the Network Offering if specifyVlan=true is set. Only the ROOT
	// admin can set this value.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`

	// The name or ID of the zone where this network will be
	// available. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NetworkObservation struct {

	// The ACL ID that should be attached to the network or
	// none if you do not want to attach an ACL. You can dynamically attach and
	// swap ACL's, but if you want to detach an attached ACL and revert to using
	// none, this will force a new resource to be created. (defaults none)
	ACLID *string `json:"aclId,omitempty" tf:"acl_id,omitempty"`

	// The CIDR block for the network. Changing this forces a new
	// resource to be created.
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The display text of the network.
	DisplayText *string `json:"displayText,omitempty" tf:"display_text,omitempty"`

	// End of the IP block that will be available on the
	// network. Defaults to the last available IP in the range.
	Endip *string `json:"endip,omitempty" tf:"endip,omitempty"`

	// Gateway that will be provided to the instances in this
	// network. Defaults to the first usable IP in the range.
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The ID of the network.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the network.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// DNS domain for the network.
	NetworkDomain *string `json:"networkDomain,omitempty" tf:"network_domain,omitempty"`

	// The name or ID of the network offering to use
	// for this network.
	NetworkOffering *string `json:"networkOffering,omitempty" tf:"network_offering,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If set to true a public IP will be associated
	// with the network. This is mainly used when the network supports the source
	// NAT service which claims the first associated IP address. This prevents the
	// ability to manage the IP address as an independent entity.
	SourceNATIP *bool `json:"sourceNatIp,omitempty" tf:"source_nat_ip,omitempty"`

	// The associated source NAT IP.
	SourceNATIPAddress *string `json:"sourceNatIpAddress,omitempty" tf:"source_nat_ip_address,omitempty"`

	// The ID of the associated source NAT IP.
	SourceNATIPID *string `json:"sourceNatIpId,omitempty" tf:"source_nat_ip_id,omitempty"`

	// Start of the IP block that will be available on the
	// network. Defaults to the second available IP in the range.
	Startip *string `json:"startip,omitempty" tf:"startip,omitempty"`

	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID in which to create this network. Changing
	// this forces a new resource to be created.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VLAN number (1-4095) the network will use. This might be
	// required by the Network Offering if specifyVlan=true is set. Only the ROOT
	// admin can set this value.
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`

	// The name or ID of the zone where this network will be
	// available. Changing this forces a new resource to be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NetworkParameters struct {

	// The ACL ID that should be attached to the network or
	// none if you do not want to attach an ACL. You can dynamically attach and
	// swap ACL's, but if you want to detach an attached ACL and revert to using
	// none, this will force a new resource to be created. (defaults none)
	// +kubebuilder:validation:Optional
	ACLID *string `json:"aclId,omitempty" tf:"acl_id,omitempty"`

	// The CIDR block for the network. Changing this forces a new
	// resource to be created.
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// The display text of the network.
	// +kubebuilder:validation:Optional
	DisplayText *string `json:"displayText,omitempty" tf:"display_text,omitempty"`

	// End of the IP block that will be available on the
	// network. Defaults to the last available IP in the range.
	// +kubebuilder:validation:Optional
	Endip *string `json:"endip,omitempty" tf:"endip,omitempty"`

	// Gateway that will be provided to the instances in this
	// network. Defaults to the first usable IP in the range.
	// +kubebuilder:validation:Optional
	Gateway *string `json:"gateway,omitempty" tf:"gateway,omitempty"`

	// The name of the network.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// DNS domain for the network.
	// +kubebuilder:validation:Optional
	NetworkDomain *string `json:"networkDomain,omitempty" tf:"network_domain,omitempty"`

	// The name or ID of the network offering to use
	// for this network.
	// +kubebuilder:validation:Optional
	NetworkOffering *string `json:"networkOffering,omitempty" tf:"network_offering,omitempty"`

	// The name or ID of the project to deploy this
	// instance to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If set to true a public IP will be associated
	// with the network. This is mainly used when the network supports the source
	// NAT service which claims the first associated IP address. This prevents the
	// ability to manage the IP address as an independent entity.
	// +kubebuilder:validation:Optional
	SourceNATIP *bool `json:"sourceNatIp,omitempty" tf:"source_nat_ip,omitempty"`

	// Start of the IP block that will be available on the
	// network. Defaults to the second available IP in the range.
	// +kubebuilder:validation:Optional
	Startip *string `json:"startip,omitempty" tf:"startip,omitempty"`

	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The VPC ID in which to create this network. Changing
	// this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VLAN number (1-4095) the network will use. This might be
	// required by the Network Offering if specifyVlan=true is set. Only the ROOT
	// admin can set this value.
	// +kubebuilder:validation:Optional
	Vlan *float64 `json:"vlan,omitempty" tf:"vlan,omitempty"`

	// The name or ID of the zone where this network will be
	// available. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NetworkInitParameters `json:"initProvider,omitempty"`
}

// NetworkStatus defines the observed state of Network.
type NetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Network is the Schema for the Networks API. Creates a network.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cidr) || (has(self.initProvider) && has(self.initProvider.cidr))",message="spec.forProvider.cidr is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.networkOffering) || (has(self.initProvider) && has(self.initProvider.networkOffering))",message="spec.forProvider.networkOffering is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   NetworkSpec   `json:"spec"`
	Status NetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Networks
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

// Repository type metadata.
var (
	Network_Kind             = "Network"
	Network_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Network_Kind}.String()
	Network_KindAPIVersion   = Network_Kind + "." + CRDGroupVersion.String()
	Network_GroupVersionKind = CRDGroupVersion.WithKind(Network_Kind)
)

func init() {
	SchemeBuilder.Register(&Network{}, &NetworkList{})
}
