package adminPanel

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/harshLamba2/feedbackF/models/databaseSchema"
	"github.com/harshLamba2/feedbackF/models/structEntities"
)

func GetAllUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.PaginationStruct
		if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
			c.Error(err)
			return
		}

		var limit int = payload.Offset
		var offset int = (payload.Offset * payload.Page)

		// userDetails := []structEntities.UsersDetails{}

		var usersDetails []map[string]interface{}

		if err := db.Raw(`SELECT 
			name AS name, 
			email AS email, 
			phone_number AS phoneNumber, 
			gender AS gender,
			user_type AS userType, 
			ticket_generation_status AS ticketGenerationStatus, 
			DATE_FORMAT(created_at, '%y %b, %d | %h:%i %p') AS createdAt 
		FROM users 
		ORDER BY name
		LIMIT ? 
		OFFSET ?`, limit, offset).Scan(&usersDetails).Error; err != nil {
			c.Error(err)
			return
		}

		var total int64
		db.Model(&databaseSchema.Users{}).Count(&total)

		c.JSON(http.StatusOK, gin.H{
			"success":	true,
			"data":  	usersDetails,
			"total": 	total,
			"from":  	offset,
			"to":    	offset + limit,
		})
	}
}

func PermissionAccess(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.TicketingPermissionAccess

		if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
			c.Error(err)
			return
		}

		var updatedTicketGenerationStatus int = 401
		if payload.TicketGenerationStatus {
			updatedTicketGenerationStatus = 200
		}
		
		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", payload.UserEmail).Update("ticket_generation_status", updatedTicketGenerationStatus).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"ticketGenerationStatus": updatedTicketGenerationStatus,
			"message":"Status Changed",
		})
	}
}
