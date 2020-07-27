package main

import (
	"time"

	"ponglehub.co.uk/crd-lib/pkg/pipeline/v1alpha1"
	log "github.com/sirupsen/logrus"
)

func main() {
	for {
		spec := v1alpha1.PipelineSpec{}
		log.Infof("%+v", spec)
		log.Info("tick...")
		time.Sleep(10 * time.Second)
	}
}
