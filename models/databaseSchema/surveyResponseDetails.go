package databaseSchema

type SurveyResponseDetails struct {
	Id         uint   `json:"id" gorm:"primarykey;autoIncrement;type:bigint unsigned"`
	QuestionId uint   `json:"questionId" gorm:"type:bigint unsigned;not null;index"`
	SummaryId  uint   `json:"summaryId" gorm:"type:bigint unsigned;not null;index"`
	Response   string `json:"response" gorm:"type:text"`

	QuestionSchema        QuestionSchema        `gorm:"foreignKey:QuestionId;references:Id;not null"`
	SurveyResponseSummary SurveyResponseSummary `gorm:"foreignKey:SummaryId;references:Id;not null"`
}
