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

package nginxapp

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	v1alpha1 "github.com/u2takey/extend_k8s_example/pkg/apis/simple/v1alpha1"
	nclient "github.com/u2takey/extend_k8s_example/pkg/client/clientset_generated/clientset"
	listers "github.com/u2takey/extend_k8s_example/pkg/client/listers_generated/simple/v1alpha1"
	"github.com/u2takey/extend_k8s_example/pkg/controller/sharedinformers"
	v1beta1 "k8s.io/api/apps/v1beta1"
	corev1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// +controller:group=simple,version=v1alpha1,kind=NginxApp,resource=nginxapps
type NginxAppControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about NginxApp
	lister listers.NginxAppLister

	kclient *kubernetes.Clientset
	nclient *nclient.Clientset
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *NginxAppControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing nginxapps labels
	c.lister = arguments.GetSharedInformers().Factory.Simple().V1alpha1().NginxApps().Lister()
	c.kclient, _ = kubernetes.NewForConfig(arguments.GetRestConfig())
	c.nclient, _ = nclient.NewForConfig(arguments.GetRestConfig())
}

// Reconcile handles enqueued messages
func (c *NginxAppControllerImpl) Reconcile(u *v1alpha1.NginxApp) error {
	// Implement controller logic here
	log.Printf("Running reconcile NginxApp for %s\n", u.Name)
	app, err := c.Get(u.Namespace, u.Name)
	if err != nil {
		return err
	}
	if app.DeletionTimestamp != nil {
		//
		log.Println("delete app")
	} else {
		label := map[string]string{"qcloud-app": u.Name}
		d := &v1beta1.Deployment{
			Spec: v1beta1.DeploymentSpec{
				Replicas: u.Spec.Replicas,
				Selector: &meta_v1.LabelSelector{MatchLabels: label},
				Template: corev1.PodTemplateSpec{
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							corev1.Container{
								Name:  "nginx",
								Image: u.Spec.Image,
							},
						},
					},
				},
			},
		}
		var controllerKind = v1alpha1.SchemeGroupVersion.WithKind("NginxApp")
		d.Spec.Template.Labels = label
		d.Namespace = u.Namespace
		d.OwnerReferences = append(d.OwnerReferences, meta_v1.OwnerReference{
			APIVersion: controllerKind.GroupVersion().String(),
			Kind:       controllerKind.Kind,
			Name:       u.Name,
			UID:        u.UID,
		})
		d.Name = u.Name
		_, err = c.kclient.AppsV1beta1().Deployments(d.Namespace).Create(d)
		if err != nil {
			log.Println("create Deployment error", err, d)
		}
		log.Println("add/update app", u)
	}
	return nil
}

// Get...
func (c *NginxAppControllerImpl) Get(namespace, name string) (*v1alpha1.NginxApp, error) {
	return c.lister.NginxApps(namespace).Get(name)
}
