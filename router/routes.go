package router

import (
	"lexa_wabsite_backend/handlers"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	UserHandler *handlers.UserHandler
}

func RegisterRoutes(r *gin.Engine, deps *Dependencies) {
	// Создаем группу для пользователя
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", deps.UserHandler.GetUsers) // Убедитесь, что GetUsers - метод *UserHandler
		userGroup.POST("/", deps.UserHandler.CreateUser)
		userGroup.GET("/:id", deps.UserHandler.GetUserByID)
	}
}
