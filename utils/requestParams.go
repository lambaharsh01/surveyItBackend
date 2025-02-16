package utils

import (
	"github.com/harshLamba2/feedbackF/models/structEntities"

	"github.com/gin-gonic/gin"
)

func GetRequestString(c *gin.Context, key string) string {
	if value, exists := c.Get(key); exists {
		if stringValue, ok := value.(string); ok {
			return stringValue
		}
	}
	return ""
}

func GetRequestUint(c *gin.Context, key string) uint {
	if value, exists := c.Get(key); exists {
		if uintValue, ok := value.(uint); ok {
			return uintValue
		}
	}
	return 0
}

func GetRequestInt(c *gin.Context, key string) int {
	if value, exists := c.Get(key); exists {
		if intValue, ok := value.(int); ok {
			return intValue
		}
	}
	return 0
}

func GetRequestParameters(c *gin.Context) *structEntities.AuthToken {

	var userId uint = GetRequestUint(c, "userId")
	var userEmail string = GetRequestString(c, "userEmail")
	var userName string = GetRequestString(c, "userName")
	var userGender string = GetRequestString(c, "userGender")
	var userType string = GetRequestString(c, "userType")
	var ticketGenerationStatus int = GetRequestInt(c, "ticketGenerationStatus")
	
	return &structEntities.AuthToken{
		UserId: userId,
		UserEmail: userEmail,
		UserName: userName,
		UserGender: userGender,
		UserType: userType,
		TicketGenerationStatus: ticketGenerationStatus,
	}
}
