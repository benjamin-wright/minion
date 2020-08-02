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

type versionListerWatcher struct {
	client *client.MinionCRDClient
}

func (w *versionListerWatcher) List(options metav1.ListOptions) (runtime.Object, error) {
	return w.client.ListVersions(options)
}

func (w *versionListerWatcher) Watch(options metav1.ListOptions) (watch.Interface, error) {
	return w.client.WatchVersions(options)
}

// Versions returns a channel into which new mongo user events will be pushed
func Versions() (*VersionInformer, error) {
	client, err := client.New()
	if err != nil {
		return nil, err
	}

	events := make(chan VersionEvent, 20)
	stopper := make(chan struct{})

	_, informer := cache.NewIndexerInformer(
		&versionListerWatcher{client: &client},
		&v1alpha1.Version{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    addVersionFunc(events),
			UpdateFunc: updateVersionFunc(events),
			DeleteFunc: deleteVersionFunc(events),
		},
		cache.Indexers{},
	)

	go informer.Run(stopper)

	return &VersionInformer{
		Events:  events,
		Stopper: stopper,
	}, nil
}

func addVersionFunc(events chan<- VersionEvent) func(interface{}) {
	return func(obj interface{}) {
		events <- VersionEvent{
			Kind:    ADDED,
			Current: obj.(*v1alpha1.Version),
		}
	}
}

func updateVersionFunc(events chan<- VersionEvent) func(interface{}, interface{}) {
	return func(obj1 interface{}, obj2 interface{}) {
		logrus.Info("Updated event")

		events <- VersionEvent{
			Kind:     UPDATED,
			Current:  obj2.(*v1alpha1.Version),
			Previous: obj1.(*v1alpha1.Version),
		}
	}
}

func deleteVersionFunc(events chan<- VersionEvent) func(interface{}) {
	return func(obj interface{}) {
		logrus.Info("Deleted event")

		events <- VersionEvent{
			Kind:     DELETED,
			Previous: obj.(*v1alpha1.Version),
		}
	}
}
