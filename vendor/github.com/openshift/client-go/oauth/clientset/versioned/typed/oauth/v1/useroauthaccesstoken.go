// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/oauth/v1"
	oauthv1 "github.com/openshift/client-go/oauth/applyconfigurations/oauth/v1"
	scheme "github.com/openshift/client-go/oauth/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserOAuthAccessTokensGetter has a method to return a UserOAuthAccessTokenInterface.
// A group's client should implement this interface.
type UserOAuthAccessTokensGetter interface {
	UserOAuthAccessTokens() UserOAuthAccessTokenInterface
}

// UserOAuthAccessTokenInterface has methods to work with UserOAuthAccessToken resources.
type UserOAuthAccessTokenInterface interface {
	Create(ctx context.Context, userOAuthAccessToken *v1.UserOAuthAccessToken, opts metav1.CreateOptions) (*v1.UserOAuthAccessToken, error)
	Update(ctx context.Context, userOAuthAccessToken *v1.UserOAuthAccessToken, opts metav1.UpdateOptions) (*v1.UserOAuthAccessToken, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.UserOAuthAccessToken, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserOAuthAccessTokenList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserOAuthAccessToken, err error)
	Apply(ctx context.Context, userOAuthAccessToken *oauthv1.UserOAuthAccessTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.UserOAuthAccessToken, err error)
	UserOAuthAccessTokenExpansion
}

// userOAuthAccessTokens implements UserOAuthAccessTokenInterface
type userOAuthAccessTokens struct {
	client rest.Interface
}

// newUserOAuthAccessTokens returns a UserOAuthAccessTokens
func newUserOAuthAccessTokens(c *OauthV1Client) *userOAuthAccessTokens {
	return &userOAuthAccessTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the userOAuthAccessToken, and returns the corresponding userOAuthAccessToken object, and an error if there is any.
func (c *userOAuthAccessTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.UserOAuthAccessToken, err error) {
	result = &v1.UserOAuthAccessToken{}
	err = c.client.Get().
		Resource("useroauthaccesstokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserOAuthAccessTokens that match those selectors.
func (c *userOAuthAccessTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UserOAuthAccessTokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UserOAuthAccessTokenList{}
	err = c.client.Get().
		Resource("useroauthaccesstokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userOAuthAccessTokens.
func (c *userOAuthAccessTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("useroauthaccesstokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a userOAuthAccessToken and creates it.  Returns the server's representation of the userOAuthAccessToken, and an error, if there is any.
func (c *userOAuthAccessTokens) Create(ctx context.Context, userOAuthAccessToken *v1.UserOAuthAccessToken, opts metav1.CreateOptions) (result *v1.UserOAuthAccessToken, err error) {
	result = &v1.UserOAuthAccessToken{}
	err = c.client.Post().
		Resource("useroauthaccesstokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userOAuthAccessToken).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a userOAuthAccessToken and updates it. Returns the server's representation of the userOAuthAccessToken, and an error, if there is any.
func (c *userOAuthAccessTokens) Update(ctx context.Context, userOAuthAccessToken *v1.UserOAuthAccessToken, opts metav1.UpdateOptions) (result *v1.UserOAuthAccessToken, err error) {
	result = &v1.UserOAuthAccessToken{}
	err = c.client.Put().
		Resource("useroauthaccesstokens").
		Name(userOAuthAccessToken.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userOAuthAccessToken).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the userOAuthAccessToken and deletes it. Returns an error if one occurs.
func (c *userOAuthAccessTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("useroauthaccesstokens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userOAuthAccessTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("useroauthaccesstokens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched userOAuthAccessToken.
func (c *userOAuthAccessTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserOAuthAccessToken, err error) {
	result = &v1.UserOAuthAccessToken{}
	err = c.client.Patch(pt).
		Resource("useroauthaccesstokens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied userOAuthAccessToken.
func (c *userOAuthAccessTokens) Apply(ctx context.Context, userOAuthAccessToken *oauthv1.UserOAuthAccessTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.UserOAuthAccessToken, err error) {
	if userOAuthAccessToken == nil {
		return nil, fmt.Errorf("userOAuthAccessToken provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(userOAuthAccessToken)
	if err != nil {
		return nil, err
	}
	name := userOAuthAccessToken.Name
	if name == nil {
		return nil, fmt.Errorf("userOAuthAccessToken.Name must be provided to Apply")
	}
	result = &v1.UserOAuthAccessToken{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("useroauthaccesstokens").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
