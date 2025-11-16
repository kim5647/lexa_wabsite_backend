package router

import (
	"lexa_wabsite_backend/handlers"
	"lexa_wabsite_backend/service"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	UserHandler *handlers.UserHandler
	AuthService *service.IAuthService
}

func RegisterRoutes(r *gin.Engine, deps *Dependencies) {
	// Создаем группу для пользователя
	userGroup := r.Group("/users")
	{
		userGroup.POST("/reg", deps.UserHandler.CreateUser)
		userGroup.POST("/login", deps.UserHandler.LoginUser)
	}
}
