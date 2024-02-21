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

// SamlProvidersGetter has a method to return a SamlProviderInterface.
// A group's client should implement this interface.
type SamlProvidersGetter interface {
	SamlProviders() SamlProviderInterface
}

// SamlProviderInterface has methods to work with SamlProvider resources.
type SamlProviderInterface interface {
	Create(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.CreateOptions) (*v3.SamlProvider, error)
	Update(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.UpdateOptions) (*v3.SamlProvider, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.SamlProvider, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.SamlProviderList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SamlProvider, err error)
	SamlProviderExpansion
}

// samlProviders implements SamlProviderInterface
type samlProviders struct {
	client rest.Interface
}

// newSamlProviders returns a SamlProviders
func newSamlProviders(c *ManagementV3Client) *samlProviders {
	return &samlProviders{
		client: c.RESTClient(),
	}
}

// Get takes name of the samlProvider, and returns the corresponding samlProvider object, and an error if there is any.
func (c *samlProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.SamlProvider, err error) {
	result = &v3.SamlProvider{}
	err = c.client.Get().
		Resource("samlproviders").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SamlProviders that match those selectors.
func (c *samlProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.SamlProviderList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.SamlProviderList{}
	err = c.client.Get().
		Resource("samlproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested samlProviders.
func (c *samlProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("samlproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a samlProvider and creates it.  Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *samlProviders) Create(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.CreateOptions) (result *v3.SamlProvider, err error) {
	result = &v3.SamlProvider{}
	err = c.client.Post().
		Resource("samlproviders").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(samlProvider).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a samlProvider and updates it. Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *samlProviders) Update(ctx context.Context, samlProvider *v3.SamlProvider, opts v1.UpdateOptions) (result *v3.SamlProvider, err error) {
	result = &v3.SamlProvider{}
	err = c.client.Put().
		Resource("samlproviders").
		Name(samlProvider.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(samlProvider).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the samlProvider and deletes it. Returns an error if one occurs.
func (c *samlProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("samlproviders").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *samlProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("samlproviders").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched samlProvider.
func (c *samlProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.SamlProvider, err error) {
	result = &v3.SamlProvider{}
	err = c.client.Patch(pt).
		Resource("samlproviders").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
