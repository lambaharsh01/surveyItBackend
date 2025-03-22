package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lambaharsh01/surveyItBackend/models/databaseSchema"
	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
	"github.com/lambaharsh01/surveyItBackend/utils"
	"github.com/lambaharsh01/surveyItBackend/utils/constants"
)

func FetchSurveyAndQuestionary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var surveyCode string= c.Param("surveyCode")

		survey := structEntities.SurveysResponseStruct{}
		questionaryRaw := []structEntities.QuestionRawResponseStruct{}
		questionary := []structEntities.QuestionResponseStruct{}

		if err := db.Model(&databaseSchema.SurveySchema{}).
			Select(`id, 
			survey_code, 
			survey_name, 
			survey_description, 
			survey_target_audience, 
			survey_alignment, 
			survey_color_theme, 
			allow_multiple_submissions,
			DATE_FORMAT(active_from, '%Y-%m-%d') AS active_from, 
			DATE_FORMAT(active_to, '%Y-%m-%d') AS active_to,
			CASE WHEN CURRENT_DATE BETWEEN active_from AND active_to THEN 1 ELSE 0 END AS active,
			DATE_FORMAT(created_at, '%D %M %Y') AS created_at`).
			Where("survey_code = ? AND deleted_at IS NULL", surveyCode).
			Order("id DESC").
			First(&survey).Error; err != nil {
			c.Error(err)
			return
		}

		if !survey.Active {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"survey":    survey,
			})
			return
		}

		var getQuestionsQuery string =`SELECT 
			qs.id, 
			qs.question AS text, 
			qt.id AS question_type_id,
			qt.question_type,
			ft.id AS file_type_id,
			ft.file_type,
			qs.options, 
			qs.required, 
			qs.validation , 
			qs.min, 
			qs.max
		FROM question_schemas qs
		JOIN survey_schemas ss ON ss.id = qs.survey_id
		JOIN question_types qt ON qt.id = qs.question_type_id 
		LEFT JOIN file_types ft ON ft.id = qs.file_type_id
		WHERE ss.survey_code = ? AND qs.deleted_at IS NULL`

		if err := db.Raw(getQuestionsQuery, surveyCode).Scan(&questionaryRaw).Error; err != nil {
			c.Error(err)
			return
		}
		
		for _, raw := range questionaryRaw {

			var unmarshaledSlice []string

			if err:= json.Unmarshal(raw.Options, &unmarshaledSlice); err!=nil {
				c.Error(err)
				return
			}

			question := structEntities.QuestionResponseStruct{
				ID: raw.ID,
				Text: raw.Text,
				QuestionTypeID: raw.QuestionTypeID,
				QuestionType: raw.QuestionType,
				FileTypeID: raw.FileTypeID,
				FileType: raw.FileType,
				Options: unmarshaledSlice,
				Required: raw.Required,
				Validation: raw.Validation,
				Min: raw.Min,
				Max: raw.Max,
			}
			questionary = append(questionary, question)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"survey":    survey,
			"questionary":    questionary,

		})

	}
}



func SurveySubmission(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.SurveySubmissionPayloadStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		
		var getFromDataQuery string = `
		SELECT
			1 AS not_first_response,
			ss.allow_multiple_submissions,
			COUNT(DISTINCT srs.id) AS previous_responses,
			u.name AS user_name
		FROM survey_schemas ss
		LEFT JOIN survey_response_summaries srs ON srs.survey_id = ss.id
		LEFT JOIN users u ON u.id = ss.created_by
		WHERE ss.id = ? AND srs.respondent_email = ?
		GROUP BY ss.id, u.name
		`

		var formData structEntities.SurveyResponseFormDataResponseStruct
		if err := db.Raw(getFromDataQuery, payload.SurveyId, payload.RespondentEmail).Scan(&formData).Error; err != nil {
			c.Error(err)
			return
		}

		
		if formData.NotFirstResponse && (!formData.AllowMultipleSubmissions && formData.PreviousResponses > 0)  {
			utils.AbortWithStatusJSON(c, http.StatusConflict, "A response has already been recorded from this device.")
			return
		}

		// VALIDATION DONE CREATING SURVEY AND QUESTION RESPONSE 

		var surveySummary databaseSchema.SurveyResponseSummary = databaseSchema.SurveyResponseSummary{
			RespondentEmail:payload.RespondentEmail,
			SurveyId: payload.SurveyId,
		}

		if err:= db.Create(&surveySummary).Error; err != nil {
			c.Error(err)
			return
		}

		// Response summary id obtained preparing for detailed response

		var responses []databaseSchema.SurveyResponseDetails

		for _, rawResponse := range payload.SurveyResponse {

			var response databaseSchema.SurveyResponseDetails = databaseSchema.SurveyResponseDetails{
				QuestionId: rawResponse.QuestionId,
				SummaryId:	surveySummary.Id,
				Response: rawResponse.Response,
			}

			responses = append(responses, response)
		}

		if err:= db.Create(&responses).Error; err!=nil {
			c.Error(err)
			return
		}		


		emailParameter := &structEntities.MailerModel{
			ReceiverEmailId: payload.RespondentEmail,
			Subject:         "Thank You for Filling Out the Form!",
			Body:            fmt.Sprintf(constants.ThankYouEmailTemplate, formData.UserName),
			BodyType:        "html",
		}

		go utils.SendEmail(emailParameter)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Response Saved",
		})

	}
}