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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "github.com/rancher-sandbox/cluster-api-provider-harvester/pkg/clientset/versioned/scheme"
)

// VirtualMachineBackupsGetter has a method to return a VirtualMachineBackupInterface.
// A group's client should implement this interface.
type VirtualMachineBackupsGetter interface {
	VirtualMachineBackups(namespace string) VirtualMachineBackupInterface
}

// VirtualMachineBackupInterface has methods to work with VirtualMachineBackup resources.
type VirtualMachineBackupInterface interface {
	Create(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.CreateOptions) (*v1beta1.VirtualMachineBackup, error)
	Update(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (*v1beta1.VirtualMachineBackup, error)
	UpdateStatus(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (*v1beta1.VirtualMachineBackup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineBackup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineBackupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineBackup, err error)
	VirtualMachineBackupExpansion
}

// virtualMachineBackups implements VirtualMachineBackupInterface
type virtualMachineBackups struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineBackups returns a VirtualMachineBackups
func newVirtualMachineBackups(c *HarvesterhciV1beta1Client, namespace string) *virtualMachineBackups {
	return &virtualMachineBackups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineBackup, and returns the corresponding virtualMachineBackup object, and an error if there is any.
func (c *virtualMachineBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	result = &v1beta1.VirtualMachineBackup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineBackups that match those selectors.
func (c *virtualMachineBackups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineBackupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VirtualMachineBackupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineBackups.
func (c *virtualMachineBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineBackup and creates it.  Returns the server's representation of the virtualMachineBackup, and an error, if there is any.
func (c *virtualMachineBackups) Create(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.CreateOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	result = &v1beta1.VirtualMachineBackup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineBackup and updates it. Returns the server's representation of the virtualMachineBackup, and an error, if there is any.
func (c *virtualMachineBackups) Update(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	result = &v1beta1.VirtualMachineBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		Name(virtualMachineBackup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualMachineBackups) UpdateStatus(ctx context.Context, virtualMachineBackup *v1beta1.VirtualMachineBackup, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineBackup, err error) {
	result = &v1beta1.VirtualMachineBackup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		Name(virtualMachineBackup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineBackup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineBackup and deletes it. Returns an error if one occurs.
func (c *virtualMachineBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineBackup.
func (c *virtualMachineBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineBackup, err error) {
	result = &v1beta1.VirtualMachineBackup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinebackups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
