package middleware

import (
	"lexa_wabsite_backend/service" // Используем ваш пакет service
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationHeader = "Authorization"
const userContextKey = "userClaims" // Ключ для хранения данных пользователя в контексте Gin

// AuthMiddleware - фабрика для создания middleware с внедренным сервисом
func AuthMiddleware(authService service.IAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Извлечение заголовка Authorization
		authHeader := c.GetHeader(authorizationHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// 2. Проверка формата "Bearer <token>"
		fields := strings.Fields(authHeader)
		if len(fields) < 2 || strings.ToLower(fields[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Must be Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := fields[1]

		// 3. Валидация токена через Service Layer
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 4. Успех: Сохранение данных пользователя в контексте Gin
		c.Set(userContextKey, claims)

		// 5. Продолжение обработки запроса
		c.Next()
	}
}

// GetUserClaimsFromContext - Вспомогательная функция для Handler
func GetUserClaimsFromContext(c *gin.Context) *service.UserClaims {
	if claims, exists := c.Get(userContextKey); exists {
		if userClaims, ok := claims.(*service.UserClaims); ok {
			return userClaims
		}
	}
	return nil
}
