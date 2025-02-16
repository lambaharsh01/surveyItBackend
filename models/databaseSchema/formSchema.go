package databaseSchema

import "time"

type FormSchema struct {
	Id           	uint   		`json:"id" gorm:"primarykey;autoIncrement"`
	SurveyId 		string 		`json:"surveyId" gorm:"type:varchar(50);unique;not null"`
	SurveyName 		string 		`json:"surveyName" gorm:"type:varchar(100);not null"`
	SurveyMessage 	string 		`json:"surveyMessage" gorm:"type:varchar(500)"`
	CreatedAt    	time.Time 	`json:"createdAt" gorm:"autoCreateTime"`
	DeletedAt    	*time.Time 	`json:"deletedAt,omitempty" gorm:"type:timestamp;null"`
	CreatedBy 	 	uint    	`json:"createdBy" gorm:"type:int;not null"`
	Users 		 	Users     	`gorm:"foreignkey:CreatedBy;references:Id"`
}

