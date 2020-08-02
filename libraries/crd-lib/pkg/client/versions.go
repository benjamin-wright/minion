package client

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	v1alpha1 "ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

func (c MinionCRDClient) PostVersion(version *v1alpha1.Version, namespace string) error {
	return c.client.
		Post().
		Namespace(namespace).
		Resource("versions").
		VersionedParams(&metav1.CreateOptions{}, scheme.ParameterCodec).
		Body(version).
		Do().
		Error()
}

// GetVersion fetch a pipeline CRD definition
func (c MinionCRDClient) GetVersion(name string, namespace string) (*v1alpha1.Version, error) {
	existing := v1alpha1.Version{}
	err := c.client.
		Get().
		Namespace(namespace).
		Resource("versions").
		Name(name).
		Do().
		Into(&existing)

	return &existing, err
}

// ListVersions get a list of pipeline CRD definitions
func (c MinionCRDClient) ListVersions(options metav1.ListOptions) (*v1alpha1.VersionList, error) {
	result := v1alpha1.VersionList{}
	err := c.client.
		Get().
		Resource("versions").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

// WatchVersions get informed when the list of pipeline CRDs changes
func (c MinionCRDClient) WatchVersions(options metav1.ListOptions) (watch.Interface, error) {
	options.Watch = true

	return c.client.
		Get().
		Resource("versions").
		Timeout(time.Second*20).
		VersionedParams(&options, scheme.ParameterCodec).
		Watch()
}
