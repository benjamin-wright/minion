package config

import (
	"fmt"
	"os"
)

// Config represents the current application configuration
type Config struct {
	LogLevel string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"{LogLevel: %s}",
		c.LogLevel,
	)
}

// Get an instance of the system configuration
func Get() Config {
	logLevel, exists := os.LookupEnv("LOG_LEVEL")
	if !exists {
		logLevel = "info"
	}

	return Config{
		LogLevel: logLevel,
	}
}
