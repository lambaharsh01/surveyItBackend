package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/harshLamba2/feedbackF/models/databaseSchema"
	"github.com/harshLamba2/feedbackF/models/structEntities"
	"github.com/harshLamba2/feedbackF/utils"
	"github.com/harshLamba2/feedbackF/utils/constants"
)

func InitForgotPassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ForgotOtpPayload structEntities.ForgotOtpPayloadStruct

		if err := c.ShouldBindJSON(&ForgotOtpPayload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
		}

		var OtpDetails structEntities.ConfirmOtpResponseStruct

		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", ForgotOtpPayload.UserEmail).First(&OtpDetails).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.AbortWithStatusJSON(c, http.StatusNotFound, "User Not Found")
				return
			}
			c.Error(err)
			return
		}

		var minuteDiff int = utils.TimeDiffMinuet(OtpDetails.InitialOtpSentAt, time.Now())

		if minuteDiff < constants.AttemptsBlockTime && OtpDetails.OtpCount >= constants.AttemptsAllowed {
			var inactivityMessage string = fmt.Sprintf("Too many password change attempts, your authentication activity has been blocked for next %s minutes", strconv.Itoa(constants.AttemptsBlockTime-minuteDiff))
			utils.AbortWithStatusJSON(c, http.StatusForbidden, inactivityMessage)
			return

		}

		query := db.Model(&databaseSchema.Users{}).Where("email = ?", ForgotOtpPayload.UserEmail)

		if minuteDiff < constants.AttemptsBlockTime {
			// UPDATING OTP COUNT BY 1
			if err := query.Update("otp_count", OtpDetails.OtpCount+1).Error; err != nil {
				c.Error(err)
				return
			}
		} else {
			// RESETING OTP COUNT TO 0 AND INTIAL OTP SENT TO CURRENT_TIME
			otpLimitUpdation := map[string]interface{}{
				"otp_count":           0,
				"initial_otp_sent_at": time.Now(),
			}
			if err := query.Updates(otpLimitUpdation).Error; err != nil {
				c.Error(err)
				return
			}
		}

		var randomNumber string = utils.RandomNumber()
		var hashedOtpString string = utils.HashString(randomNumber)

		otpUpdate := &structEntities.ConfirmOtpResponseStruct{
			OTP:       hashedOtpString,
			OtpSentAt: time.Now(),
		}

		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", ForgotOtpPayload.UserEmail).Updates(&otpUpdate).Error; err != nil {
			c.Error(err)
			return
		}

		emailParameter := &structEntities.MailerModel{
			ReciverEmailId: ForgotOtpPayload.UserEmail,
			Subject:        "OTP For Charter Password Reset",
			Body:           fmt.Sprintf(constants.OtpHtmlDesign, randomNumber),
			BodyType:       "html",
		}

		if err := utils.SendEmail(emailParameter); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusInternalServerError, "Email could not be sent at")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": fmt.Sprintf("Otp sent to %s", ForgotOtpPayload.UserEmail),
		})
	}
}

func InitSignUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Step 1: request data validation
		var initSignUpPayloadStruct structEntities.InitSignUpPayloadStruct

		err := c.ShouldBindJSON(&initSignUpPayloadStruct)
		if err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "Invalid request Parameters")
			return
		}

		if validEmail := utils.RegexEmail(initSignUpPayloadStruct.UserEmail); !validEmail {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
			return
		}
		if validPhone := utils.RegexPhone(initSignUpPayloadStruct.PhoneNumber); !validPhone {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
			return
		}

		if validDate := utils.RegexDate(initSignUpPayloadStruct.DateOfBirth); !validDate {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
			return
		}

		DateOfBirth, dateConversionError := time.Parse("2006-01-02", initSignUpPayloadStruct.DateOfBirth)
		if dateConversionError != nil {
			c.Error(dateConversionError)
			return
		}

		// Step 2: Check if user with the same emal exists
		var previousEntery *databaseSchema.Users

		// FETCH HERE
		if err := db.First(&previousEntery, "email = ?", initSignUpPayloadStruct.UserEmail).Error; err == nil {
			utils.AbortWithStatusJSON(c, http.StatusConflict, "User exists, please initeate forgot password process")
			return
		} else {
			if !strings.Contains(err.Error(), "record not found") {
				c.Error(err)
				return
			}
		}

		// Step 3: Prepare data for the email
		var emailParameter *structEntities.MailerModel
		var randomNumber string = utils.RandomNumber()

		var hashedOtpString string = utils.HashString(randomNumber)

		// Step 4: Populate the table row struct
		userRow := &databaseSchema.Users{
			Name:                  initSignUpPayloadStruct.UserName,
			Email:                 initSignUpPayloadStruct.UserEmail,
			PhoneNumber:           initSignUpPayloadStruct.PhoneNumber,
			DateOfBirth:           DateOfBirth,
			Gender:                initSignUpPayloadStruct.Gender,
			OtpSentAt:             time.Now(),
			PasswordLastUpdatedAt: time.Now(),
			InitialOtpSentAt:      time.Now(),
			OTP:                   hashedOtpString,
		}

		// step 5: send email
		emailParameter = &structEntities.MailerModel{
			ReciverEmailId: initSignUpPayloadStruct.UserEmail,
			Subject:        "OTP For Charter Sign Up",
			Body:           fmt.Sprintf(constants.OtpHtmlDesign, randomNumber),
			BodyType:       "html",
			// CC: []string { "email@gmail.com", "email2@gmail.com" },
		}

		if err := utils.SendEmail(emailParameter); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusInternalServerError, "Email could not be sent at")
			return
		}

		if rowInsertion := db.Create(&userRow); rowInsertion.Error != nil {
			c.Error(rowInsertion.Error)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": fmt.Sprintf("Otp sent to %s", initSignUpPayloadStruct.UserEmail),
		})
	}
}

