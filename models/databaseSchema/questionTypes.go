package databaseSchema

type QuestionType struct {
	Id                 uint   `json:"id" gorm:"primarykey;"`
	QuestionAcceptedAs string `json:"questionAcceptedAs" gorm:"type:varchar(20);not null"`
	QuestionTypeLabel  string `json:"questionTypeLabel" gorm:"type:varchar(100);not null"`
}

var DefaultQuestionTypes []QuestionType = []QuestionType{
	{Id: 1, QuestionAcceptedAs: "radio", QuestionTypeLabel: "Star Rating"},
	{Id: 2, QuestionAcceptedAs: "radio", QuestionTypeLabel: "Radio"},
	{Id: 3, QuestionAcceptedAs: "select", QuestionTypeLabel: "Dropdown"},
	{Id: 4, QuestionAcceptedAs: "checkbox", QuestionTypeLabel: "Checkbox"},
	{Id: 5, QuestionAcceptedAs: "file", QuestionTypeLabel: "File"},
	{Id: 6, QuestionAcceptedAs: "text", QuestionTypeLabel: "Text"},
	{Id: 7, QuestionAcceptedAs: "number", QuestionTypeLabel: "Number"},
}
