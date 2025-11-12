package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
	c.JSON(http.StatusOK, gin.H{
		"message": "Image generation endpoint",
	})
}

func getAvailableModels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"models": []string{
			"stable-diffusion-v1.5",
			"stable-diffusion-v2.1",
			"dall-e-3",
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
