// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	simple "github.com/u2takey/extend_k8s_example/pkg/apis/simple"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_NginxApp_To_simple_NginxApp,
		Convert_simple_NginxApp_To_v1alpha1_NginxApp,
		Convert_v1alpha1_NginxAppList_To_simple_NginxAppList,
		Convert_simple_NginxAppList_To_v1alpha1_NginxAppList,
		Convert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec,
		Convert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec,
		Convert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus,
		Convert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus,
		Convert_v1alpha1_NginxAppStatusStrategy_To_simple_NginxAppStatusStrategy,
		Convert_simple_NginxAppStatusStrategy_To_v1alpha1_NginxAppStatusStrategy,
		Convert_v1alpha1_NginxAppStrategy_To_simple_NginxAppStrategy,
		Convert_simple_NginxAppStrategy_To_v1alpha1_NginxAppStrategy,
	)
}

func autoConvert_v1alpha1_NginxApp_To_simple_NginxApp(in *NginxApp, out *simple.NginxApp, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_NginxApp_To_simple_NginxApp is an autogenerated conversion function.
func Convert_v1alpha1_NginxApp_To_simple_NginxApp(in *NginxApp, out *simple.NginxApp, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxApp_To_simple_NginxApp(in, out, s)
}

func autoConvert_simple_NginxApp_To_v1alpha1_NginxApp(in *simple.NginxApp, out *NginxApp, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_simple_NginxApp_To_v1alpha1_NginxApp is an autogenerated conversion function.
func Convert_simple_NginxApp_To_v1alpha1_NginxApp(in *simple.NginxApp, out *NginxApp, s conversion.Scope) error {
	return autoConvert_simple_NginxApp_To_v1alpha1_NginxApp(in, out, s)
}

func autoConvert_v1alpha1_NginxAppList_To_simple_NginxAppList(in *NginxAppList, out *simple.NginxAppList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]simple.NginxApp)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_NginxAppList_To_simple_NginxAppList is an autogenerated conversion function.
func Convert_v1alpha1_NginxAppList_To_simple_NginxAppList(in *NginxAppList, out *simple.NginxAppList, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxAppList_To_simple_NginxAppList(in, out, s)
}

func autoConvert_simple_NginxAppList_To_v1alpha1_NginxAppList(in *simple.NginxAppList, out *NginxAppList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]NginxApp)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_simple_NginxAppList_To_v1alpha1_NginxAppList is an autogenerated conversion function.
func Convert_simple_NginxAppList_To_v1alpha1_NginxAppList(in *simple.NginxAppList, out *NginxAppList, s conversion.Scope) error {
	return autoConvert_simple_NginxAppList_To_v1alpha1_NginxAppList(in, out, s)
}

func autoConvert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec(in *NginxAppSpec, out *simple.NginxAppSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	return nil
}

// Convert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec is an autogenerated conversion function.
func Convert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec(in *NginxAppSpec, out *simple.NginxAppSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxAppSpec_To_simple_NginxAppSpec(in, out, s)
}

func autoConvert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec(in *simple.NginxAppSpec, out *NginxAppSpec, s conversion.Scope) error {
	out.Image = in.Image
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	return nil
}

// Convert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec is an autogenerated conversion function.
func Convert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec(in *simple.NginxAppSpec, out *NginxAppSpec, s conversion.Scope) error {
	return autoConvert_simple_NginxAppSpec_To_v1alpha1_NginxAppSpec(in, out, s)
}

func autoConvert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus(in *NginxAppStatus, out *simple.NginxAppStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus is an autogenerated conversion function.
func Convert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus(in *NginxAppStatus, out *simple.NginxAppStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxAppStatus_To_simple_NginxAppStatus(in, out, s)
}

func autoConvert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus(in *simple.NginxAppStatus, out *NginxAppStatus, s conversion.Scope) error {
	out.Replicas = in.Replicas
	return nil
}

// Convert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus is an autogenerated conversion function.
func Convert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus(in *simple.NginxAppStatus, out *NginxAppStatus, s conversion.Scope) error {
	return autoConvert_simple_NginxAppStatus_To_v1alpha1_NginxAppStatus(in, out, s)
}

func autoConvert_v1alpha1_NginxAppStatusStrategy_To_simple_NginxAppStatusStrategy(in *NginxAppStatusStrategy, out *simple.NginxAppStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_NginxAppStatusStrategy_To_simple_NginxAppStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_NginxAppStatusStrategy_To_simple_NginxAppStatusStrategy(in *NginxAppStatusStrategy, out *simple.NginxAppStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxAppStatusStrategy_To_simple_NginxAppStatusStrategy(in, out, s)
}

func autoConvert_simple_NginxAppStatusStrategy_To_v1alpha1_NginxAppStatusStrategy(in *simple.NginxAppStatusStrategy, out *NginxAppStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_simple_NginxAppStatusStrategy_To_v1alpha1_NginxAppStatusStrategy is an autogenerated conversion function.
func Convert_simple_NginxAppStatusStrategy_To_v1alpha1_NginxAppStatusStrategy(in *simple.NginxAppStatusStrategy, out *NginxAppStatusStrategy, s conversion.Scope) error {
	return autoConvert_simple_NginxAppStatusStrategy_To_v1alpha1_NginxAppStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_NginxAppStrategy_To_simple_NginxAppStrategy(in *NginxAppStrategy, out *simple.NginxAppStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_NginxAppStrategy_To_simple_NginxAppStrategy is an autogenerated conversion function.
func Convert_v1alpha1_NginxAppStrategy_To_simple_NginxAppStrategy(in *NginxAppStrategy, out *simple.NginxAppStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_NginxAppStrategy_To_simple_NginxAppStrategy(in, out, s)
}

func autoConvert_simple_NginxAppStrategy_To_v1alpha1_NginxAppStrategy(in *simple.NginxAppStrategy, out *NginxAppStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_simple_NginxAppStrategy_To_v1alpha1_NginxAppStrategy is an autogenerated conversion function.
func Convert_simple_NginxAppStrategy_To_v1alpha1_NginxAppStrategy(in *simple.NginxAppStrategy, out *NginxAppStrategy, s conversion.Scope) error {
	return autoConvert_simple_NginxAppStrategy_To_v1alpha1_NginxAppStrategy(in, out, s)
}
