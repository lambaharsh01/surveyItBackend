package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/harshLamba2/feedbackF/models/databaseSchema"
	"github.com/harshLamba2/feedbackF/models/structEntities"
	"github.com/harshLamba2/feedbackF/utils"
)



func Me(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		user := utils.GetRequestParameters(c)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data": user,
		})

	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var LoginPayload *structEntities.LoginPayloadStruct
		var LoginUserDataResponse *structEntities.LoginUserDataResponseStruct

		if err := c.BindJSON(&LoginPayload); err != nil {
			c.Error(err)
			return
		}

		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", LoginPayload.UserEmail).First(&LoginUserDataResponse); err.Error != nil {
			if errors.Is(err.Error, gorm.ErrRecordNotFound) {
				utils.AbortWithStatusJSON(c, http.StatusNotFound, "User Not Found")
				return
			}
			c.Error(err.Error)
			return
		}

		if passwordMatched := utils.CompareHashes(LoginPayload.Password, LoginUserDataResponse.Password); !passwordMatched {
			utils.AbortWithStatusJSON(c, http.StatusUnauthorized, "Incorrect Credentials")
			return
		}

		fmt.Println(LoginUserDataResponse, "LoginUserDataResponse")

		tokenInfo := &structEntities.AuthToken{
			UserId: LoginUserDataResponse.ID,
			UserEmail: LoginUserDataResponse.Email,
			UserName: LoginUserDataResponse.Name,
			UserGender: LoginUserDataResponse.Gender,
			UserType: LoginUserDataResponse.UserType,
			TicketGenerationStatus: LoginUserDataResponse.TicketGenerationStatus,
		}

		token, tokenGenerationErr := utils.GenerateJWT(tokenInfo)

		if tokenGenerationErr != nil {
			c.Error(tokenGenerationErr)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Login Successful",
			"token":   token,
		})

	}
}



