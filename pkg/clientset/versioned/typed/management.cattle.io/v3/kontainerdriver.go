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

package v3

import (
	"context"
	"time"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
)

// KontainerDriversGetter has a method to return a KontainerDriverInterface.
// A group's client should implement this interface.
type KontainerDriversGetter interface {
	KontainerDrivers() KontainerDriverInterface
}

// KontainerDriverInterface has methods to work with KontainerDriver resources.
type KontainerDriverInterface interface {
	Create(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.CreateOptions) (*v3.KontainerDriver, error)
	Update(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (*v3.KontainerDriver, error)
	UpdateStatus(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (*v3.KontainerDriver, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.KontainerDriver, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.KontainerDriverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.KontainerDriver, err error)
	KontainerDriverExpansion
}

// kontainerDrivers implements KontainerDriverInterface
type kontainerDrivers struct {
	client rest.Interface
}

// newKontainerDrivers returns a KontainerDrivers
func newKontainerDrivers(c *ManagementV3Client) *kontainerDrivers {
	return &kontainerDrivers{
		client: c.RESTClient(),
	}
}

// Get takes name of the kontainerDriver, and returns the corresponding kontainerDriver object, and an error if there is any.
func (c *kontainerDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.KontainerDriver, err error) {
	result = &v3.KontainerDriver{}
	err = c.client.Get().
		Resource("kontainerdrivers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KontainerDrivers that match those selectors.
func (c *kontainerDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v3.KontainerDriverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.KontainerDriverList{}
	err = c.client.Get().
		Resource("kontainerdrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kontainerDrivers.
func (c *kontainerDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("kontainerdrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kontainerDriver and creates it.  Returns the server's representation of the kontainerDriver, and an error, if there is any.
func (c *kontainerDrivers) Create(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.CreateOptions) (result *v3.KontainerDriver, err error) {
	result = &v3.KontainerDriver{}
	err = c.client.Post().
		Resource("kontainerdrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kontainerDriver).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kontainerDriver and updates it. Returns the server's representation of the kontainerDriver, and an error, if there is any.
func (c *kontainerDrivers) Update(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (result *v3.KontainerDriver, err error) {
	result = &v3.KontainerDriver{}
	err = c.client.Put().
		Resource("kontainerdrivers").
		Name(kontainerDriver.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kontainerDriver).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kontainerDrivers) UpdateStatus(ctx context.Context, kontainerDriver *v3.KontainerDriver, opts v1.UpdateOptions) (result *v3.KontainerDriver, err error) {
	result = &v3.KontainerDriver{}
	err = c.client.Put().
		Resource("kontainerdrivers").
		Name(kontainerDriver.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kontainerDriver).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kontainerDriver and deletes it. Returns an error if one occurs.
func (c *kontainerDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("kontainerdrivers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kontainerDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("kontainerdrivers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kontainerDriver.
func (c *kontainerDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.KontainerDriver, err error) {
	result = &v3.KontainerDriver{}
	err = c.client.Patch(pt).
		Resource("kontainerdrivers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
