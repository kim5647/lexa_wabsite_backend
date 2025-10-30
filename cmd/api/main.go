package main

import (
	"log"

	"lexa_wabsite_backend/config" // <--- 1. Только для вызова config.Config()
	"lexa_wabsite_backend/handlers"
	"lexa_wabsite_backend/router"
	"lexa_wabsite_backend/service"

	database "lexa_wabsite_backend/db"                  // <--- 2. Пакет с ConnectPool()
	sqlcgen "lexa_wabsite_backend/repository"           // <--- 3. Сгенерированный SQLC-код (пакет 'repository')
	userRepoImpl "lexa_wabsite_backend/user_repository" // <--- 4. Ваш пользовательский репозиторий (содержит NewUserRepository)

	"github.com/gin-gonic/gin"
)

func main() {
	config.Config() // Если config.Config() что-то делает, вызываем его.

	// 1. Подключение к БД
	conn, err := database.ConnectPool() // <--- Используем 'database'
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer conn.Close()

	// --- 2. Инициализация слоев ---

	// 2.1. Инициализация Репозитория.
	// New(conn) находится в сгенерированном sqlcgen.
	sqlQueries := sqlcgen.New(conn)

	// NewUserRepository находится в вашем пользовательском пакете user_repository
	userRepo := userRepoImpl.NewUserRepository(sqlQueries)

	// 2.2. Инициализация Сервисов
	authService := service.NewAuthService(userRepo)

	// 2.3. Инициализация Обработчиков (Последняя ошибка!)
	userHandler := handlers.NewUserHandler(authService) // <-- Требует реализации Register

	// --- 3. Запуск ---
	deps := &router.Dependencies{UserHandler: userHandler}
	r := gin.Default()
	router.RegisterRoutes(r, deps)

	if err := r.Run(":8080"); err != nil {
		log.Panicf("Ошибка при запуске сервера: %v", err)
	}
}
