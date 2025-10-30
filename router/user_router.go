package router

import (
	"lexa_wabsite_backend/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes настраивает все маршруты для ресурса "user"
func AddUserRoutes(rg *gin.RouterGroup, h *handlers.UserHandler) {
	// Создаем группу с префиксом /users
	userGroup := rg.Group("/users")
	{
		// Привязываем маршруты к функциям-обработчикам из пакета handlers
		// userGroup.GET("/", handlers.GetUsers)
		userGroup.POST("/", h.CreateUser)
		// userGroup.GET("/:id", handlers.GetUserByID)
		// ... и т.д.
	}
}
