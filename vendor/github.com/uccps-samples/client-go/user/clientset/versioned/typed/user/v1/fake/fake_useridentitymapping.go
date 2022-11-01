// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	userv1 "github.com/uccps-samples/api/user/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeUserIdentityMappings implements UserIdentityMappingInterface
type FakeUserIdentityMappings struct {
	Fake *FakeUserV1
}

var useridentitymappingsResource = schema.GroupVersionResource{Group: "user.uccp.io", Version: "v1", Resource: "useridentitymappings"}

var useridentitymappingsKind = schema.GroupVersionKind{Group: "user.uccp.io", Version: "v1", Kind: "UserIdentityMapping"}

// Get takes name of the userIdentityMapping, and returns the corresponding userIdentityMapping object, and an error if there is any.
func (c *FakeUserIdentityMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *userv1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(useridentitymappingsResource, name), &userv1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.UserIdentityMapping), err
}

// Create takes the representation of a userIdentityMapping and creates it.  Returns the server's representation of the userIdentityMapping, and an error, if there is any.
func (c *FakeUserIdentityMappings) Create(ctx context.Context, userIdentityMapping *userv1.UserIdentityMapping, opts v1.CreateOptions) (result *userv1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(useridentitymappingsResource, userIdentityMapping), &userv1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.UserIdentityMapping), err
}

// Update takes the representation of a userIdentityMapping and updates it. Returns the server's representation of the userIdentityMapping, and an error, if there is any.
func (c *FakeUserIdentityMappings) Update(ctx context.Context, userIdentityMapping *userv1.UserIdentityMapping, opts v1.UpdateOptions) (result *userv1.UserIdentityMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(useridentitymappingsResource, userIdentityMapping), &userv1.UserIdentityMapping{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.UserIdentityMapping), err
}

// Delete takes name of the userIdentityMapping and deletes it. Returns an error if one occurs.
func (c *FakeUserIdentityMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(useridentitymappingsResource, name, opts), &userv1.UserIdentityMapping{})
	return err
}
