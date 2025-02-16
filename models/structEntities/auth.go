package structEntities

import "time"

type AuthToken struct {
	UserId 		  			uint `json:"userId"`
	UserEmail 				string	`json:"userEmail"`
	UserName    			string	`json:"userName"`
	UserGender   			string	`json:"userGender"`
	UserType 				string	`json:"userType"`
	TicketGenerationStatus 	int	`json:"ticketGenerationStatus"`
}

type LoginPayloadStruct struct {
	UserEmail string `json:"userEmail"`
	Password  string `json:"password"`
}

type LoginUserDataResponseStruct struct {
	ID 		 uint
	Name     string
	Email    string
	Password string
	Gender   string
	UserType string
	TicketGenerationStatus int
}

type ForgotOtpPayloadStruct struct {
	UserEmail string `json:"userEmail"`
}

type ConfirmOtpResponseStruct struct {
	OTP              string
	OtpSentAt        time.Time
	InitialOtpSentAt time.Time
	OtpCount         int
}

type InitSignUpPayloadStruct struct {
	UserName    string `json:"userName"`
	UserEmail   string `json:"userEmail"`
	PhoneNumber string `json:"phoneNumber"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
}

type OtpPayloadStruct struct {
	UserEmail string `json:"userEmail"`
	OTP       string `json:"otp"`
}

type ChangePasswordPayloadStruct struct {
	UserEmail string `json:"userEmail"`
	OTP       string `json:"otp"`
	Password  string `json:"password"`
}
