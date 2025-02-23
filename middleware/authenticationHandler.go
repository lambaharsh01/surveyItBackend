package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/utils"
)

func AuthenticationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Authorization header missing",
			})
			return
		}

		// Extract token from "Bearer <token>" format
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   "Invalid or expired token",
			})
			return
		}

		c.Set("userId", claims.UserId)
		c.Set("userEmail", claims.UserEmail)
		c.Set("userName", claims.UserName)
		c.Set("userGender", claims.UserGender)
		c.Set("userType", claims.UserType)
		c.Next()
	}
}
