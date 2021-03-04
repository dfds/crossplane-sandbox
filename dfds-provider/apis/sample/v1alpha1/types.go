/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// ServiceProxyOperatorParameters are the configurable fields of a ServiceProxyOperator.
type ServiceProxyOperatorParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ServiceProxyOperatorObservation are the observable fields of a ServiceProxyOperator.
type ServiceProxyOperatorObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ServiceProxyOperatorSpec defines the desired state of a ServiceProxyOperator.
type ServiceProxyOperatorSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServiceProxyOperatorParameters `json:"forProvider"`
}

// A ServiceProxyOperatorStatus represents the observed state of a ServiceProxyOperator.
type ServiceProxyOperatorStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServiceProxyOperatorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A ServiceProxyOperator is an example API type
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.bindingPhase"
// +kubebuilder:printcolumn:name="STATE",type="string",JSONPath=".status.atProvider.state"
// +kubebuilder:printcolumn:name="CLASS",type="string",JSONPath=".spec.classRef.name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster
type ServiceProxyOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceProxyOperatorSpec   `json:"spec"`
	Status ServiceProxyOperatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceProxyOperatorList contains a list of ServiceProxyOperator
type ServiceProxyOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceProxyOperator `json:"items"`
}
