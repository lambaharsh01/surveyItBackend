package databaseSchema

// import (
// 	"gorm.io/datatypes"
// )

type QuestionSchema struct {
	
	Id              uint           	`json:"id" gorm:"primarykey;autoIncrement;type:bigint unsigned"`
	Question        string         	`json:"surveyMessage" gorm:"type:varchar(500);not null"`
	ResponseOptions string 			`json:"responseOptions" gorm:"type:json"`
	Required        bool           	`json:"required" gorm:"type:tinyint(1);not null;default:0"`
	Validation      bool           	`json:"validation" gorm:"type:tinyint(1);not null;default:0"`
	Min        		int           	`json:"min" gorm:"type:int;not null;default:0"`
	Max        		int           	`json:"max" gorm:"type:int;not null;default:0"`

	FormId         	uint 			`json:"formId" gorm:"type:bigint unsigned;not null"`
	QuestionTypeId 	uint 			`json:"questionTypeId" gorm:"type:bigint unsigned;not null"`
	FileTypeId     	uint 			`json:"fileTypeId" gorm:"type:bigint unsigned"`

	SurveySchema   	SurveySchema   	`gorm:"foreignKey:FormId;references:Id;constraint:OnDelete:CASCADE"`
	QuestionType 	QuestionType 	`gorm:"foreignKey:QuestionTypeId;references:Id;constraint:OnDelete:CASCADE"`
	FileType     	FileType     	`gorm:"foreignKey:FileTypeId;references:Id;constraint:OnDelete:SET NULL"`
}
