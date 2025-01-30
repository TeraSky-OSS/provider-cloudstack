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

type ForwardForwardInitParameters struct {

	// The private port to forward to.
	PrivatePort *float64 `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp and udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public port to forward from.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`

	// The virtual machine IP address for the port
	// forwarding rule (useful when the virtual machine has secondairy NICs
	// or IP addresses).
	VMGuestIP *string `json:"vmGuestIp,omitempty" tf:"vm_guest_ip,omitempty"`

	// The ID of the virtual machine to forward to.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type ForwardForwardObservation struct {

	// The private port to forward to.
	PrivatePort *float64 `json:"privatePort,omitempty" tf:"private_port,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp and udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public port to forward from.
	PublicPort *float64 `json:"publicPort,omitempty" tf:"public_port,omitempty"`

	// The ID of the IP address for which the port forwards are created.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// The virtual machine IP address for the port
	// forwarding rule (useful when the virtual machine has secondairy NICs
	// or IP addresses).
	VMGuestIP *string `json:"vmGuestIp,omitempty" tf:"vm_guest_ip,omitempty"`

	// The ID of the virtual machine to forward to.
	VirtualMachineID *string `json:"virtualMachineId,omitempty" tf:"virtual_machine_id,omitempty"`
}

type ForwardForwardParameters struct {

	// The private port to forward to.
	// +kubebuilder:validation:Optional
	PrivatePort *float64 `json:"privatePort" tf:"private_port,omitempty"`

	// The name of the protocol to allow. Valid options are:
	// tcp and udp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The public port to forward from.
	// +kubebuilder:validation:Optional
	PublicPort *float64 `json:"publicPort" tf:"public_port,omitempty"`

	// The virtual machine IP address for the port
	// forwarding rule (useful when the virtual machine has secondairy NICs
	// or IP addresses).
	// +kubebuilder:validation:Optional
	VMGuestIP *string `json:"vmGuestIp,omitempty" tf:"vm_guest_ip,omitempty"`

	// The ID of the virtual machine to forward to.
	// +kubebuilder:validation:Optional
	VirtualMachineID *string `json:"virtualMachineId" tf:"virtual_machine_id,omitempty"`
}

type ForwardInitParameters struct {

	// Can be specified multiple times. Each forward block supports
	// fields documented below.
	Forward []ForwardForwardInitParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// The IP address ID for which to create the port
	// forwards. Changing this forces a new resource to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the port forwards for
	// this IP address will be managed by this resource. This means it will delete
	// all port forwards that are not in your config! (defaults false)
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// The name or ID of the project to create this port forward
	// in. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ForwardObservation struct {

	// Can be specified multiple times. Each forward block supports
	// fields documented below.
	Forward []ForwardForwardObservation `json:"forward,omitempty" tf:"forward,omitempty"`

	// The ID of the IP address for which the port forwards are created.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The IP address ID for which to create the port
	// forwards. Changing this forces a new resource to be created.
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the port forwards for
	// this IP address will be managed by this resource. This means it will delete
	// all port forwards that are not in your config! (defaults false)
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// The name or ID of the project to create this port forward
	// in. Changing this forces a new resource to be created.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ForwardParameters struct {

	// Can be specified multiple times. Each forward block supports
	// fields documented below.
	// +kubebuilder:validation:Optional
	Forward []ForwardForwardParameters `json:"forward,omitempty" tf:"forward,omitempty"`

	// The IP address ID for which to create the port
	// forwards. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	IPAddressID *string `json:"ipAddressId,omitempty" tf:"ip_address_id,omitempty"`

	// USE WITH CAUTION! If enabled all the port forwards for
	// this IP address will be managed by this resource. This means it will delete
	// all port forwards that are not in your config! (defaults false)
	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// The name or ID of the project to create this port forward
	// in. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ForwardSpec defines the desired state of Forward
type ForwardSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ForwardParameters `json:"forProvider"`
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
	InitProvider ForwardInitParameters `json:"initProvider,omitempty"`
}

// ForwardStatus defines the observed state of Forward.
type ForwardStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ForwardObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Forward is the Schema for the Forwards API. Creates port forwards.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudstack}
type Forward struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.forward) || (has(self.initProvider) && has(self.initProvider.forward))",message="spec.forProvider.forward is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ipAddressId) || (has(self.initProvider) && has(self.initProvider.ipAddressId))",message="spec.forProvider.ipAddressId is a required parameter"
	Spec   ForwardSpec   `json:"spec"`
	Status ForwardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ForwardList contains a list of Forwards
type ForwardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Forward `json:"items"`
}

// Repository type metadata.
var (
	Forward_Kind             = "Forward"
	Forward_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Forward_Kind}.String()
	Forward_KindAPIVersion   = Forward_Kind + "." + CRDGroupVersion.String()
	Forward_GroupVersionKind = CRDGroupVersion.WithKind(Forward_Kind)
)

func init() {
	SchemeBuilder.Register(&Forward{}, &ForwardList{})
}
