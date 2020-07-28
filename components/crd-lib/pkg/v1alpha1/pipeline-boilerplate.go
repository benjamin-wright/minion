package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
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
