package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineSpec The specification for a Minion-CI pipeline CRD
type PipelineSpec struct {
	Resources []PipelineResource `json:"resources"`
	Steps     []PipelineStep     `json:"steps"`
}

type PipelineResource struct {
	Name    string `json:"name"`
	Trigger bool   `json:"trigger"`
}

type PipelineStep struct {
	Name     string `json:"name"`
	Resource string `json:"resource,omitempty"`
	Action   string `json:"action,omitempty"`
	Path     string `json:"path,omitempty"`
	Image    string `json:"image,omitempty"`
	Command  string `json:"command,omitempty"`
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

// PipelineList a list of sqs notifications
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items"`
}
