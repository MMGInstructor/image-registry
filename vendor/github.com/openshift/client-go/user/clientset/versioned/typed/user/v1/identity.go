// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/user/v1"
	userv1 "github.com/openshift/client-go/user/applyconfigurations/user/v1"
	scheme "github.com/openshift/client-go/user/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IdentitiesGetter has a method to return a IdentityInterface.
// A group's client should implement this interface.
type IdentitiesGetter interface {
	Identities() IdentityInterface
}

// IdentityInterface has methods to work with Identity resources.
type IdentityInterface interface {
	Create(ctx context.Context, identity *v1.Identity, opts metav1.CreateOptions) (*v1.Identity, error)
	Update(ctx context.Context, identity *v1.Identity, opts metav1.UpdateOptions) (*v1.Identity, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Identity, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IdentityList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Identity, err error)
	Apply(ctx context.Context, identity *userv1.IdentityApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Identity, err error)
	IdentityExpansion
}

// identities implements IdentityInterface
type identities struct {
	client rest.Interface
}

// newIdentities returns a Identities
func newIdentities(c *UserV1Client) *identities {
	return &identities{
		client: c.RESTClient(),
	}
}

// Get takes name of the identity, and returns the corresponding identity object, and an error if there is any.
func (c *identities) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Identity, err error) {
	result = &v1.Identity{}
	err = c.client.Get().
		Resource("identities").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Identities that match those selectors.
func (c *identities) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IdentityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IdentityList{}
	err = c.client.Get().
		Resource("identities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested identities.
func (c *identities) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("identities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a identity and creates it.  Returns the server's representation of the identity, and an error, if there is any.
func (c *identities) Create(ctx context.Context, identity *v1.Identity, opts metav1.CreateOptions) (result *v1.Identity, err error) {
	result = &v1.Identity{}
	err = c.client.Post().
		Resource("identities").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(identity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a identity and updates it. Returns the server's representation of the identity, and an error, if there is any.
func (c *identities) Update(ctx context.Context, identity *v1.Identity, opts metav1.UpdateOptions) (result *v1.Identity, err error) {
	result = &v1.Identity{}
	err = c.client.Put().
		Resource("identities").
		Name(identity.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(identity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the identity and deletes it. Returns an error if one occurs.
func (c *identities) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("identities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *identities) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("identities").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched identity.
func (c *identities) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Identity, err error) {
	result = &v1.Identity{}
	err = c.client.Patch(pt).
		Resource("identities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied identity.
func (c *identities) Apply(ctx context.Context, identity *userv1.IdentityApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Identity, err error) {
	if identity == nil {
		return nil, fmt.Errorf("identity provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(identity)
	if err != nil {
		return nil, err
	}
	name := identity.Name
	if name == nil {
		return nil, fmt.Errorf("identity.Name must be provided to Apply")
	}
	result = &v1.Identity{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("identities").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
