package structEntities

type PaginationStruct struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

type UpsertSurveyPayloadStruct struct {
	SurveyName           string `json:"surveyName" binding:"required"`
	SurveyDescription    string `json:"surveyDescription" binding:"required"`
	SurveyTargetAudience string `json:"surveyTargetAudience" binding:"required"`
	SurveyAlignment      string `json:"surveyAlignment" binding:"required"`
	SurveyColorTheme     string `json:"surveyColorTheme" binding:"required"`
	ActiveFrom           string `json:"activeFrom" binding:"required"`
	ActiveTo             string `json:"activeTo" binding:"required"`
}

type SurveysResponseStruct struct {
	Id                   uint   `json:"id"`
	SurveyCode           string `json:"surveyCode"`
	SurveyName           string `json:"surveyName"`
	SurveyDescription    string `json:"surveyDescription"`
	SurveyTargetAudience string `json:"surveyTargetAudience"`
	SurveyAlignment      string `json:"surveyAlignment"`
	SurveyColorTheme     string `json:"surveyColorTheme"`
	ActiveFrom           string `json:"activeFrom"`
	ActiveTo             string `json:"activeTo"`
	Active               bool   `json:"active"`
	CreatedAt            string `json:"createdAt"`
}
