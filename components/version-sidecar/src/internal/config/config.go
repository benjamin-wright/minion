package config

import (
	"errors"
	"fmt"
	"os"
)

// Config represents the current application configuration
type Config struct {
	LogLevel string
	Resource string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"{LogLevel: %s, Resource: %s}",
		c.LogLevel,
		c.Resource,
	)
}

// Get an instance of the system configuration
func Get() (Config, error) {
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "info"
	}

	resource, exists := os.LookupEnv("RESOURCE")
	if !exists {
		return Config{}, errors.New("RESOURCE environment variable not defined")
	}

	return Config{
		LogLevel: logLevel,
		Resource: resource,
	}, nil
}
