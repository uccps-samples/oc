// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/openshift/api/image/v1"
	"github.com/openshift/client-go/image/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type ImageV1Interface interface {
	RESTClient() rest.Interface
	ImagesGetter
	ImageSignaturesGetter
	ImageStreamsGetter
	ImageStreamImagesGetter
	ImageStreamImportsGetter
	ImageStreamMappingsGetter
	ImageStreamTagsGetter
	ImageTagsGetter
}

// ImageV1Client is used to interact with features provided by the image.uccp.io group.
type ImageV1Client struct {
	restClient rest.Interface
}

func (c *ImageV1Client) Images() ImageInterface {
	return newImages(c)
}

func (c *ImageV1Client) ImageSignatures() ImageSignatureInterface {
	return newImageSignatures(c)
}

func (c *ImageV1Client) ImageStreams(namespace string) ImageStreamInterface {
	return newImageStreams(c, namespace)
}

func (c *ImageV1Client) ImageStreamImages(namespace string) ImageStreamImageInterface {
	return newImageStreamImages(c, namespace)
}

func (c *ImageV1Client) ImageStreamImports(namespace string) ImageStreamImportInterface {
	return newImageStreamImports(c, namespace)
}

func (c *ImageV1Client) ImageStreamMappings(namespace string) ImageStreamMappingInterface {
	return newImageStreamMappings(c, namespace)
}

func (c *ImageV1Client) ImageStreamTags(namespace string) ImageStreamTagInterface {
	return newImageStreamTags(c, namespace)
}

func (c *ImageV1Client) ImageTags(namespace string) ImageTagInterface {
	return newImageTags(c, namespace)
}

// NewForConfig creates a new ImageV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ImageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ImageV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ImageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ImageV1Client{client}, nil
}

// NewForConfigOrDie creates a new ImageV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ImageV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ImageV1Client for the given RESTClient.
func New(c rest.Interface) *ImageV1Client {
	return &ImageV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ImageV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
