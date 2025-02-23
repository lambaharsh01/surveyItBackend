package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/controllers"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
	"gorm.io/gorm"
)

func FetchRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.GET(constants.GetQuestionTypes, controllers.GetQuestionTypes(db))
	router.GET(constants.GetFileTypes, controllers.GetFileTypes(db))

	router.POST(constants.GetFileTypes, controllers.AddSurvey(db))

}
