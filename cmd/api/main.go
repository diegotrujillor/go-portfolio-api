package main

import (
	"net/http"

	"github.com/diegotrujillor/go-portfolio-api/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}
	log.Info().Msgf("Loaded configuration: %+v", cfg)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		log.Info().Msg("Received ping request")
		c.JSON(http.StatusOK, gin.H{
			"message": "Server listening on :8080",
		})
	})

	// Create HTTP server with config values
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	}

	log.Info().Str("port", cfg.Port).Str("env", cfg.Env).Msg("Starting API server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start API server")
	}
}
