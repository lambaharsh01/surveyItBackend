package databaseSchema

import "time"

type SurveySchema struct {
	Id                   uint       `json:"id" gorm:"primarykey;autoIncrement"`
	SurveyCode           string     `json:"surveyCode" gorm:"type:varchar(50);unique;not null"`
	SurveyName           string     `json:"surveyName" gorm:"type:varchar(100);not null"`
	SurveyDescription    string     `json:"surveyDescription" gorm:"type:varchar(500)"`
	SurveyTargetAudience string     `json:"surveyTargetAudience" gorm:"type:varchar(100)"`
	SurveyAlignment      string     `json:"surveyAlignment" gorm:"type:varchar(50)"`
	SurveyColorTheme     string     `json:"surveyColorTheme" gorm:"type:varchar(50)"`
	AllowMultipleSubmissions     	 bool     	`json:"allowMultipleSubmissions" gorm:"type:tinyint(1);not null;default:0"`
	ActiveFrom           time.Time  `json:"activeFrom" gorm:"autoCreateTime"`
	ActiveTo             time.Time  `json:"activeTo" gorm:"autoCreateTime"`
	CreatedBy            uint       `json:"createdBy" gorm:"type:int;not null"`
	Users                Users      `gorm:"foreignkey:CreatedBy;references:Id"`
	CreatedAt            time.Time  `json:"createdAt" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt            *time.Time `json:"deletedAt,omitempty" gorm:"type:timestamp NULL DEFAULT NULL"`
}
