package middleware

import (
	"net/http"
	"strings"
	"vibex-api/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header not found"})
			c.Abort()
			return
		}

		token := strings.Split(authHeader, "Bearer ")
		if len(token) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			c.Abort()
			return
		}

		// Validate token
		userId, err := jwtService.ValidateToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
