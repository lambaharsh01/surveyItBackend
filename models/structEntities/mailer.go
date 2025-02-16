package structEntities

type MailerModel struct {
	ReciverEmailId string 
	Subject string
	Body string
	CC []string
	BCC []string
	BodyType string // plain || html
}