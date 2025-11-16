package service

import "github.com/golang-jwt/jwt/v5"

// UserClaims - пользовательские данные в токене.
type UserClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// JWTConfig - Структура для хранения настроек JWT.
type JWTConfig struct {
	SecretKey       string
	ExpirationHours int
}

// NewJWTConfig - создает конфигурацию.
func NewJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKey:       "super_secret_signing_key_change_me",
		ExpirationHours: 24,
	}
}
