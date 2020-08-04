package config

import (
	"fmt"
	"os"
)

// Config represents the current application configuration
type Config struct {
	LogLevel     string
	SidecarImage string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"{LogLevel: %s, SidecarImage: %s}",
		c.LogLevel,
		c.SidecarImage,
	)
}

// Get an instance of the system configuration
func Get() Config {
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "info"
	}

	sidecarImage, exists := os.LookupEnv("SIDECAR_IMAGE")
	if !exists {
		logLevel = "info"
	}

	return Config{
		LogLevel:     logLevel,
		SidecarImage: sidecarImage,
	}
}
