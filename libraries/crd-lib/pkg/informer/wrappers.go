package informer

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"ponglehub.co.uk/crd-lib/pkg/client"
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
