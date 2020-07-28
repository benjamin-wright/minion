package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

const groupName = "minion.ponglehub.co.uk"
const groupVersion = "v1alpha1"

var schemeGroupVersion = schema.GroupVersion{Group: groupName, Version: groupVersion}

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	addToScheme   = schemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		schemeGroupVersion,
		&Pipeline{},
		&PipelineList{},
	)

	scheme.AddKnownTypes(
		schema.GroupVersion{Group: groupName, Version: "__internal"},
		&Pipeline{},
		&PipelineList{},
	)

	metav1.AddToGroupVersion(scheme, schemeGroupVersion)
	return nil
}

// Init configure the kubernetes environment with details of the CRD
func Init() error {
	return addToScheme(scheme.Scheme)
}

// TypedClient return a new versioned REST client for accessing CRD resources
func TypedClient() (*rest.RESTClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: groupName, Version: groupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	return rest.UnversionedRESTClientFor(&crdConfig)
}
