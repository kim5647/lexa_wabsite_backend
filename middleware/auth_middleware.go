package middleware

import (
	"lexa_wabsite_backend/service" // Используем ваш пакет service
	"net/http"

	"github.com/gin-gonic/gin"
)

const userContextKey = "userClaims" // Ключ для хранения данных пользователя

func AuthMiddleware(authService service.IAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Извлечение токена из куки (предпочтительный способ)
		tokenString, err := c.Cookie("auth_token")

		if err != nil {
			// Токен отсутствует
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			c.Abort()
			return
		}

		// 2. Валидация через внедренный сервис (чистая DI)
		// 💡 Используем метод authService, который знает о секрете.
		claims, err := authService.ValidateToken(tokenString)
		if err != nil {
			// Токен недействителен или истёк
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 3. Успех: Сохраняем данные для использования в обработчике
		c.Set(userContextKey, claims)

		// 4. Переход к следующему обработчику
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
