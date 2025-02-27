package databaseSchema

type SurveyResponseDetails struct {
	Id               uint   `json:"id" gorm:"primarykey;autoIncrement;type:bigint unsigned"`
	QuestionSchemaId uint   `json:"questionSchemaId" gorm:"type:bigint unsigned;not null;index"`
	SummaryId        uint   `json:"summaryId" gorm:"type:bigint unsigned;not null;index"`
	Response         string `json:"response" gorm:"type:text"`

	QuestionSchema        QuestionType          `gorm:"foreignKey:QuestionSchemaId;references:Id;not null"`
	SurveyResponseSummary SurveyResponseSummary `gorm:"foreignKey:SummaryId;references:Id;not null"`
}
