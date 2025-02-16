package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func ErrorHandler() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Next()

		if c.IsAborted() {
			return // If the context is already aborted, stop further processing
		}
		
// import "errors"
// err:= errors.New("Hello World");
// c.Error(errors.New("Hello World")); // adds/ collects error in the Context which can be processed later
// return;
		
		if( len(c.Errors) > 0 ){	

			log.Println("Error :" + c.Errors.String())
			
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"success":false,
				"error":c.Errors.String(),
			})
		}
	}
}