package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harshLamba2/feedbackF/controllers/fetch"
	"github.com/harshLamba2/feedbackF/utils/constants"
	"gorm.io/gorm"
)


func FetchRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.GET(constants.GetQuestionTypes, fetch.GetQuestionTypes(db))
	router.GET(constants.GetFileTypes, fetch.GetFileTypes(db))

}
