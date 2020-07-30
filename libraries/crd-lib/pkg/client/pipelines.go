package client

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	v1alpha1 "ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

// GetPipeline fetch a pipeline CRD definition
func (c MinionCRDClient) GetPipeline(name string, namespace string) (*v1alpha1.Pipeline, error) {
	existing := v1alpha1.Pipeline{}
	err := c.client.
		Get().
		Namespace(namespace).
		Resource("pipelines").
		Name(name).
		Do().
		Into(&existing)

	return &existing, err
}

// ListPipelines get a list of pipeline CRD definitions
func (c MinionCRDClient) ListPipelines(options metav1.ListOptions) (*v1alpha1.PipelineList, error) {
	result := v1alpha1.PipelineList{}
	err := c.client.
		Get().
		Resource("pipelines").
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(&result)

	return &result, err
}

// Watch get informed when the list of pipeline CRDs changes
func (c MinionCRDClient) Watch(options metav1.ListOptions) (watch.Interface, error) {
	options.Watch = true

	return c.client.
		Get().
		Resource("pipelines").
		Timeout(time.Second*20).
		VersionedParams(&options, scheme.ParameterCodec).
		Watch()
}
