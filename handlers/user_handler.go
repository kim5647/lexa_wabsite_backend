package handlers

import (
	"lexa_wabsite_backend/dto"
	"lexa_wabsite_backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	AuthService service.IAuthService
}

func NewUserHandler(authService service.IAuthService) *UserHandler {
	return &UserHandler{
		AuthService: authService,
	}
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}
	createdUser, err := h.AuthService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"user_id": createdUser.ID, // Предполагая, что sqlc.User имеет поле ID
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req dto.LoginRequest

	// ... (1. Валидация входных данных)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// 2. Аутентификация и получение токена
	token, err := h.AuthService.Login(c.Request.Context(), req)

	if err != nil {
		// ... (обработка ошибок аутентификации)
		if err.Error() == "неверный email или пароль" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Произошла внутренняя ошибка"})
		return
	}

	// 3. 🍪 Установка HTTP-Only Cookie
	// Время жизни куки должно совпадать с ExpiresAt в JWT (наш пример: 24 часа)
	const maxAgeSeconds = 24 * 60 * 60

	c.SetCookie(
		"auth_token",  // Имя куки
		token,         // Значение (JWT)
		maxAgeSeconds, // MaxAge
		"/",           // Путь: доступно на всех путях
		"",            // Домен: пусто для текущего домена
		true,          // Secure: true - только по HTTPS (Обязательно для продакшена!)
		true,          // HttpOnly: true - JS не имеет доступа (Защита от XSS)
	)

	// 4. Успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Вход выполнен успешно. Токен записан в куки.",
	})
}

// ДОБАВИТЬ: Метод GetUsers
func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUsers is working"})
}

// ДОБАВИТЬ: Метод GetUserByID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByID is working"})
}
