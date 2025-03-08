package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lambaharsh01/surveyItBackend/models/databaseSchema"
	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
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