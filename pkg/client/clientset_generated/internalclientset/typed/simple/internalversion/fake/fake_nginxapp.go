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
package fake

import (
	simple "github.com/u2takey/extend_k8s_example/pkg/apis/simple"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNginxApps implements NginxAppInterface
type FakeNginxApps struct {
	Fake *FakeSimple
	ns   string
}

var nginxappsResource = schema.GroupVersionResource{Group: "simple.example.com", Version: "", Resource: "nginxapps"}

var nginxappsKind = schema.GroupVersionKind{Group: "simple.example.com", Version: "", Kind: "NginxApp"}

// Get takes name of the nginxApp, and returns the corresponding nginxApp object, and an error if there is any.
func (c *FakeNginxApps) Get(name string, options v1.GetOptions) (result *simple.NginxApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nginxappsResource, c.ns, name), &simple.NginxApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*simple.NginxApp), err
}

// List takes label and field selectors, and returns the list of NginxApps that match those selectors.
func (c *FakeNginxApps) List(opts v1.ListOptions) (result *simple.NginxAppList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nginxappsResource, nginxappsKind, c.ns, opts), &simple.NginxAppList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &simple.NginxAppList{}
	for _, item := range obj.(*simple.NginxAppList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nginxApps.
func (c *FakeNginxApps) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nginxappsResource, c.ns, opts))

}

// Create takes the representation of a nginxApp and creates it.  Returns the server's representation of the nginxApp, and an error, if there is any.
func (c *FakeNginxApps) Create(nginxApp *simple.NginxApp) (result *simple.NginxApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nginxappsResource, c.ns, nginxApp), &simple.NginxApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*simple.NginxApp), err
}

// Update takes the representation of a nginxApp and updates it. Returns the server's representation of the nginxApp, and an error, if there is any.
func (c *FakeNginxApps) Update(nginxApp *simple.NginxApp) (result *simple.NginxApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nginxappsResource, c.ns, nginxApp), &simple.NginxApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*simple.NginxApp), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNginxApps) UpdateStatus(nginxApp *simple.NginxApp) (*simple.NginxApp, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nginxappsResource, "status", c.ns, nginxApp), &simple.NginxApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*simple.NginxApp), err
}

// Delete takes name of the nginxApp and deletes it. Returns an error if one occurs.
func (c *FakeNginxApps) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nginxappsResource, c.ns, name), &simple.NginxApp{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNginxApps) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nginxappsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &simple.NginxAppList{})
	return err
}

// Patch applies the patch and returns the patched nginxApp.
func (c *FakeNginxApps) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *simple.NginxApp, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nginxappsResource, c.ns, name, data, subresources...), &simple.NginxApp{})

	if obj == nil {
		return nil, err
	}
	return obj.(*simple.NginxApp), err
}
