package utils

import (
	"github.com/gin-gonic/gin"
)

func AbortWithStatusJSON(c *gin.Context, errorCode int, message string) {

	c.AbortWithStatusJSON(errorCode, gin.H{
		"success" : false, 
		"error" : message,
	})
	c.Abort()
}