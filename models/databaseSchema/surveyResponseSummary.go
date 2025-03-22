package databaseSchema

import (
	"time"
)

// SurveyResponseDetails
type SurveyResponseSummary struct {
	Id             uint         `json:"id" gorm:"primarykey;autoIncrement;type:bigint unsigned"`
	RespondentEmail     string       `json:"respondentEmail" gorm:"type:varchar(50);not null;index"`
	SurveyId       uint         `json:"surveyId" gorm:"type:bigint unsigned;not null;index"`
	CreatedAt      time.Time  	`json:"createdAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	SurveySchema   SurveySchema `gorm:"foreignKey:SurveyId;references:Id;not null"`
	DeletedAt      *time.Time 	`json:"deletedAt,omitempty" gorm:"type:timestamp NULL DEFAULT NULL"`
}
