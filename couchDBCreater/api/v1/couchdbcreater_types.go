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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CouchDBCreaterSpec defines the desired state of CouchDBCreater
type CouchDBCreaterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of CouchDBCreater. Edit couchdbcreater_types.go to remove/update
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Url string `json:"url,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Database string `json:"database,omitempty"`
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	UserName string                     `json:"username,omitempty"`
	Password CouchDBCreaterSpecPassword `json:"password,omitempty"`
}

type CouchDBCreaterSpecPassword struct {
	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Secret string `json:"secret,omitempty"`

	//+kubebuilder:validation:Required
	//+kubebuilder:validation:Type=string
	Key string `json:"key,omitempty"`
}

// CouchDBCreaterStatus defines the observed state of CouchDBCreater
type CouchDBCreaterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CouchDBCreater is the Schema for the couchdbcreaters API
type CouchDBCreater struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CouchDBCreaterSpec   `json:"spec,omitempty"`
	Status CouchDBCreaterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CouchDBCreaterList contains a list of CouchDBCreater
type CouchDBCreaterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CouchDBCreater `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CouchDBCreater{}, &CouchDBCreaterList{})
}
