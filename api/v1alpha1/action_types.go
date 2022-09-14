/*
Copyright 2022.

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
)

// SourceRef defines the reference to a Source.
type SourceRef struct {
	Name string `json:"name"`
}

// ComponentRef defines a reference to a component.
type ComponentRef struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
}

// ActionSpec defines the desired state of Action
type ActionSpec struct {
	ComponentRef ComponentRef `json:"componentRef"`
	SourceRef    SourceRef    `json:"sourceRef"`
	ProviderRef  ProviderRef  `json:"providerRef"`
}

// ActionStatus defines the observed state of Action
type ActionStatus struct {
	Ready    bool   `json:"read"`
	Snapshot string `json:"snapshot"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Action is the Schema for the actions API
type Action struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionSpec   `json:"spec,omitempty"`
	Status ActionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ActionList contains a list of Action
type ActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Action `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Action{}, &ActionList{})
}
