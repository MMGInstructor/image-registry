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

// OAuthAuthorizeTokensGetter has a method to return a OAuthAuthorizeTokenInterface.
// A group's client should implement this interface.
type OAuthAuthorizeTokensGetter interface {
	OAuthAuthorizeTokens() OAuthAuthorizeTokenInterface
}

// OAuthAuthorizeTokenInterface has methods to work with OAuthAuthorizeToken resources.
type OAuthAuthorizeTokenInterface interface {
	Create(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.CreateOptions) (*v1.OAuthAuthorizeToken, error)
	Update(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.UpdateOptions) (*v1.OAuthAuthorizeToken, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.OAuthAuthorizeToken, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.OAuthAuthorizeTokenList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OAuthAuthorizeToken, err error)
	Apply(ctx context.Context, oAuthAuthorizeToken *oauthv1.OAuthAuthorizeTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OAuthAuthorizeToken, err error)
	OAuthAuthorizeTokenExpansion
}

// oAuthAuthorizeTokens implements OAuthAuthorizeTokenInterface
type oAuthAuthorizeTokens struct {
	client rest.Interface
}

// newOAuthAuthorizeTokens returns a OAuthAuthorizeTokens
func newOAuthAuthorizeTokens(c *OauthV1Client) *oAuthAuthorizeTokens {
	return &oAuthAuthorizeTokens{
		client: c.RESTClient(),
	}
}

// Get takes name of the oAuthAuthorizeToken, and returns the corresponding oAuthAuthorizeToken object, and an error if there is any.
func (c *oAuthAuthorizeTokens) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.OAuthAuthorizeToken, err error) {
	result = &v1.OAuthAuthorizeToken{}
	err = c.client.Get().
		Resource("oauthauthorizetokens").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OAuthAuthorizeTokens that match those selectors.
func (c *oAuthAuthorizeTokens) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OAuthAuthorizeTokenList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.OAuthAuthorizeTokenList{}
	err = c.client.Get().
		Resource("oauthauthorizetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oAuthAuthorizeTokens.
func (c *oAuthAuthorizeTokens) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("oauthauthorizetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a oAuthAuthorizeToken and creates it.  Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *oAuthAuthorizeTokens) Create(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.CreateOptions) (result *v1.OAuthAuthorizeToken, err error) {
	result = &v1.OAuthAuthorizeToken{}
	err = c.client.Post().
		Resource("oauthauthorizetokens").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oAuthAuthorizeToken).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a oAuthAuthorizeToken and updates it. Returns the server's representation of the oAuthAuthorizeToken, and an error, if there is any.
func (c *oAuthAuthorizeTokens) Update(ctx context.Context, oAuthAuthorizeToken *v1.OAuthAuthorizeToken, opts metav1.UpdateOptions) (result *v1.OAuthAuthorizeToken, err error) {
	result = &v1.OAuthAuthorizeToken{}
	err = c.client.Put().
		Resource("oauthauthorizetokens").
		Name(oAuthAuthorizeToken.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(oAuthAuthorizeToken).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the oAuthAuthorizeToken and deletes it. Returns an error if one occurs.
func (c *oAuthAuthorizeTokens) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("oauthauthorizetokens").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oAuthAuthorizeTokens) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("oauthauthorizetokens").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched oAuthAuthorizeToken.
func (c *oAuthAuthorizeTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.OAuthAuthorizeToken, err error) {
	result = &v1.OAuthAuthorizeToken{}
	err = c.client.Patch(pt).
		Resource("oauthauthorizetokens").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied oAuthAuthorizeToken.
func (c *oAuthAuthorizeTokens) Apply(ctx context.Context, oAuthAuthorizeToken *oauthv1.OAuthAuthorizeTokenApplyConfiguration, opts metav1.ApplyOptions) (result *v1.OAuthAuthorizeToken, err error) {
	if oAuthAuthorizeToken == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(oAuthAuthorizeToken)
	if err != nil {
		return nil, err
	}
	name := oAuthAuthorizeToken.Name
	if name == nil {
		return nil, fmt.Errorf("oAuthAuthorizeToken.Name must be provided to Apply")
	}
	result = &v1.OAuthAuthorizeToken{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("oauthauthorizetokens").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
