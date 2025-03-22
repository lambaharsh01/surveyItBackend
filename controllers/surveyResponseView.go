package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lambaharsh01/surveyItBackend/models/databaseSchema"
	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
	"github.com/lambaharsh01/surveyItBackend/utils"
)


func GetSurveysAndResponses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.PaginationStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		var total int64
		var limit int = payload.Offset
		var offset int = (payload.Offset * payload.Page)

		user := utils.GetRequestParameters(c)

		surveyResponses := []structEntities.SurveyResponsesResponseStruct{}

		if err := db.Raw(`SELECT 
			ss.id AS survey_id, 
			ss.survey_code, 
			ss.survey_name, 
			CASE WHEN CURRENT_DATE BETWEEN ss.active_from AND ss.active_to THEN 1 ELSE 0 END AS active,
			COUNT(DISTINCT srs.id) AS responses
			FROM survey_schemas ss
			LEFT JOIN survey_response_summaries srs ON 
			srs.survey_id = ss.id
			WHERE created_by = ? AND deleted_at IS NULL
			GROUP BY ss.id, ss.survey_code, ss.survey_name, ss.active_from, ss.active_to
			ORDER BY id DESC 
			LIMIT ? OFFSET ?`, user.UserId, limit, offset).Scan(&surveyResponses).Error; err != nil {
			c.Error(err)
			return
		}
		
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    surveyResponses,
			"total":   total,
			"from":    offset,
			"to":      offset + limit,
		})

	}
}


func GetResponses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var surveyCode string= c.Param("surveyCode")

		user := utils.GetRequestParameters(c)

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
			Where("survey_code = ? AND created_by = ? AND deleted_at IS NULL", surveyCode, user.UserId).
			Order("id DESC").
			First(&survey).Error; err != nil {
			c.Error(err)
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
		WHERE ss.survey_code = ? AND qs.deleted_at IS NULL AND created_by = ?`

		if err := db.Raw(getQuestionsQuery, surveyCode, user.UserId).Scan(&questionaryRaw).Error; err != nil {
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















// SELECT
// 	DENSE_RANK() OVER (ORDER BY srs.id) AS response_sno,
// 	srs.respondent_email,
// 	qs.question,
//     srd.response,
//     srs.created_at AS submitted_at
// FROM survey_schemas ss
// JOIN question_schemas qs ON qs.survey_id = ss.id
// JOIN survey_response_summaries srs ON srs.survey_id = ss.id
// JOIN survey_response_details srd ON
// 	srd.question_id = qs.id AND
//     srd.summary_id = srs.id
// WHERE 1=1
// ORDER BY srs.id, qs.id


// SELECT 
//     ss.survey_name, 
//     DATE_FORMAT(ss.created_at, '%d-%m-%Y') AS created_date,
//     CASE WHEN CURRENT_DATE BETWEEN ss.active_from AND ss.active_to THEN 1 ELSE 0 END AS active,
//     SUM(CASE WHEN srs.id IS NOT NULL THEN 1 ELSE 0 END) AS responses
// FROM survey_schemas ss
// LEFT JOIN survey_response_summaries srs ON srs.survey_id = ss.id
// -- WHERE CURRENT_DATE BETWEEN ss.active_from AND ss.active_to
// GROUP BY ss.survey_name, ss.created_at, ss.active_from, ss.active_to
// ORDER BY active DESC, ss.id ASC
