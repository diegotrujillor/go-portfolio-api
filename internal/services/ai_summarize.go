package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type OllamaClient struct {
	BaseURL    string
	Model      string
	HTTPClient *http.Client
}

func NewOllamaClient(baseURL, model string, timeout time.Duration) (*OllamaClient, error) {
	baseURL = strings.TrimRight(baseURL, "/")
	if baseURL == "" {
		return nil, errors.New("LLM base URL is empty")
	}
	if model == "" {
		return nil, errors.New("LLM model is empty")
	}

	return &OllamaClient{
		BaseURL: baseURL,
		Model:   model,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

type ollamaGenerateRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaGenerateResponse struct {
	Response string `json:"response"`
	// Other fields exist (created_at, done, etc.) but don't needed now.
}

func (c *OllamaClient) Summarize(ctx context.Context, text string) (string, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return "", errors.New("text is empty")
	}

	// Keep the instruction stable for eval
	prompt := fmt.Sprintf(
		`Summarize the following text in 5 bullet points. Keep it factual, avoid assumptions, and do not invent details.

TEXT:
%s`, text)

	reqBody, err := json.Marshal(ollamaGenerateRequest{
		Model:  c.Model,
		Prompt: prompt,
		Stream: false,
	})
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	url := c.BaseURL + "/api/generate"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", fmt.Errorf("ollama returned non-2xx: %s", resp.Status)
	}

	var out ollamaGenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	summary := strings.TrimSpace(out.Response)
	if summary == "" {
		return "", errors.New("empty summary from model")
	}

	return summary, nil
}
