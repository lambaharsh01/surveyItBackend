package databaseSchema

import "time"

type QuestionSchema struct {
	Id             uint         `json:"id" gorm:"primarykey;autoIncrement;type:bigint unsigned"`
	Question       string       `json:"surveyMessage" gorm:"type:varchar(500);not null"`
	Options        []string     `json:"options" gorm:"type:json"`
	Required       bool         `json:"required" gorm:"type:tinyint(1);not null;default:0"`
	Validation     bool         `json:"validation" gorm:"type:tinyint(1);not null;default:0"`
	Min            float64      `json:"min" gorm:"type:int;not null;default:0"`
	Max            float64      `json:"max" gorm:"type:int;not null;default:0"`
	SurveyId       uint         `json:"surveyId" gorm:"type:bigint unsigned;not null;index"`
	QuestionTypeId uint         `json:"questionTypeId" gorm:"type:bigint unsigned;not null"`
	FileTypeId     *uint        `json:"fileTypeId" gorm:"type:bigint unsigned"`
	DeletedAt      *time.Time   `json:"deletedAt,omitempty" gorm:"type:timestamp NULL DEFAULT NULL"`
	
	SurveySchema   SurveySchema `gorm:"foreignKey:SurveyId;references:Id;constraint:OnDelete:CASCADE"`
	QuestionType   QuestionType `gorm:"foreignKey:QuestionTypeId;references:Id;constraint:OnDelete:CASCADE"`
	FileType       *FileType    `gorm:"foreignKey:FileTypeId;references:Id;constraint:OnDelete:SET NULL"`
}
