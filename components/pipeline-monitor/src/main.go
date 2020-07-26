package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	for {
		log.Info("tick...")
		time.Sleep(10 * time.Second)
	}
}
