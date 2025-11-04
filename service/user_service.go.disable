// lexa_wabsite_backend/handlers/user_handler.go
package handlers

// Импорт Service, который вы хотите внедрить
import (
	"lexa_wabsite_backend/service"
	// ...
)

// Обязательно добавьте поле для сервиса
type UserHandler struct {
	AuthService service.IAuthService // Убедитесь, что IAuthService существует
}

// Конструктор должен принимать сервис!
func NewUserHandler(authService service.IAuthService) *UserHandler {
	return &UserHandler{
		AuthService: authService,
	}
}
