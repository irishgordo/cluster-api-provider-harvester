/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1beta2

import (
	"context"
	"time"

	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
)

// SystemRestoresGetter has a method to return a SystemRestoreInterface.
// A group's client should implement this interface.
type SystemRestoresGetter interface {
	SystemRestores(namespace string) SystemRestoreInterface
}

// SystemRestoreInterface has methods to work with SystemRestore resources.
type SystemRestoreInterface interface {
	Create(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.CreateOptions) (*v1beta2.SystemRestore, error)
	Update(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.UpdateOptions) (*v1beta2.SystemRestore, error)
	UpdateStatus(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.UpdateOptions) (*v1beta2.SystemRestore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.SystemRestore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.SystemRestoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.SystemRestore, err error)
	SystemRestoreExpansion
}

// systemRestores implements SystemRestoreInterface
type systemRestores struct {
	client rest.Interface
	ns     string
}

// newSystemRestores returns a SystemRestores
func newSystemRestores(c *LonghornV1beta2Client, namespace string) *systemRestores {
	return &systemRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the systemRestore, and returns the corresponding systemRestore object, and an error if there is any.
func (c *systemRestores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.SystemRestore, err error) {
	result = &v1beta2.SystemRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("systemrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SystemRestores that match those selectors.
func (c *systemRestores) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.SystemRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.SystemRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("systemrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested systemRestores.
func (c *systemRestores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("systemrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a systemRestore and creates it.  Returns the server's representation of the systemRestore, and an error, if there is any.
func (c *systemRestores) Create(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.CreateOptions) (result *v1beta2.SystemRestore, err error) {
	result = &v1beta2.SystemRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("systemrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemRestore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a systemRestore and updates it. Returns the server's representation of the systemRestore, and an error, if there is any.
func (c *systemRestores) Update(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.UpdateOptions) (result *v1beta2.SystemRestore, err error) {
	result = &v1beta2.SystemRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("systemrestores").
		Name(systemRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemRestore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *systemRestores) UpdateStatus(ctx context.Context, systemRestore *v1beta2.SystemRestore, opts v1.UpdateOptions) (result *v1beta2.SystemRestore, err error) {
	result = &v1beta2.SystemRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("systemrestores").
		Name(systemRestore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(systemRestore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the systemRestore and deletes it. Returns an error if one occurs.
func (c *systemRestores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("systemrestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *systemRestores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("systemrestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched systemRestore.
func (c *systemRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.SystemRestore, err error) {
	result = &v1beta2.SystemRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("systemrestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
