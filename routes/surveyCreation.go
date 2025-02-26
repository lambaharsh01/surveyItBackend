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

	router.POST(constants.AddSurvey, controllers.AddSurvey(db))
	router.PUT(constants.UpdateSurvey, controllers.UpdateSurvey(db))
	router.GET(constants.GetSurveyAndQuestionary, controllers.GetSurveyAndQuestionary(db))
	router.POST(constants.GetSurveys, controllers.GetSurveys(db))
	router.DELETE(constants.DeleteSurvey, controllers.DeleteSurvey(db))

	router.POST(constants.UpdateQuestionary, controllers.UpdateQuestionary(db))

}
