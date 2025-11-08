package handlers

import (
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
	// Внутри этого метода вы будете вызывать: h.AuthService.RegisterNewUser(...)
	c.JSON(http.StatusOK, "Ты лучший")
}

// ДОБАВИТЬ: Метод GetUsers
func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUsers is working"})
}

// ДОБАВИТЬ: Метод GetUserByID
func (h *UserHandler) GetUserByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetUserByID is working"})
}
