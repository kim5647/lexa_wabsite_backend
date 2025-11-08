package main

import (
	database "lexa_wabsite_backend/db"
	repository "lexa_wabsite_backend/db/repository"
	db "lexa_wabsite_backend/db/sqlc"
	"lexa_wabsite_backend/handlers"
	"lexa_wabsite_backend/router"
	"lexa_wabsite_backend/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Подключение к БД
	conn, err := database.ConnectPool() // <--- Используем 'database'
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer conn.Close()

	// 2.1. Инициализация Репозитория.
	sqlQueries := db.New(conn)
	// userRepo находится в пакете 'db' (по вашему импорту)
	userRepo := repository.NewUserRepository(sqlQueries)

	// 2.2. Инициализация Сервисов
	// 💡 Создаем сервис, передавая в него репозиторий
	authService := service.NewAuthService(userRepo)

	// 2.3. Инициализация Обработчиков
	// 💡 ПЕРЕДАЕМ СЕРВИС (authService) в Handler
	userHandler := handlers.NewUserHandler(authService)

	// --- 3. Запуск ---
	deps := &router.Dependencies{UserHandler: userHandler}
	r := gin.Default()
	router.RegisterRoutes(r, deps)

	if err := r.Run(":8080"); err != nil {
		log.Panicf("Ошибка при запуске сервера: %v", err)
	}
}
