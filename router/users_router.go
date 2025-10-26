package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes настраивает все маршруты для ресурса "user"
func RegisterUserRoutes(r *gin.Engine) {
	// Создаем группу с префиксом /users
	userGroup := r.Group("/users")
	{
		// Привязываем маршруты к функциям-обработчикам из пакета handlers
		userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", handlers.CreateUser)
		userGroup.GET("/:id", handlers.GetUserByID)
		// ... и т.д.
	}
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
