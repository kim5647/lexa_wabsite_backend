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
		userGroup.POST("/", deps.UserHandler.CreateUser)
	}
}
