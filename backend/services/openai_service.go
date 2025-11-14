package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// OpenAIService handles OpenAI API calls for GPT-4 Vision image generation
type OpenAIService struct {
	apiKey string
	client *http.Client
}

// OpenAIRequest represents the request payload for OpenAI API
type OpenAIRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	MaxTokens   int             `json:"max_tokens"`
	Temperature float32         `json:"temperature"`
}

// OpenAIMessage represents a message in the request
type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response from OpenAI API
type OpenAIResponse struct {
	ID      string         `json:"id"`
	Choices []OpenAIChoice `json:"choices"`
	Error   *OpenAIError   `json:"error,omitempty"`
}

// OpenAIChoice represents a choice in the response
type OpenAIChoice struct {
	Message OpenAIMessage `json:"message"`
	Index   int           `json:"index"`
}

// OpenAIError represents an error in the response
type OpenAIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// NewOpenAIService creates a new OpenAI service
func NewOpenAIService() *OpenAIService {
	return &OpenAIService{
		apiKey: os.Getenv("OPENAI_API_KEY"),
		client: &http.Client{},
	}
}

// GenerateImageDescription generates a detailed image description using GPT-4
// This can be used to create better prompts or direct image generation
func (os *OpenAIService) GenerateImageDescription(prompt string) (string, error) {
	if os.apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY not set")
	}

	openaiReq := OpenAIRequest{
		Model: "gpt-4-turbo",
		Messages: []OpenAIMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf("Create a detailed and vivid image generation prompt based on this description: %s. The prompt should be suitable for Stable Diffusion and include style, lighting, and composition details.", prompt),
			},
		},
		MaxTokens:   500,
		Temperature: 0.7,
	}

	body, err := json.Marshal(openaiReq)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.apiKey))

	resp, err := os.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to call OpenAI API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("OpenAI API error: status %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var openaiResp OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&openaiResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if openaiResp.Error != nil {
		return "", fmt.Errorf("OpenAI error: %s", openaiResp.Error.Message)
	}

	if len(openaiResp.Choices) == 0 {
		return "", fmt.Errorf("no choices in response")
	}

	return openaiResp.Choices[0].Message.Content, nil
}
