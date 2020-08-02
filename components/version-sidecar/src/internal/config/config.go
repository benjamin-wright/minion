package config

import (
	"errors"
	"fmt"
	"os"
)

// Config represents the current application configuration
type Config struct {
	LogLevel string
	Pipeline string
	Resource string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"{LogLevel: %s, Pipeline: %s, Resource: %s}",
		c.LogLevel,
		c.Pipeline,
		c.Resource,
	)
}

// Get an instance of the system configuration
func Get() (Config, error) {
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "info"
	}

	pipeline, exists := os.LookupEnv("PIPELINE")
	if !exists {
		return Config{}, errors.New("PIPELINE environment variable not defined")
	}

	resource, exists := os.LookupEnv("RESOURCE")
	if !exists {
		return Config{}, errors.New("RESOURCE environment variable not defined")
	}

	return Config{
		LogLevel: logLevel,
		Pipeline: pipeline,
		Resource: resource,
	}, nil
}
