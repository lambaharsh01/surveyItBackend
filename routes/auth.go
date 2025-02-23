package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/controllers"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
	"gorm.io/gorm"
)

func MeRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET(constants.Me, controllers.Me(db))
}
func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST(constants.SignIn, controllers.Login(db))
	router.POST(constants.InitForgotPassword, controllers.InitForgotPassword(db))
	router.POST(constants.InitSignUp, controllers.InitSignUp(db))
	router.POST(constants.CheckOTP, controllers.CheckOTP(db))
	router.POST(constants.SetPassword, controllers.SetPassword(db))
}
