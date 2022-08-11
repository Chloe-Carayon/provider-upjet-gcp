/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ServiceNetworkingPeeredDNSDomainObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

type ServiceNetworkingPeeredDNSDomainParameters struct {

	// The DNS domain name suffix of the peered DNS domain.
	// +kubebuilder:validation:Required
	DNSSuffix *string `json:"dnsSuffix" tf:"dns_suffix,omitempty"`

	// Network in the consumer project to peer with.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The ID of the project that the service account will be created in. Defaults to the provider project configuration.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The name of the service to create a peered DNS domain for, e.g. servicenetworking.googleapis.com
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ServiceNetworkingPeeredDNSDomainSpec defines the desired state of ServiceNetworkingPeeredDNSDomain
type ServiceNetworkingPeeredDNSDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceNetworkingPeeredDNSDomainParameters `json:"forProvider"`
}

// ServiceNetworkingPeeredDNSDomainStatus defines the observed state of ServiceNetworkingPeeredDNSDomain.
type ServiceNetworkingPeeredDNSDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceNetworkingPeeredDNSDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkingPeeredDNSDomain is the Schema for the ServiceNetworkingPeeredDNSDomains API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceNetworkingPeeredDNSDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceNetworkingPeeredDNSDomainSpec   `json:"spec"`
	Status            ServiceNetworkingPeeredDNSDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceNetworkingPeeredDNSDomainList contains a list of ServiceNetworkingPeeredDNSDomains
type ServiceNetworkingPeeredDNSDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceNetworkingPeeredDNSDomain `json:"items"`
}

// Repository type metadata.
var (
	ServiceNetworkingPeeredDNSDomain_Kind             = "ServiceNetworkingPeeredDNSDomain"
	ServiceNetworkingPeeredDNSDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceNetworkingPeeredDNSDomain_Kind}.String()
	ServiceNetworkingPeeredDNSDomain_KindAPIVersion   = ServiceNetworkingPeeredDNSDomain_Kind + "." + CRDGroupVersion.String()
	ServiceNetworkingPeeredDNSDomain_GroupVersionKind = CRDGroupVersion.WithKind(ServiceNetworkingPeeredDNSDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceNetworkingPeeredDNSDomain{}, &ServiceNetworkingPeeredDNSDomainList{})
}
