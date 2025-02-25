package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lambaharsh01/surveyItBackend/models/databaseSchema"
	"github.com/lambaharsh01/surveyItBackend/models/structEntities"
	"github.com/lambaharsh01/surveyItBackend/utils"
)

func GetQuestionTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		questionType := []databaseSchema.QuestionType{}

		if err := db.Find(&questionType).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":      true,
			"questionType": questionType,
		})

	}
}

func GetFileTypes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		fileType := []databaseSchema.FileType{}

		if err := db.Find(&fileType).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":  true,
			"fileType": fileType,
		})

	}
}

func AddSurvey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.UpsertSurveyPayloadStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		user := utils.GetRequestParameters(c)

		var survey databaseSchema.SurveySchema = databaseSchema.SurveySchema{
			SurveyName:           payload.SurveyName,
			SurveyDescription:    payload.SurveyDescription,
			SurveyTargetAudience: payload.SurveyTargetAudience,
			SurveyAlignment:      payload.SurveyAlignment,
			SurveyColorTheme:     payload.SurveyColorTheme,
			AllowMultipleSubmissions:   payload.AllowMultipleSubmissions,
			DeletedAt:            nil,
		}

		if err := utils.ParseDate(payload.ActiveFrom, &survey.ActiveFrom); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active from date")
			return
		}

		if err := utils.ParseDate(payload.ActiveTo, &survey.ActiveTo); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active to date")
			return
		}

		survey.SurveyCode = utils.GenerateUniqueKey(15)
		survey.CreatedBy = user.UserId

		if err := db.Create(&survey).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "survey created successfully",
		})

	}
}

func UpdateSurvey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var id string = c.Param("id")
		user := utils.GetRequestParameters(c)

		var payload structEntities.UpsertSurveyPayloadStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		var updates map[string]interface{} = map[string]interface{}{
			"survey_name":            payload.SurveyName,
			"survey_description":     payload.SurveyDescription,
			"survey_target_audience": payload.SurveyTargetAudience,
			"survey_alignment":       payload.SurveyAlignment,
			"survey_color_theme":     payload.SurveyColorTheme,
			"allow_multiple_submissions":     payload.AllowMultipleSubmissions,
		}

		var activeFrom time.Time
		var activeTo time.Time
		if err := utils.ParseDate(payload.ActiveFrom, &activeFrom); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active from date")
			return
		}

		if err := utils.ParseDate(payload.ActiveTo, &activeTo); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active to date")
			return
		}

		updates["active_from"] = activeFrom
		updates["active_to"] = activeTo

		if err := db.Model(&databaseSchema.SurveySchema{}).
			Where("id = ? AND created_by = ?", id, user.UserId).
			Updates(updates).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "survey created successfully",
		})

	}
}

func GetSurvey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var surveyCode string= c.Param("surveyCode")

		user := utils.GetRequestParameters(c)

		survey := structEntities.SurveysResponseStruct{}

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

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    survey,
		})

	}
}

func GetSurveys(db *gorm.DB) gin.HandlerFunc {
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

		surveys := []structEntities.SurveysResponseStruct{}

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
			DATE_FORMAT(created_at, '%D %M %Y') AS created_at
			`).
			Where("created_by = ? AND deleted_at IS NULL", user.UserId).
			Order("id DESC").
			Limit(limit).
			Offset(offset).Find(&surveys).Error; err != nil {
			c.Error(err)
			return
		}

		if err := db.Model(&databaseSchema.SurveySchema{}).
			Where("created_by = ? AND deleted_at IS NULL", user.UserId).
			Count(&total).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    surveys,
			"total":   total,
			"from":    offset,
			"to":      offset + limit,
		})

	}
}

func DeleteSurvey(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var id string = c.Param("id")

		if err := db.Model(&databaseSchema.SurveySchema{}).
			Where("id = ?", id).
			Update("deleted_at", time.Now()).
			Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})

	}
}


func SaveQuestionary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.UpsertSurveyPayloadStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}

		user := utils.GetRequestParameters(c)

		var survey databaseSchema.SurveySchema = databaseSchema.SurveySchema{
			SurveyName:           payload.SurveyName,
			SurveyDescription:    payload.SurveyDescription,
			SurveyTargetAudience: payload.SurveyTargetAudience,
			SurveyAlignment:      payload.SurveyAlignment,
			SurveyColorTheme:     payload.SurveyColorTheme,
			DeletedAt:            nil,
		}

		if err := utils.ParseDate(payload.ActiveFrom, &survey.ActiveFrom); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active from date")
			return
		}

		if err := utils.ParseDate(payload.ActiveTo, &survey.ActiveTo); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, "invalid active to date")
			return
		}

		survey.SurveyCode = utils.GenerateUniqueKey(15)
		survey.CreatedBy = user.UserId

		if err := db.Create(&survey).Error; err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "survey created successfully",
		})

	}
}


func UpdateQuestionary(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var payload structEntities.AddQuestionaryPayloadStruct
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.AbortWithStatusJSON(c, http.StatusBadRequest, err.Error())
			return
		}
		
		transaction := db.Begin()
		if transaction.Error !=nil {
			c.Error(transaction.Error)
			return
		}

		for _, question := range payload.Questionary {

			questionSchema := databaseSchema.QuestionSchema{
				Id: question.Id,
				Question: question.Text,
				Options: question.Options,
				Required: question.Required,
				Validation: question.Validation,
				Min: question.Min,
				Max: question.Max,
				FormId: payload.FormId,
				QuestionTypeId: question.QuestionTypeId,
				FileTypeId: question.FileTypeId,
			}

			if err := transaction.Save(&questionSchema).Error; err!= nil {
				transaction.Rollback()
				c.Error(err)
				return
			}
		}


		if len(payload.DeletedQuestionIds) > 0 {

			var setDeletedAt map[string]interface{} = map[string]interface{}{
				"deleted_at":   time.Now(),
			}

			if err := db.Model(&databaseSchema.QuestionSchema{}).
				Where("id IN ?", payload.DeletedQuestionIds).
				Updates(setDeletedAt).Error; err!=nil {
					transaction.Rollback()
					c.Error(err)
					return
			}
		}
		
		if err:= transaction.Commit().Error; err !=nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "questionary updated successfully",
		})
	}
}