package informer

import "ponglehub.co.uk/crd-lib/pkg/v1alpha1"

type PipelineEvent struct {
	Kind     EventType
	Current  *v1alpha1.Pipeline
	Previous *v1alpha1.Pipeline
}

type PipelineInformer struct {
	Events  <-chan PipelineEvent
	Stopper chan<- struct{}
}
