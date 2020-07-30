package informer

import "ponglehub.co.uk/crd-lib/pkg/v1alpha1"

const (
	ADDED   string = "ADDED"
	UPDATED string = "UPDATED"
	DELETED string = "DELETED"
)

type PipelineEvent struct {
	Kind     string
	Current  *v1alpha1.Pipeline
	Previous *v1alpha1.Pipeline
}

type PipelineInformer struct {
	Events  <-chan PipelineEvent
	Stopper chan<- struct{}
}
