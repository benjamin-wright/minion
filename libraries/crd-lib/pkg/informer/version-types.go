package informer

import "ponglehub.co.uk/crd-lib/pkg/v1alpha1"

type VersionEvent struct {
	Kind     EventType
	Current  *v1alpha1.Version
	Previous *v1alpha1.Version
}

type VersionInformer struct {
	Events  <-chan VersionEvent
	Stopper chan<- struct{}
}
