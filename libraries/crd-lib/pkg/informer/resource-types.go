package informer

import "ponglehub.co.uk/crd-lib/pkg/v1alpha1"

type ResourceEvent struct {
	Kind     EventType
	Current  *v1alpha1.Resource
	Previous *v1alpha1.Resource
}

type ResourceInformer struct {
	Events  <-chan ResourceEvent
	Stopper chan<- struct{}
}
