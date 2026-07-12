package router

import (
	"password-generator/internal/infrastructure/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handlers.PasswordHandler) *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.GET("/health", handler.HealthCheck)

	api := r.Group("/api/v1")
	{
		password := api.Group("/password")
		{
			password.POST("/generate", handler.GeneratePassword)
			password.POST("/validate", handler.ValidatePassword)
		}
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}