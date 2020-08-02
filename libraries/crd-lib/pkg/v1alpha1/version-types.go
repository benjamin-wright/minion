package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VersionSpec The specification for a Minion-CI resource CRD
type VersionSpec struct {
	Resource string `json:"resource,omitempty"`
	Pipeline string `json:"pipeline,omitempty"`
	Version  string `json:"version,omitempty"`
}

// Version A Minion-CI resource CRD object
type Version struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VersionSpec `json:"spec"`
}

// VersionList a list of sqs notifications
type VersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Version `json:"items"`
}
