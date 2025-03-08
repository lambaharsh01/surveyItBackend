package constants

const (
	RootUser string = "lambaharsh01@gmail.com"
)
const AlphanumericCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const (
	OtpValidityMinuets             int = 3
	PasswordChangeValidTillMinuets int = 10
)

const (
	InvalidRequestParameters string = "Invalid Request Parameters"
	AttemptsBlockTime        int    = 30
	AttemptsAllowed          int    = 5
)

const (
	SignIn string = "/sign-in"
	Me     string = "/me"

	InitForgotPassword string = "/forgot-password"

	InitSignUp       string = "/initial-sign-up"
	CheckOTP         string = "/check-otp"
	SetPassword      string = "/set-password"
	GetQuestionTypes string = "/get-question-types"
	GetFileTypes     string = "/get-file-types"

	AddSurvey               string = "/add-survey"
	UpdateSurvey            string = "/update-survey/:id"
	GetSurveys              string = "/get-surveys"
	GetSurveyAndQuestionary string = "/get-survey-and-questionary/:surveyCode"
	DeleteSurvey            string = "/delete-survey/:id"
	UpdateQuestionary       string = "/update-questionary"

	FetchSurveyAndQuestionary string = "/fetch-survey-and-questionary/:surveyCode"
)

const OtpHtmlDesign string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>OTP Verification</title>
</head>
<body style="margin: 0; padding: 0; font-family: Arial, sans-serif; text-align: center; background-color: #f5f5f5; height: 86vh;">

  <div style="padding: 20px; font-size: 10px; font-weight: 500; background-color: #355ba6; color: white;">
    Your One-Time Password (OTP) for Charter Authentication is:
	<br/>
	<span style="font-size: 8px; font-weight: 400;">(If you did not request this process, please disregard this message)</span>
  </div>
  <br/>
  <br/>
  <br/>
  <br/>
  <br/>
  <br/>
  <br/>
      <span style="font-size: 40px; font-weight: 600; color: #355ba6; margin-top:80px">
	  %s
	  </span>
	  <br/>
	  <br/>
</body>
</html>
`
