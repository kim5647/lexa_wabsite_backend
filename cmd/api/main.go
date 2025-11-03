package main

import (
	"lexa_wabsite_backend/handlers"
	repository "lexa_wabsite_backend/repository"
	sqlc "lexa_wabsite_backend/repository/sqlc"
	"lexa_wabsite_backend/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Подключение к БД
	conn, err := repository.ConnectPool() // <--- Используем 'database'
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer conn.Close()

	// --- 2. Инициализация слоев ---

	// 2.1. Инициализация Репозитория.
	// New(conn) находится в сгенерированном sqlcgen.
	sqlQueries := sqlc.New(conn)

	// 2.2. Инициализация Сервисов
	// authService := service.NewAuthService(userRepo)
	userRepo := sqlc.NewUserRepository(sqlQueries)
	// 2.3. Инициализация Обработчиков (Последняя ошибка!)
	userHandler := handlers.NewUserHandler(userRepo) // <-- Требует реализации Register

	// --- 3. Запуск ---
	deps := &router.Dependencies{UserHandler: userHandler}
	r := gin.Default()
	router.RegisterRoutes(r, deps)

	if err := r.Run(":8080"); err != nil {
		log.Panicf("Ошибка при запуске сервера: %v", err)
	}
}
