package main

import (
	"context"
	"log"
	"net/http"

	"github.com/BowieHe/img-gacha/models"
	"github.com/BowieHe/img-gacha/services"
	"github.com/gin-gonic/gin"
)

var imageGenerator *services.ImageGenerator

func main() {
	router := gin.Default()

	// Initialize context
	ctx := context.Background()

	// Initialize image generator
	var err error
	imageGenerator, err = services.NewImageGenerator(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize image generator: %v", err)
	}

	// Enable CORS
	router.Use(corsMiddleware())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		api.POST("/generate", generateImage)
		api.GET("/models", getAvailableModels)
		api.GET("/status/:taskId", getGenerationStatus)
	}

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func generateImage(c *gin.Context) {
	var req models.GenerationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Set default values
	if req.Steps == 0 {
		req.Steps = 50
	}
	if req.CFGScale == 0 {
		req.CFGScale = 7.5
	}
	if req.Count == 0 {
		req.Count = 1
	}
	if req.ImageHeight == 0 {
		req.ImageHeight = 768
	}
	if req.ImageWidth == 0 {
		req.ImageWidth = 768
	}

	resp, err := imageGenerator.GenerateImages(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func getAvailableModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"models": []gin.H{
			{
				"id":          "gemini-2.0-flash",
				"name":        "Gemini 2.0 Flash",
				"description": "Fast and powerful image generation with Gemini",
				"provider":    "Google Gemini",
			},
			{
				"id":          "gpt-4-image",
				"name":        "GPT-4 Enhanced Generation",
				"description": "Uses GPT-4 to enhance prompts, then generates with Gemini",
				"provider":    "OpenAI + Google Gemini",
			},
		},
	})
}

func getGenerationStatus(c *gin.Context) {
	taskId := c.Param("taskId")
	c.JSON(http.StatusOK, gin.H{
		"taskId": taskId,
		"status": "pending",
	})
}
