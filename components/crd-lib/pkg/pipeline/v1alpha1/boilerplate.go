package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// DeepCopyInto create a Pipeline from another one
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = PipelineSpec{
		Queue: in.Spec.Queue,
		URL:   in.Spec.URL,
	}
	out.Status = PipelineStatus{
		Status: in.Status.Status,
	}
}

// DeepCopyObject Create a copy of the current Pipeline instance
func (in *Pipeline) DeepCopyObject() runtime.Object {
	out := Pipeline{}
	in.DeepCopyInto(&out)

	return &out
}

// PipelineList a list of sqs notifications
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}

// DeepCopyObject Create a copy of the current PipelineList instance
func (in *PipelineList) DeepCopyObject() runtime.Object {
	out := PipelineList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Pipeline, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}

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

// NewClient return a new versioned REST client for accessing CRD resources
func NewClient() (*rest.RESTClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	addToScheme(scheme.Scheme)

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &schema.GroupVersion{Group: groupName, Version: groupVersion}
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	return rest.UnversionedRESTClientFor(&crdConfig)
}
