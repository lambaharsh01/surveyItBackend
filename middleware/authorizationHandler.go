package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/utils"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := utils.GetRequestParameters(c)

		if user.UserType != "admin" {
			utils.AbortWithStatusJSON(c, http.StatusForbidden, "You are not authorized to access this resource")
			return
		}

		c.Next()
	}
}

// authorizationHandler
