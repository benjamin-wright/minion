package listener

import (
	log "github.com/sirupsen/logrus"
	crdInformer "ponglehub.co.uk/crd-lib/pkg/informer"
	"ponglehub.co.uk/resource-monitor/internal/config"
)

// Listen responds to resource events
func (l *Listener) Listen(events <-chan crdInformer.ResourceEvent, cfg config.Config) {
	for event := range events {
		switch event.Kind {
		case crdInformer.ADDED:
			err := l.m.Create(event.Current, cfg)
			if err != nil {
				log.Errorf("Failed to add resource: %+v", err)
			} else {
				log.Infof("Resource added: %s", event.Current.ObjectMeta.Name)
			}
		case crdInformer.UPDATED:
			log.Info("Resource Updated")
		case crdInformer.DELETED:
			log.Info("Resource Deleted")
		}
	}
}
