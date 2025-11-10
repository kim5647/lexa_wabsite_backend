package service

import (
	"errors"
	"time"

	// 💡 Вам нужно будет создать структуру Config,
	// или временно использовать глобальные переменные, как здесь.

	"github.com/golang-jwt/jwt/v5"
)

// UserClaims - пользовательские данные, которые будут храниться в токене.
// Поле UserID должно соответствовать типу ID в вашей sqlc.User (обычно int64).
type UserClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// JWTConfig - Структура для хранения настроек JWT.
// В реальном приложении это поле должно быть частью вашей основной Config.
type JWTConfig struct {
	SecretKey       string
	ExpirationHours int
}

// NewJWTConfig - создает конфигурацию. В реальном приложении данные берутся из YAML.
func NewJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKey:       "super_secret_signing_key_change_me", // !!! ВАШ СЕКРЕТНЫЙ КЛЮЧ !!!
		ExpirationHours: 24,
	}
}

// GenerateToken - генерирует JWT для пользователя.
// Принимает ID пользователя и Email, а также конфигурацию.
func GenerateToken(userID int64, email string, cfg JWTConfig) (string, error) {

	secretKey := []byte(cfg.SecretKey)
	expirationTime := time.Now().Add(time.Hour * time.Duration(cfg.ExpirationHours))

	claims := &UserClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func ValidateToken(tokenString string, cfg JWTConfig) (*UserClaims, error) {
	secretKey := []byte(cfg.SecretKey)

	// Парсинг токена
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверка алгоритма подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверка валидности (срок годности и т.д.)
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
