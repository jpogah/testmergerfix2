package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	defaultEnvironment          = "development"
	defaultPort                 = 8080
	defaultReadHeaderTimeoutSec = 5
	defaultShutdownTimeoutSec   = 10
	defaultHelloMessage         = "hello from hello-web"
)

// Config captures runtime settings for hello-web.
type Config struct {
	Environment       string
	Port              int
	ReadHeaderTimeout time.Duration
	ShutdownTimeout   time.Duration
	HelloMessage      string
}

func Load() (Config, error) {
	cfg := Config{
		Environment:       getEnv("APP_ENV", defaultEnvironment),
		Port:              getInt("HTTP_PORT", defaultPort),
		ReadHeaderTimeout: time.Duration(getInt("READ_HEADER_TIMEOUT_SECONDS", defaultReadHeaderTimeoutSec)) * time.Second,
		ShutdownTimeout:   time.Duration(getInt("SHUTDOWN_TIMEOUT_SECONDS", defaultShutdownTimeoutSec)) * time.Second,
		HelloMessage:      getEnv("HELLO_MESSAGE", defaultHelloMessage),
	}

	if cfg.Port < 1 || cfg.Port > 65535 {
		return Config{}, fmt.Errorf("invalid HTTP_PORT %d: must be between 1 and 65535", cfg.Port)
	}

	if cfg.ReadHeaderTimeout <= 0 {
		return Config{}, fmt.Errorf("READ_HEADER_TIMEOUT_SECONDS must be greater than zero")
	}

	if cfg.ShutdownTimeout <= 0 {
		return Config{}, fmt.Errorf("SHUTDOWN_TIMEOUT_SECONDS must be greater than zero")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}
	return value
}

func getInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
