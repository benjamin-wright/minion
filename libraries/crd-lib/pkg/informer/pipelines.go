package informer

import (
	"time"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

type pipelineListerWatcher struct {
	client *client.MinionCRDClient
}

func (w *pipelineListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	return w.client.ListPipelines(options)
}

func (w *pipelineListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	return w.client.WatchPipelines(options)
}

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
			AddFunc:    addPipelineFunc(events),
			UpdateFunc: updatePipelineFunc(events),
			DeleteFunc: deletePipelineFunc(events),
		},
		cache.Indexers{},
	)

	go informer.Run(stopper)

	return &PipelineInformer{
		Events:  events,
		Stopper: stopper,
	}, nil
}

func addPipelineFunc(events chan<- PipelineEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- PipelineEvent{
			Kind:    ADDED,
			Current: obj.(*v1alpha1.Pipeline),
		}
	}
}

func updatePipelineFunc(events chan<- PipelineEvent) func(interface{}, interface{}) {
	return func(obj1 interface{}, obj2 interface{}) {
		logrus.Info("Updated event")

		events <- PipelineEvent{
			Kind:     UPDATED,
			Current:  obj2.(*v1alpha1.Pipeline),
			Previous: obj1.(*v1alpha1.Pipeline),
		}
	}
}

func deletePipelineFunc(events chan<- PipelineEvent) func(interface{}) {
	return func(obj interface{}) {
		logrus.Info("Deleted event")

		events <- PipelineEvent{
			Kind:     DELETED,
			Previous: obj.(*v1alpha1.Pipeline),
		}
	}
}
