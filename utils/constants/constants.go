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
	SurveySubmission          string = "/survey-submission"

	GetResponseData string = "/get-response-data/:surveyCode"
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

const ThankYouEmailTemplate = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Thank You!</title>
</head>
<body style="font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #f8f9fa; margin: 0; padding: 40px 20px; color: #333;">

<div style="max-width: 600px; background: #ffffff; padding: 30px; border-radius: 4px; border: 1px solid #e1e4e8; box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05); margin: auto;">
    <div style="text-align: center; padding-bottom: 20px; border-bottom: 1px solid #eaeaea; margin-bottom: 25px;">
        <h2 style="color: #2c3e50; font-weight: 600; margin-top: 0;">Thank You <span style="color: #3498db;">for Your Response</span></h2>
    </div>
    
    <div style="text-align: left; line-height: 1.6;">
        <p style="color: #4a4a4a; font-size: 16px; margin-bottom: 16px;">I sincerely appreciate you taking the time to complete the response form.</p>
        <p style="color: #4a4a4a; font-size: 16px; margin-bottom: 16px;">Your valuable input will help me improve myself.</p>
        <p style="color: #4a4a4a; font-size: 16px; margin-bottom: 16px;">If you have any additional questions or feedback, please don't hesitate to reach out.</p>
        
        <div style="margin-top: 30px; padding-top: 20px; border-top: 1px solid #eaeaea;">
            <p style="color: #4a4a4a; font-size: 16px; margin-bottom: 16px;">Best regards,</p>
            <p style="color: #4a4a4a; font-size: 16px; margin-bottom: 16px;"><strong>%s</strong></p>
        </div>
    </div>
</div>

</body>
</html>
`
