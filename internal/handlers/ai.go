package handlers

import (
	"net/http"
	"strings"

	"github.com/diegotrujillor/go-portfolio-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AISummarizeRequest struct {
	Text string `json:"text"`
}

type AISummarizeResponse struct {
	Summary string `json:"summary"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type AIHandler struct {
	Ollama *services.OllamaClient
}

func NewAIHandler(ollama *services.OllamaClient) *AIHandler {
	return &AIHandler{Ollama: ollama}
}

func (h *AIHandler) Summarize(c *gin.Context) {
	var req AISummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid JSON body"})
		return
	}

	req.Text = strings.TrimSpace(req.Text)
	if req.Text == "" {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "text is required"})
		return
	}

	// Basic guardrail to avoid huge payloads early on (tune later)
	if len(req.Text) > 20_000 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "text is too long (max 20000 chars)"})
		return
	}

	log.Info().
		Int("text_len", len(req.Text)).
		Msg("ai summarize requested")

	summary, err := h.Ollama.Summarize(c.Request.Context(), req.Text)
	if err != nil {
		log.Warn().Err(err).Msg("ai summarize failed")
		c.JSON(http.StatusBadGateway, ErrorResponse{Error: "LLM request failed"})
		return
	}

	c.JSON(http.StatusOK, AISummarizeResponse{Summary: summary})
}
