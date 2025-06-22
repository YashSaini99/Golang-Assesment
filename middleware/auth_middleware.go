package middleware

import (
	"net/http"
	"strings"

	"github.com/YashSaini99/Golang-Assesment/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := service.ParseToken(token)
		if err != nil || (role != "" && user.Role != role) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
