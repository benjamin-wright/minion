package client

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	v1alpha1 "ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

// GetResource fetch a resource CRD definition
func (c MinionCRDClient) GetResource(name string, namespace string) (*v1alpha1.Resource, error) {
	existing := v1alpha1.Resource{}
	err := c.client.
		Get().
		Namespace(namespace).
		Resource("resources").
		Name(name).
		Do().
		Into(&existing)

	return &existing, err
}

// ListResources get a list of resource CRD definitions
func (c MinionCRDClient) ListResources(options metav1.ListOptions) (*v1alpha1.ResourceList, error) {
	result := v1alpha1.ResourceList{}
	err := c.client.
		Get().
		Resource("resources").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

// WatchResources get informed when the list of resource CRDs changes
func (c MinionCRDClient) WatchResources(options metav1.ListOptions) (watch.Interface, error) {
	options.Watch = true

	return c.client.
		Get().
		Resource("resources").
		Timeout(time.Second*20).
		VersionedParams(&options, scheme.ParameterCodec).
		Watch()
}
