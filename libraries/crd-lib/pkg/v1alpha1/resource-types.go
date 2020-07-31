package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceSpec The specification for a Minion-CI resource CRD
type ResourceSpec struct {
	Image   string            `json:"image,omitempty"`
	Env     map[string]string `json:"env,omitempty"`
	Secrets []Secret          `json:"secrets,omitempty"`
}

// ResourceStatus The status for a Minion-CI resource CRD
type ResourceStatus struct {
	Status string `json:"status,omitempty"`
}

// Resource A Minion-CI resource CRD object
type Resource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceSpec   `json:"spec"`
	Status            ResourceStatus `json:"status"`
}

// ResourceList a list of sqs notifications
type ResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Resource `json:"items"`
}
