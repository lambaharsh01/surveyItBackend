package fetch

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/harshLamba2/feedbackF/models/databaseSchema"
)

func GetQuestionTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		questionType := databaseSchema.QuestionType{}

		if err:= db.Find(&questionType).Error; err!=nil{
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"questionType":questionType,
		})

	}
}

func GetFileTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileType := databaseSchema.FileType{}

		if err:= db.Find(&fileType).Error; err!=nil{
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":true,
			"fileType":fileType,
		})

	}
}


