package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(CORSMiddleware())

	// API routes
	api := router.Group("/api/v1")
	{
		// Profile routes
		api.GET("/profile", GetProfile)
		
		// Experience routes
		api.GET("/experiences", GetExperiences)
		api.GET("/experiences/:id", GetExperience)
		
		// Education routes
		api.GET("/education", GetEducation)
		api.GET("/education/:id", GetEducationDetail)
		
		// Skills routes
		api.GET("/skills", GetSkills)
		
		// Projects routes
		api.GET("/projects", GetProjects)
		api.GET("/projects/:id", GetProject)
		
		// Contact routes
		api.POST("/contact", SendContactMessage)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// CORSMiddleware handles CORS
func CORSMiddleware() gin.HandlerFunc {
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
