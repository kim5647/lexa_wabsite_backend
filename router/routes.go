package router

import (
	"lexa_wabsite_backend/handlers"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	UserHandler *handlers.UserHandler
	// Другие хендлеры
	// ProductHandler *handlers.ProductHandler
}

// RegisterRoutes - Главная функция, которая собирает все маршруты.
func RegisterRoutes(r *gin.Engine, deps *Dependencies) {

	// Создаем корневую группу, которая использует путь "/"
	// Это гарантирует, что маршруты будут иметь только тот префикс, который вы зададите
	rootGroup := r.Group("/")
	{
		// 1. Регистрируем роуты пользователей внутри корневой группы
		// Маршруты будут: /users, /users/:id
		AddUserRoutes(rootGroup, deps.UserHandler)

		// 2. Добавляем /ping (будет доступен как /ping)
		addPingRoutes(rootGroup)
	}
}

// addPingRoutes (остается без изменений, так как принимает группу)
func addPingRoutes(rg *gin.RouterGroup) {
	rg.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
}
