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

// AuthTokensGetter has a method to return a AuthTokenInterface.
// A group's client should implement this interface.
type AuthTokensGetter interface {
	AuthTokens() AuthTokenInterface
}

// AuthTokenInterface has methods to work with AuthToken resources.
type AuthTokenInterface interface {
	Create(ctx context.Context, authToken *v3.AuthToken, opts v1.CreateOptions) (*v3.AuthToken, error)
	Update(ctx context.Context, authToken *v3.AuthToken, opts v1.UpdateOptions) (*v3.AuthToken, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.AuthToken, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.AuthTokenList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthToken, err error)
	AuthTokenExpansion
}

// authTokens implements AuthTokenInterface
type authTokens struct {
	client rest.Interface
}

// newAuthTokens returns a AuthTokens
func newAuthTokens(c *ManagementV3Client) *authTokens {
	return &authTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the authToken, and returns the corresponding authToken object, and an error if there is any.
func (c *authTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.AuthToken, err error) {
	result = &v3.AuthToken{}
	err = c.client.Get().
		Resource("authtokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AuthTokens that match those selectors.
func (c *authTokens) List(ctx context.Context, opts v1.ListOptions) (result *v3.AuthTokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.AuthTokenList{}
	err = c.client.Get().
		Resource("authtokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authTokens.
func (c *authTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("authtokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a authToken and creates it.  Returns the server's representation of the authToken, and an error, if there is any.
func (c *authTokens) Create(ctx context.Context, authToken *v3.AuthToken, opts v1.CreateOptions) (result *v3.AuthToken, err error) {
	result = &v3.AuthToken{}
	err = c.client.Post().
		Resource("authtokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authToken).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a authToken and updates it. Returns the server's representation of the authToken, and an error, if there is any.
func (c *authTokens) Update(ctx context.Context, authToken *v3.AuthToken, opts v1.UpdateOptions) (result *v3.AuthToken, err error) {
	result = &v3.AuthToken{}
	err = c.client.Put().
		Resource("authtokens").
		Name(authToken.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(authToken).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the authToken and deletes it. Returns an error if one occurs.
func (c *authTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("authtokens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("authtokens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched authToken.
func (c *authTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.AuthToken, err error) {
	result = &v3.AuthToken{}
	err = c.client.Patch(pt).
		Resource("authtokens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
