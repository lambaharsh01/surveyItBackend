package structEntities

type MailerModel struct {
	ReceiverEmailId string 
	Subject string
	Body string
	CC []string
	BCC []string
	BodyType string // plain || html
}