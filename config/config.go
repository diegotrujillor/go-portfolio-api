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

	// AI / Ollama
	LLMBaseURL string
	LLMModel   string
	LLMTimeout time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		Env:        getEnv("ENV", "local"),
		Port:       getEnv("PORT", "8080"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		LLMBaseURL: getEnv("LLM_BASE_URL", "http://localhost:11434"),
		LLMModel:   getEnv("LLM_MODEL", "llama3.1"),
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
	cfg.LLMTimeout, err = time.ParseDuration(getEnv("LLM_TIMEOUT", "20s"))
	if err != nil {
		return Config{}, fmt.Errorf("invalid LLM_TIMEOUT: %w", err)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
