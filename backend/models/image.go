package models

import "time"

// GenerationRequest represents an image generation request
type GenerationRequest struct {
	Prompt      string  `json:"prompt" binding:"required"`
	Model       string  `json:"model" binding:"required"`
	Count       int     `json:"count" binding:"min=1,max=10"`
	Seed        int64   `json:"seed"`
	Steps       int     `json:"steps" binding:"min=1,max=150"`
	CFGScale    float32 `json:"cfg_scale" binding:"min=1,max=20"`
	NegPrompt   string  `json:"negative_prompt"`
	ImageHeight int     `json:"height" binding:"oneof=512 768 1024"`
	ImageWidth  int     `json:"width" binding:"oneof=512 768 1024"`
}

// GenerationResponse represents the response from image generation
type GenerationResponse struct {
	TaskID    string    `json:"task_id"`
	Status    string    `json:"status"`
	Images    []string  `json:"images,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// ImageResult represents a generated image
type ImageResult struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	URL       string    `json:"url"`
	Seed      int64     `json:"seed"`
	CreatedAt time.Time `json:"created_at"`
}

// GenerationTask represents the state of an image generation task
type GenerationTask struct {
	ID        string
	Prompt    string
	Model     string
	Status    string // pending, processing, completed, failed
	Results   []ImageResult
	CreatedAt time.Time
	UpdatedAt time.Time
}
