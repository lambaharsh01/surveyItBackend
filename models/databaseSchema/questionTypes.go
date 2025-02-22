package databaseSchema

type QuestionType struct {
	Id                uint   `json:"id" gorm:"primarykey;"`
	QuestionType      string `json:"questionType" gorm:"type:varchar(20);not null"`
	QuestionTypeLabel string `json:"questionTypeLabel" gorm:"type:varchar(100);not null"`
}

var DefaultQuestionTypes []QuestionType = []QuestionType{
	{Id: 1, QuestionType: "text", QuestionTypeLabel: "Text"},
	{Id: 2, QuestionType: "email", QuestionTypeLabel: "Email"},
	{Id: 3, QuestionType: "textarea", QuestionTypeLabel: "Long Text"},
	{Id: 4, QuestionType: "number", QuestionTypeLabel: "Number"},
	{Id: 5, QuestionType: "date", QuestionTypeLabel: "Date"},
	{Id: 6, QuestionType: "star", QuestionTypeLabel: "Star Rating(5)"},
	{Id: 7, QuestionType: "radio", QuestionTypeLabel: "Radio"},
	{Id: 9, QuestionType: "checkbox", QuestionTypeLabel: "Checkbox"},
	{Id: 8, QuestionType: "select", QuestionTypeLabel: "Dropdown"},
	{Id: 10, QuestionType: "file", QuestionTypeLabel: "File"},
}
