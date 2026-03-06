package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type AiSummarizeClient struct {
	llmBaseUrl string
	llmModel   string
	timeout    time.Duration
}

func NewAiSummarizeClient(baseUrl, model string, timeout time.Duration) *AiSummarizeClient {
	return &AiSummarizeClient{
		llmBaseUrl: baseUrl,
		llmModel:   model,
		timeout:    timeout,
	}
}

func (c *AiSummarizeClient) Summarize(text string) (Summary, error) {
	req, err := http.NewRequest("POST", c.llmBaseUrl, strings.NewReader(text))
	if err != nil {
		return Summary{}, err
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Model", c.llmModel)

	client := &http.Client{
		Timeout: c.timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return Summary{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return Summary{}, fmt.Errorf("LLM call failed: %d", resp.StatusCode)
	}

	var summary Summary
	err = json.NewDecoder(resp.Body).Decode(&summary)
	if err != nil {
		return Summary{}, err
	}

	return summary, nil
}

type Summary struct {
	Summary string `json:"summary"`
}
