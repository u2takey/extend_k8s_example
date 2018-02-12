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
	v1alpha1 "github.com/u2takey/extend_k8s_example/pkg/apis/simple/v1alpha1"
	scheme "github.com/u2takey/extend_k8s_example/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NginxAppsGetter has a method to return a NginxAppInterface.
// A group's client should implement this interface.
type NginxAppsGetter interface {
	NginxApps(namespace string) NginxAppInterface
}

// NginxAppInterface has methods to work with NginxApp resources.
type NginxAppInterface interface {
	Create(*v1alpha1.NginxApp) (*v1alpha1.NginxApp, error)
	Update(*v1alpha1.NginxApp) (*v1alpha1.NginxApp, error)
	UpdateStatus(*v1alpha1.NginxApp) (*v1alpha1.NginxApp, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NginxApp, error)
	List(opts v1.ListOptions) (*v1alpha1.NginxAppList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NginxApp, err error)
	NginxAppExpansion
}

// nginxApps implements NginxAppInterface
type nginxApps struct {
	client rest.Interface
	ns     string
}

// newNginxApps returns a NginxApps
func newNginxApps(c *SimpleV1alpha1Client, namespace string) *nginxApps {
	return &nginxApps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nginxApp, and returns the corresponding nginxApp object, and an error if there is any.
func (c *nginxApps) Get(name string, options v1.GetOptions) (result *v1alpha1.NginxApp, err error) {
	result = &v1alpha1.NginxApp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nginxapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NginxApps that match those selectors.
func (c *nginxApps) List(opts v1.ListOptions) (result *v1alpha1.NginxAppList, err error) {
	result = &v1alpha1.NginxAppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nginxapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nginxApps.
func (c *nginxApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nginxapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nginxApp and creates it.  Returns the server's representation of the nginxApp, and an error, if there is any.
func (c *nginxApps) Create(nginxApp *v1alpha1.NginxApp) (result *v1alpha1.NginxApp, err error) {
	result = &v1alpha1.NginxApp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nginxapps").
		Body(nginxApp).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nginxApp and updates it. Returns the server's representation of the nginxApp, and an error, if there is any.
func (c *nginxApps) Update(nginxApp *v1alpha1.NginxApp) (result *v1alpha1.NginxApp, err error) {
	result = &v1alpha1.NginxApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nginxapps").
		Name(nginxApp.Name).
		Body(nginxApp).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nginxApps) UpdateStatus(nginxApp *v1alpha1.NginxApp) (result *v1alpha1.NginxApp, err error) {
	result = &v1alpha1.NginxApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nginxapps").
		Name(nginxApp.Name).
		SubResource("status").
		Body(nginxApp).
		Do().
		Into(result)
	return
}

// Delete takes name of the nginxApp and deletes it. Returns an error if one occurs.
func (c *nginxApps) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nginxapps").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nginxApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nginxapps").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nginxApp.
func (c *nginxApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NginxApp, err error) {
	result = &v1alpha1.NginxApp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nginxapps").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
