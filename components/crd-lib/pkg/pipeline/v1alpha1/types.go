package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineSpec The specification for a Minion-CI pipeline CRD
type PipelineSpec struct {
	Queue  string `json:"queue,omitempty"`
	URL    string `json:"url,omitempty"`
	Filter string `json:"filter,omitempty"`
}

// PipelineStatus The status for a Minion-CI pipeline CRD
type PipelineStatus struct {
	Status string `json:"status,omitempty"`
}

// Pipeline A Minion-CI pipeline CRD object
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PipelineSpec   `json:"spec"`
	Status            PipelineStatus `json:"status"`
}
