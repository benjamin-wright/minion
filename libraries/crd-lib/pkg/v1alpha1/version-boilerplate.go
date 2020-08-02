package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto create a Version from another one
func (in *Version) DeepCopyInto(out *Version) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = VersionSpec{
		Resource: in.Spec.Resource,
		Pipeline: in.Spec.Pipeline,
		Version:  in.Spec.Version,
	}
}

// DeepCopyObject Create a copy of the current Version instance
func (in *Version) DeepCopyObject() runtime.Object {
	out := Version{}
	in.DeepCopyInto(&out)

	return &out
}

// DeepCopyObject Create a copy of the current VersionList instance
func (in *VersionList) DeepCopyObject() runtime.Object {
	out := VersionList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Version, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
