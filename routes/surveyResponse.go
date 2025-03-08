package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/controllers"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
	"gorm.io/gorm"
)

func SurveyResponseRoutes(router *gin.RouterGroup, db *gorm.DB) {
	
	router.GET(constants.FetchSurveyAndQuestionary, controllers.FetchSurveyAndQuestionary(db))

}
