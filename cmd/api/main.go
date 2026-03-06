package main

import (
	"net/http"

	"github.com/diegotrujillor/go-portfolio-api/config"
	"github.com/diegotrujillor/go-portfolio-api/internal/handlers"
	"github.com/diegotrujillor/go-portfolio-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load configuration")
	}

	ollamaClient, err := services.NewOllamaClient(cfg.LLMBaseURL, cfg.LLMModel, cfg.LLMTimeout)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to init ollama client")
	}
	aiHandler := handlers.NewAIHandler(ollamaClient)

	router := gin.Default()

	// health
	router.GET("/health", handlers.Health)

	// AI
	router.POST("/ai/summarize", aiHandler.Summarize)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Info().
		Str("port", cfg.Port).
		Str("env", cfg.Env).
		Str("llm_base_url", cfg.LLMBaseURL).
		Str("llm_model", cfg.LLMModel).
		Msg("starting API server")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("failed to start API server")
	}
}
