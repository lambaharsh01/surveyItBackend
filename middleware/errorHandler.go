package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.IsAborted() {
			return
		}

		if len(c.Errors) > 0 {

			log.Println("Error :" + c.Errors.String())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   c.Errors.String(),
			})
		}
	}
}
