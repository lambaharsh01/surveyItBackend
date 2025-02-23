package utils

import (
	"errors"
	"fmt"

	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
	"gopkg.in/gomail.v2"
)

func SendEmail(mailerOption *structEntities.MailerModel) error {

	m := gomail.NewMessage()

	var senderEmailId string = GetEnv("EMAIL_ID")
	var senderEmailPassword string = GetEnv("EMAIL_PASSWORD")

	m.SetHeader("From", senderEmailId)
	m.SetHeader("To", mailerOption.ReceiverEmailId)

	for _, cc := range mailerOption.CC {
		m.SetHeader("Cc", cc)
	}

	for _, bcc := range mailerOption.BCC {
		m.SetHeader("Bcc", bcc)
	}

	m.SetHeader("Subject", mailerOption.Subject)

	if mailerOption.BodyType != "plain" && mailerOption.BodyType != "html" {
		return errors.New("BodyType is not valid")
	}

	var emailBodyType string = fmt.Sprintf("text/%s", mailerOption.BodyType)

	m.SetBody(emailBodyType, mailerOption.Body)

	d := gomail.NewDialer("smtp.gmail.com", 587, senderEmailId, senderEmailPassword)

	return d.DialAndSend(m)
}
