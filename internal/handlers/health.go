package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Health(c *gin.Context) {
	// Keep this low-noise in production; switch to Info() if want it always visible
	log.Debug().Msg("healthcheck requested")

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
