/*
Copyright 2024.

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

// EmailSenderConfigSpec defines the desired state of EmailSenderConfig
type EmailSenderConfigSpec struct {
	// ApiTokenSecretRef is a reference to the secret containing the API token
	ApiTokenSecretRef string `json:"apiTokenSecretRef"`
	// SenderEmail is the email address used for sending emails
	SenderEmail string `json:"senderEmail"`
	// Add other fields as needed
}

// EmailSenderConfigStatus defines the observed state of EmailSenderConfig
type EmailSenderConfigStatus struct {
	// Add observed state fields as needed
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EmailSenderConfig is the Schema for the emailsenderconfigs API
type EmailSenderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailSenderConfigSpec   `json:"spec,omitempty"`
	Status EmailSenderConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EmailSenderConfigList contains a list of EmailSenderConfig
type EmailSenderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailSenderConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmailSenderConfig{}, &EmailSenderConfigList{})
}
