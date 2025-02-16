package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/controllers/fetch"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
	"gorm.io/gorm"
)

func FetchRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.GET(constants.GetQuestionTypes, fetch.GetQuestionTypes(db))
	router.GET(constants.GetFileTypes, fetch.GetFileTypes(db))

}
