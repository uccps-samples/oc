// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/uccps-samples/api/authorization/v1"
	scheme "github.com/uccps-samples/client-go/authorization/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// LocalSubjectAccessReviewsGetter has a method to return a LocalSubjectAccessReviewInterface.
// A group's client should implement this interface.
type LocalSubjectAccessReviewsGetter interface {
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewInterface
}

// LocalSubjectAccessReviewInterface has methods to work with LocalSubjectAccessReview resources.
type LocalSubjectAccessReviewInterface interface {
	Create(ctx context.Context, localSubjectAccessReview *v1.LocalSubjectAccessReview, opts metav1.CreateOptions) (*v1.SubjectAccessReviewResponse, error)

	LocalSubjectAccessReviewExpansion
}

// localSubjectAccessReviews implements LocalSubjectAccessReviewInterface
type localSubjectAccessReviews struct {
	client rest.Interface
	ns     string
}

// newLocalSubjectAccessReviews returns a LocalSubjectAccessReviews
func newLocalSubjectAccessReviews(c *AuthorizationV1Client, namespace string) *localSubjectAccessReviews {
	return &localSubjectAccessReviews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a localSubjectAccessReview and creates it.  Returns the server's representation of the subjectAccessReviewResponse, and an error, if there is any.
func (c *localSubjectAccessReviews) Create(ctx context.Context, localSubjectAccessReview *v1.LocalSubjectAccessReview, opts metav1.CreateOptions) (result *v1.SubjectAccessReviewResponse, err error) {
	result = &v1.SubjectAccessReviewResponse{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("localsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}