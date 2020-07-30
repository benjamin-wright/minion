package informer

import (
	"time"

	"k8s.io/client-go/tools/cache"
	"ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

// Pipelines returns a channel into which new mongo user events will be pushed
func Pipelines() (*PipelineInformer, error) {
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	events := make(chan PipelineEvent, 20)
	stopper := make(chan struct{})

	_, informer := cache.NewIndexerInformer(
		&pipelineListerWatcher{client: &client},
		&v1alpha1.Pipeline{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    addFunc(events),
			UpdateFunc: updateFunc(events),
			DeleteFunc: deleteFunc(events),
		},
		cache.Indexers{},
	)

	go informer.Run(stopper)

	return &PipelineInformer{
		Events:  events,
		Stopper: stopper,
	}, nil
}

func addFunc(events chan<- PipelineEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- PipelineEvent{
			Kind:    ADDED,
			Current: obj.(*v1alpha1.Pipeline),
		}
	}
}

func updateFunc(events chan<- PipelineEvent) func(interface{}, interface{}) {
	return func(obj1 interface{}, obj2 interface{}) {
		events <- PipelineEvent{
			Kind:     UPDATED,
			Current:  obj2.(*v1alpha1.Pipeline),
			Previous: obj1.(*v1alpha1.Pipeline),
		}
	}
}

func deleteFunc(events chan<- PipelineEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- PipelineEvent{
			Kind:     DELETED,
			Previous: obj.(*v1alpha1.Pipeline),
		}
	}
}
