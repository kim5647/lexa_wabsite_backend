package handlers

import (
	"lexa_wabsite_backend/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	// Register принимает DTO и возвращает ошибку.
	Register(input dto.RegisterRequest) error
}

type UserHandler struct {
	// Добавляем ЗАВИСИМОСТЬ
	AuthService IAuthService
}

func NewUserHandler(authService IAuthService) *UserHandler {
	// Принимаем зависимость в конструкторе (Dependency Injection)
	return &UserHandler{
		AuthService: authService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user dto.RegisterRequest

	// 1. Валидация и привязка DTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Вызов СЛОЯ СЕРВИСА для выполнения бизнес-логики
	// Сервис отвечает за: проверку уникальности, хеширование и сохранение в БД.
	if err := h.AuthService.Register(user); err != nil {
		// Логика обработки ошибок (например, 409 Conflict для "уже существует")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		// В реальном коде здесь можно добавить проверку на ошибку типа "уже существует" и вернуть 409
		return
	}

	// 3. Успешный ответ
	c.JSON(http.StatusCreated, gin.H{ // Использовать 201 Created для создания ресурса
		"message": "Пользователь успешно зарегистрирован",
		"email":   user.Email, // Используем email из DTO, чтобы показать, что регистрация прошла
	})
}
