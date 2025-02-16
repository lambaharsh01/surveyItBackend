package databaseSchema

type FileType struct {
	Id            uint   `json:"id" gorm:"primarykey;"`
	FileType      string `json:"fileType" gorm:"type:varchar(100);not null"`
	FileTypeLabel string `json:"fileTypeLabel" gorm:"type:varchar(100);not null"`
}

var DefaultFileTypes []FileType = []FileType{
	{Id: 1, FileType: "*/*", FileTypeLabel: "All File Types"},
	{Id: 2, FileType: "text/plain", FileTypeLabel: "Plain Text"},
	{Id: 3, FileType: "application/pdf", FileTypeLabel: "Portable Document Format (PDF)"},
	{Id: 4, FileType: "application/msword", FileTypeLabel: "Microsoft Word"},
	{Id: 5, FileType: "application/vnd.openxmlformats-officedocument.wordprocessingml.document", FileTypeLabel: "Microsoft Word (New)"},
	{Id: 6, FileType: "application/vnd.ms-excel", FileTypeLabel: "Microsoft Excel"},
	{Id: 7, FileType: "application/vnd.openxmlformats-officedocument.spreadsheet", FileTypeLabel: "Microsoft Excel (New)"},
	{Id: 8, FileType: "text/csv", FileTypeLabel: "Comma-Separated Values"},
	{Id: 9, FileType: "image/png", FileTypeLabel: "Portable Network Graphics"},
	{Id: 10, FileType: "image/jpeg", FileTypeLabel: "JPEG Image"},
	{Id: 11, FileType: "video/mp4", FileTypeLabel: "MP4 Video"},
}