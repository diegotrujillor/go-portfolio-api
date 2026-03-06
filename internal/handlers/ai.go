package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/diegotrujillor/go-portfolio-api/internal/services"
	"github.com/diegotrujillor/go-portfolio-api/internal/services/ai_summarize"
)

type AiHandler struct{}

func (h AiHandler) Summarize(c *gin.Context) {
	req := &ai_summarize.Request{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	summary := &ai_summarize.Summary{}
	if err := c.BindJSON(summary); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid response"})
		return
	}

	log.Info().Str("text_length", req.Text).Msg("Received summary request")

	llmClient := services.NewAiSummarizeClient(cfg.LLM_BASE_URL, cfg.LLM_MODEL, cfg.AITimeout)
	resp, err := llmClient.Summarize(req.Text)
	if err != nil {
		log.Error().Err(err).Msg("LLM call failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "LLM call failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"summary": resp.Summary})
}
