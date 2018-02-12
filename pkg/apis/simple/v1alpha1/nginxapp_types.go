/*
Copyright 2017 The Kubernetes Authors.

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
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/u2takey/extend_k8s_example/pkg/apis/simple"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NginxApp
// +k8s:openapi-gen=true
// +resource:path=nginxapps,strategy=NginxAppStrategy
type NginxApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   NginxAppSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status NginxAppStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// NginxAppSpec defines the desired state of NginxApp
type NginxAppSpec struct {
	Image string `json:"image,omitempty" protobuf:"varint,1,opt,name=image"`
	// Number of desired pods. This is a pointer to distinguish between explicit
	// zero and not specified. Defaults to 1.
	// +optional
	Replicas *int32 `json:"replicas,omitempty" protobuf:"varint,2,opt,name=replicas"`
}

// NginxAppStatus defines the observed state of NginxApp
type NginxAppStatus struct {
	// Total number of non-terminated pods targeted by this deployment (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
}

// Validate checks that an instance of NginxApp is well formed
func (NginxAppStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*simple.NginxApp)
	log.Printf("Validating fields for NginxApp %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid

	if o.Spec.Replicas != nil && *o.Spec.Replicas >= 10 {
		path := field.NewPath("spec")
		errors = append(errors, field.Invalid(path, o.Spec.Replicas, "replicas must less than 10"))
	}
	return errors
}

// DefaultingFunction sets default NginxApp field values
func (NginxAppSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*NginxApp)
	// set default field values here
	if obj.Spec.Replicas == nil {
		var a int32 = 1
		obj.Spec.Replicas = &a
	}
	// if obj.Finalizers == nil {
	// 	obj.Finalizers = append(obj.Finalizers, "test.example.com")
	// }
	log.Printf("Defaulting fields for NginxApp %s\n", obj.Name)
}
