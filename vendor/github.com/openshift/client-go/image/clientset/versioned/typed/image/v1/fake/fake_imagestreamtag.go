// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	imagev1 "github.com/openshift/api/image/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// FakeImageStreamTags implements ImageStreamTagInterface
type FakeImageStreamTags struct {
	Fake *FakeImageV1
	ns   string
}

var imagestreamtagsResource = schema.GroupVersionResource{Group: "image.uccp.io", Version: "v1", Resource: "imagestreamtags"}

var imagestreamtagsKind = schema.GroupVersionKind{Group: "image.uccp.io", Version: "v1", Kind: "ImageStreamTag"}

// Get takes name of the imageStreamTag, and returns the corresponding imageStreamTag object, and an error if there is any.
func (c *FakeImageStreamTags) Get(ctx context.Context, name string, options v1.GetOptions) (result *imagev1.ImageStreamTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagestreamtagsResource, c.ns, name), &imagev1.ImageStreamTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageStreamTag), err
}

// List takes label and field selectors, and returns the list of ImageStreamTags that match those selectors.
func (c *FakeImageStreamTags) List(ctx context.Context, opts v1.ListOptions) (result *imagev1.ImageStreamTagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagestreamtagsResource, imagestreamtagsKind, c.ns, opts), &imagev1.ImageStreamTagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &imagev1.ImageStreamTagList{ListMeta: obj.(*imagev1.ImageStreamTagList).ListMeta}
	for _, item := range obj.(*imagev1.ImageStreamTagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a imageStreamTag and creates it.  Returns the server's representation of the imageStreamTag, and an error, if there is any.
func (c *FakeImageStreamTags) Create(ctx context.Context, imageStreamTag *imagev1.ImageStreamTag, opts v1.CreateOptions) (result *imagev1.ImageStreamTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagestreamtagsResource, c.ns, imageStreamTag), &imagev1.ImageStreamTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageStreamTag), err
}

// Update takes the representation of a imageStreamTag and updates it. Returns the server's representation of the imageStreamTag, and an error, if there is any.
func (c *FakeImageStreamTags) Update(ctx context.Context, imageStreamTag *imagev1.ImageStreamTag, opts v1.UpdateOptions) (result *imagev1.ImageStreamTag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagestreamtagsResource, c.ns, imageStreamTag), &imagev1.ImageStreamTag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*imagev1.ImageStreamTag), err
}

// Delete takes name of the imageStreamTag and deletes it. Returns an error if one occurs.
func (c *FakeImageStreamTags) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(imagestreamtagsResource, c.ns, name, opts), &imagev1.ImageStreamTag{})

	return err
}
