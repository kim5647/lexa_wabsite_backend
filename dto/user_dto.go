package dto

// RegisterRequest - Модель для входящего запроса POST /users
type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest - Модель для запроса PUT /users/:id
type UpdateUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
