package databaseSchema

type QuestionType struct {
	Id                uint   `json:"id" gorm:"primarykey;"`
	QuestionType      string `json:"questionType" gorm:"type:varchar(20);not null"`
	QuestionTypeLabel string `json:"questionTypeLabel" gorm:"type:varchar(100);not null"`
}

var DefaultQuestionTypes []QuestionType = []QuestionType{
	{Id: 1, QuestionType: "star", QuestionTypeLabel: "Star Rating(5)"},
	{Id: 2, QuestionType: "radio", QuestionTypeLabel: "Radio"},
	{Id: 3, QuestionType: "select", QuestionTypeLabel: "Dropdown"},
	{Id: 4, QuestionType: "checkbox", QuestionTypeLabel: "Checkbox"},
	{Id: 5, QuestionType: "file", QuestionTypeLabel: "File"},
	{Id: 6, QuestionType: "text", QuestionTypeLabel: "Text"},
	{Id: 7, QuestionType: "number", QuestionTypeLabel: "Number"},
	{Id: 8, QuestionType: "textarea", QuestionTypeLabel: "Long Text"},
	{Id: 9, QuestionType: "phone", QuestionTypeLabel: "Phone Number"},
	{Id: 10, QuestionType: "email", QuestionTypeLabel: "Email"},
}
