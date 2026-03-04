package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Env          string
	Port         string
	LogLevel     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		Env:      getEnv("ENV", "local"),
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	var err error
	cfg.ReadTimeout, err = time.ParseDuration(getEnv("READ_TIMEOUT", "5s"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid READ_TIMEOUT: %w", err)
	}
	cfg.WriteTimeout, err = time.ParseDuration(getEnv("WRITE_TIMEOUT", "10s"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid WRITE_TIMEOUT: %w", err)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
