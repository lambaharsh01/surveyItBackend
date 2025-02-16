package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
)

func GenerateJWT(tokenInfo *structEntities.AuthToken) (string, error) {

	secretKey := GetEnv("SECRET_KEY")
	var jwtSecret = []byte(secretKey)

	var userIdString string = strconv.FormatUint(uint64(tokenInfo.UserId), 10)
	var ticketGenerationStatusString string = strconv.Itoa(tokenInfo.TicketGenerationStatus)

	claims := jwt.MapClaims{
		"userId":                 userIdString,
		"userEmail":              tokenInfo.UserEmail,
		"userName":               tokenInfo.UserName,
		"userGender":             tokenInfo.UserGender,
		"userType":               tokenInfo.UserType,
		"ticketGenerationStatus": ticketGenerationStatusString,
		"exp":                    time.Now().Add((time.Hour * 24) * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func ValidateJWT(tokenString string) (*structEntities.AuthToken, error) {

	secretKey := GetEnv("SECRET_KEY")
	var jwtSecret = []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // check if the encryption is done with (HS256 || HS384 || HS512) encription standards or not to begain with
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		var userEmail string = claims["userEmail"].(string)
		var userName string = claims["userName"].(string)
		var userGender string = claims["userGender"].(string)
		var userType string = claims["userType"].(string)

		var userId uint
		userIdString := claims["userId"].(string)
		parsedUserId, err := strconv.ParseUint(userIdString, 10, 64)
		if err != nil {
			return nil, err
		}
		userId = uint(parsedUserId)

		ticketGenerationStatusString := claims["ticketGenerationStatus"].(string)
		ticketGenerationStatus, err := strconv.Atoi(ticketGenerationStatusString)
		if err != nil {
			return nil, err
		}

		authToken := &structEntities.AuthToken{
			UserId:                 userId,
			UserEmail:              userEmail,
			UserName:               userName,
			UserGender:             userGender,
			UserType:               userType,
			TicketGenerationStatus: ticketGenerationStatus,
		}

		return authToken, nil
	}

	return nil, errors.New("invalid token")
}
