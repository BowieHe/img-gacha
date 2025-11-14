package services

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/BowieHe/img-gacha/models"
	"github.com/google/uuid"
)

// ImageGenerator handles image generation logic
type ImageGenerator struct {
	geminiService *GeminiService
	openaiService *OpenAIService
}

// NewImageGenerator creates a new ImageGenerator instance
func NewImageGenerator(ctx context.Context) (*ImageGenerator, error) {
	geminiService, err := NewGeminiService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Gemini service: %w", err)
	}

	return &ImageGenerator{
		geminiService: geminiService,
		openaiService: NewOpenAIService(),
	}, nil
}

// GenerateImages creates multiple images based on the request
func (ig *ImageGenerator) GenerateImages(req *models.GenerationRequest) (*models.GenerationResponse, error) {
	taskID := uuid.New().String()

	// Route to appropriate service based on model
	var imageURLs []string
	var err error

	switch req.Model {
	case "gemini-2.0-flash", "gemini-image":
		imageURLs, err = ig.geminiService.GenerateImages(req.Prompt, req.Count)
	case "gpt-4-image":
		// First, use GPT-4 to enhance the prompt
		enhancedPrompt, err := ig.openaiService.GenerateImageDescription(req.Prompt)
		if err != nil {
			return nil, fmt.Errorf("failed to enhance prompt with GPT-4: %w", err)
		}

		// Then generate with the enhanced prompt using Gemini
		imageURLs, err = ig.geminiService.GenerateImages(enhancedPrompt, req.Count)
	default:
		return nil, fmt.Errorf("unsupported model: %s", req.Model)
	}

	if err != nil {
		return nil, err
	}

	resp := &models.GenerationResponse{
		TaskID:    taskID,
		Status:    "completed",
		Images:    imageURLs,
		CreatedAt: time.Now(),
	}

	return resp, nil
}

// GetRandomSeed generates a random seed for consistent but varying results
func GetRandomSeed() int64 {
	return rand.Int63()
}

// GenerateWithSeed generates images with a specific seed for reproducibility
func (ig *ImageGenerator) GenerateWithSeed(prompt string, model string, seed int64, count int) ([]models.ImageResult, error) {
	results := make([]models.ImageResult, 0, count)

	for i := 0; i < count; i++ {
		// Generate images with slight variations
		// Each image gets its own unique seed derived from the base seed
		uniqueSeed := seed + int64(i)

		// TODO: Call the actual model API
		img := models.ImageResult{
			ID:        uuid.New().String(),
			Seed:      uniqueSeed,
			CreatedAt: time.Now(),
		}

		results = append(results, img)
	}

	return results, nil
}

// RetryWithBackoff implements exponential backoff for failed requests
func RetryWithBackoff(maxRetries int, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err := fn()
		if err == nil {
			return nil
		}

		lastErr = err

		if attempt < maxRetries-1 {
			backoffDuration := time.Duration(1<<uint(attempt)) * time.Second
			time.Sleep(backoffDuration)
		}
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}
