package services

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// GeminiService handles Google Gemini API calls for image generation
type GeminiService struct {
	client *genai.Client
	ctx    context.Context
}

// NewGeminiService creates a new Gemini service
func NewGeminiService(ctx context.Context) (*GeminiService, error) {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GOOGLE_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiService{
		client: client,
		ctx:    ctx,
	}, nil
}

// GenerateImages generates images using Google Gemini API
func (gs *GeminiService) GenerateImages(prompt string, count int) ([]string, error) {
	imageURLs := make([]string, 0, count)

	model := gs.client.GenerativeModel("gemini-2.0-flash")

	// Generate images based on count
	for i := 0; i < count; i++ {
		// Create a unique prompt for each image to get variations
		uniquePrompt := prompt
		if count > 1 {
			uniquePrompt = fmt.Sprintf("%s (variation %d)", prompt, i+1)
		}

		resp, err := model.GenerateContent(gs.ctx, genai.Text(uniquePrompt))
		if err != nil {
			return nil, fmt.Errorf("failed to generate image: %w", err)
		}

		if len(resp.Candidates) == 0 {
			return nil, fmt.Errorf("no candidates in response")
		}

		// Extract image data from the response
		for _, part := range resp.Candidates[0].Content.Parts {
			if blob, ok := part.(genai.Blob); ok {
				imageBytes := []byte(blob.Data)
				// Save the image to a file
				outputFilename := fmt.Sprintf("gemini_generated_image_%d.png", i+1)
				err := os.WriteFile(outputFilename, imageBytes, 0o644)
				if err != nil {
					return nil, fmt.Errorf("failed to write image file: %w", err)
				}
				imageURLs = append(imageURLs, outputFilename)
			}
		}
	}

	if len(imageURLs) == 0 {
		return nil, fmt.Errorf("no images generated")
	}

	return imageURLs, nil
}

// Close closes the Gemini client connection
func (gs *GeminiService) Close() error {
	if gs.client != nil {
		return gs.client.Close()
	}
	return nil
}
