package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto create a Resource from another one
func (in *Resource) DeepCopyInto(out *Resource) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ResourceSpec{
		Image:   in.Spec.Image,
		Env:     map[string]string{},
		Secrets: append([]Secret{}, in.Spec.Secrets...),
	}

	for k, v := range in.Spec.Env {
		out.Spec.Env[k] = v
	}

	out.Status = ResourceStatus{
		Status: in.Status.Status,
	}
}

// DeepCopyObject Create a copy of the current Resource instance
func (in *Resource) DeepCopyObject() runtime.Object {
	out := Resource{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject Create a copy of the current ResourceList instance
func (in *ResourceList) DeepCopyObject() runtime.Object {
	out := ResourceList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Resource, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
