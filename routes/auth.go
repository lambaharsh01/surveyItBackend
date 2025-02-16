package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lambaharsh01/surveyItBackend/controllers/auth"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
	"gorm.io/gorm"
)

func MeRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.GET(constants.Me, auth.Me(db))

}
func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST(constants.SignIn, auth.Login(db))

	router.POST(constants.InitForgotPassword, auth.InitForgotPassword(db))
	router.POST(constants.InitSignUp, auth.InitSignUp(db))
	router.POST(constants.CheckOTP, auth.CheckOTP(db))
	router.POST(constants.SetPassword, auth.SetPassword(db))

}