func CheckOTP(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var otpPlayload structEntities.OtpPayloadStruct
		if err := c.ShouldBindJSON(&otpPlayload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
			return
		}

		var confirmOtpResponse *structEntities.ConfirmOtpResponseStruct

		// FETCH HERE
		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", otpPlayload.UserEmail).First(&confirmOtpResponse).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.AbortWithStatusJSON(c, http.StatusNotFound, "User not found")
				return
			}
			c.Error(err)
			return
		}

		var minuteDiff int = utils.TimeDiffMinuet(confirmOtpResponse.InitialOtpSentAt, time.Now())

		if confirmOtpResponse.OtpCount >= constants.AttemptsAllowed {
			var inactivityMessage string = fmt.Sprintf("Too many password change attempts, your authentication activity has been blocked for next %s minutes", strconv.Itoa(constants.AttemptsBlockTime-minuteDiff))
			utils.AbortWithStatusJSON(c, http.StatusForbidden, inactivityMessage)
			return

		}

		var timeDifference int = utils.TimeDiffMinuet(confirmOtpResponse.OtpSentAt, time.Now())
		fmt.Println(timeDifference)

		if timeDifference > constants.OtpValidityMinuets {
			utils.AbortWithStatusJSON(c, http.StatusRequestTimeout, "Timeout")
			return
		}

		if passwordMatched := utils.CompareHashes(otpPlayload.OTP, confirmOtpResponse.OTP); !passwordMatched {

			if err := db.Model(&databaseSchema.Users{}).Where("email = ?", otpPlayload.UserEmail).Update("otp_count", confirmOtpResponse.OtpCount+1).Error; err != nil {
				c.Error(err)
				return
			}

			utils.AbortWithStatusJSON(c, http.StatusUnauthorized, "Incorrect OTP")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Email authenticated successfully",
		})
	}
}

func SetPassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var changePasswordPayload *structEntities.ChangePasswordPayloadStruct
		if err := c.ShouldBindJSON(&changePasswordPayload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, constants.InvalidRequestParameters)
			return
		}

		if weakPass := utils.RegexWeakPassword(changePasswordPayload.Password); weakPass {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "Validation Failed 1")
			return
		}

		var confirmOtpResponse *structEntities.ConfirmOtpResponseStruct

		// FETCH HERE
		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", changePasswordPayload.UserEmail).First(&confirmOtpResponse).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.AbortWithStatusJSON(c, http.StatusUnauthorized, "Validation Failed 2")
				return
			}
			c.Error(err)
			return
		}

		var minuteDiff int = utils.TimeDiffMinuet(confirmOtpResponse.InitialOtpSentAt, time.Now())

		if minuteDiff < constants.AttemptsBlockTime && confirmOtpResponse.OtpCount >= constants.AttemptsAllowed {
			var inactivityMessage string = fmt.Sprintf("Too many password change attempts, your authentication activity has been blocked for next %s minutes", strconv.Itoa(constants.AttemptsBlockTime-minuteDiff))
			utils.AbortWithStatusJSON(c, http.StatusForbidden, inactivityMessage)
			return

		}

		var timeDifference int = utils.TimeDiffMinuet(confirmOtpResponse.OtpSentAt, time.Now())
		fmt.Println(timeDifference)

		if timeDifference > constants.PasswordChangeValidTillMinuets {
			utils.AbortWithStatusJSON(c, http.StatusRequestTimeout, "Timeout")
			return
		}

		if passwordMatched := utils.CompareHashes(changePasswordPayload.OTP, confirmOtpResponse.OTP); !passwordMatched {

			if err := db.Model(&databaseSchema.Users{}).Where("email = ?", changePasswordPayload.UserEmail).Update("otp_count", confirmOtpResponse.OtpCount+1).Error; err != nil {
				c.Error(err)
				return
			}

			utils.AbortWithStatusJSON(c, http.StatusUnauthorized, "Validation Failed 3")
			return
		}

		var hashedPassword string = utils.HashString(changePasswordPayload.Password)

		updatePasswordFields := map[string]interface{}{
			"password": hashedPassword,
			"otp":      nil,
		}

		if err := db.Model(&databaseSchema.Users{}).Where("email = ?", changePasswordPayload.UserEmail).Updates(&updatePasswordFields).Error; err != nil {
			c.Error(errors.New("password could not be updated"))
			return
		}

		authStruct := &structEntities.AuthToken{
			UserEmail: changePasswordPayload.UserEmail,
		}

		token, err := utils.GenerateJWT(authStruct)
		if err != nil {
			c.Error(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Your password has been reset",
			"token":   token,
		})
	}
}
