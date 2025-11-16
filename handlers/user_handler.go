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
		"user_id": createdUser.ID,
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
