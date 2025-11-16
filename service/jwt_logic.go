package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken - генерирует JWT для пользователя, используя внедренную конфигурацию.
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
	return token.SignedString(secretKey)
}

// ValidateToken - парсит и проверяет токен.
func ValidateToken(tokenString string, cfg JWTConfig) (*UserClaims, error) {
	secretKey := []byte(cfg.SecretKey)

	// 1. Парсинг токена с проверкой claims и подписи
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	// 2. Проверка ошибок парсинга (включая истечение срока годности)
	if err != nil {
		return nil, err
	}

	// 3. Извлечение и проверка claims
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
