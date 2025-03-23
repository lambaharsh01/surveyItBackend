package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
	"github.com/lambaharsh01/surveyItBackend/utils"
)

func GetResponseData(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var surveyCode string = c.Param("surveyCode")

		user := utils.GetRequestParameters(c)

		surveyResponses := []structEntities.SurveyResponsesResponseStruct{}

		if err := db.Raw(`
		
		SELECT
			DENSE_RANK() OVER (ORDER BY srs.id) AS response_sno,
			srs.respondent_email,
			qs.question,
			srd.response,
			srs.created_at AS submitted_at
		FROM survey_schemas ss
		JOIN question_schemas qs ON qs.survey_id = ss.id
		JOIN survey_response_summaries srs ON srs.survey_id = ss.id
		JOIN survey_response_details srd ON
			srd.question_id = qs.id AND
			srd.summary_id = srs.id
		WHERE ss.survey_code = ? AND created_by = ? AND ss.deleted_at IS NULL
		ORDER BY srs.id, qs.id
		`, surveyCode, user.UserId).Scan(&surveyResponses).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    surveyResponses,
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
