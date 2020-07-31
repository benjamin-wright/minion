package informer

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"ponglehub.co.uk/crd-lib/pkg/client"
	"ponglehub.co.uk/crd-lib/pkg/v1alpha1"
)

type resourceListerWatcher struct {
	client *client.MinionCRDClient
}

func (w *resourceListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	return w.client.ListResources(options)
}

func (w *resourceListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	return w.client.WatchResources(options)
}

// Resources returns a channel into which new mongo user events will be pushed
func Resources() (*ResourceInformer, error) {
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	events := make(chan ResourceEvent, 20)
	stopper := make(chan struct{})

	_, informer := cache.NewIndexerInformer(
		&resourceListerWatcher{client: &client},
		&v1alpha1.Resource{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    addResourceFunc(events),
			UpdateFunc: updateResourceFunc(events),
			DeleteFunc: deleteResourceFunc(events),
		},
		cache.Indexers{},
	)

	go informer.Run(stopper)

	return &ResourceInformer{
		Events:  events,
		Stopper: stopper,
	}, nil
}

func addResourceFunc(events chan<- ResourceEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- ResourceEvent{
			Kind:    ADDED,
			Current: obj.(*v1alpha1.Resource),
		}
	}
}

func updateResourceFunc(events chan<- ResourceEvent) func(interface{}, interface{}) {
	return func(obj1 interface{}, obj2 interface{}) {
		events <- ResourceEvent{
			Kind:     UPDATED,
			Current:  obj2.(*v1alpha1.Resource),
			Previous: obj1.(*v1alpha1.Resource),
		}
	}
}

func deleteResourceFunc(events chan<- ResourceEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- ResourceEvent{
			Kind:     DELETED,
			Previous: obj.(*v1alpha1.Resource),
		}
	}
}
