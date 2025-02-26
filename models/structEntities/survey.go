package structEntities

type PaginationStruct struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

type UpsertSurveyPayloadStruct struct {
	SurveyName               string `json:"surveyName" binding:"required"`
	SurveyDescription        string `json:"surveyDescription"`
	SurveyTargetAudience     string `json:"surveyTargetAudience" binding:"required"`
	SurveyAlignment          string `json:"surveyAlignment" binding:"required"`
	SurveyColorTheme         string `json:"surveyColorTheme" binding:"required"`
	AllowMultipleSubmissions bool   `json:"allowMultipleSubmissions"`
	ActiveFrom               string `json:"activeFrom" binding:"required"`
	ActiveTo                 string `json:"activeTo" binding:"required"`
}

type SurveysResponseStruct struct {
	Id                       uint   `json:"id"`
	SurveyCode               string `json:"surveyCode"`
	SurveyName               string `json:"surveyName"`
	SurveyDescription        string `json:"surveyDescription"`
	SurveyTargetAudience     string `json:"surveyTargetAudience"`
	SurveyAlignment          string `json:"surveyAlignment"`
	SurveyColorTheme         string `json:"surveyColorTheme"`
	AllowMultipleSubmissions bool   `json:"allowMultipleSubmissions"`
	ActiveFrom               string `json:"activeFrom"`
	ActiveTo                 string `json:"activeTo"`
	Active                   bool   `json:"active"`
	CreatedAt                string `json:"createdAt"`
}

type QuestionResponseStruct struct {
	ID             uint     `json:"id"`
	Text           string   `json:"text"`
	QuestionTypeID uint     `json:"questionTypeId"`
	QuestionType   string   `json:"questionType"`
	FileTypeID     *uint    `json:"fileTypeId"`
	FileType       *string  `json:"fileType"`
	Options        []string `json:"options"`
	Required       bool     `json:"required"`
	Validation     string   `json:"validation"`
	Min            int      `json:"min"`
	Max            int      `json:"max"`
}

type QuestionaryPayloadStructs struct {
	Id             uint     `json:"id"`
	Text           string   `json:"text" binding:"required"`
	QuestionTypeId uint     `json:"questionTypeId" binding:"required"`
	FileTypeId     *uint    `json:"fileTypeId"`
	Options        []string `json:"options"`
	Required       bool     `json:"required"`
	Validation     bool     `json:"validation"`
	Min            float64  `json:"min"`
	Max            float64  `json:"max"`
}

type AddQuestionaryPayloadStruct struct {
	SurveyId           uint                        `json:"SurveyId" binding:"required"`
	Questionary        []QuestionaryPayloadStructs `json:"questionary" binding:"required"`
	DeletedQuestionIds []uint                      `json:"deletedQuestionIds" binding:"required"`
}
