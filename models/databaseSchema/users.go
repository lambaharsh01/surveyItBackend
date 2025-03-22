package databaseSchema

import "time"

type Users struct {
	Id          			uint   	  `json:"id" gorm:"primarykey;autoIncrement"`
	Name        			string 	  `json:"name" gorm:"type:varchar(100);not null"`
	Email       			string 	  `json:"email" gorm:"unique;type:varchar(100);not null"`
	PhoneNumber 			string 	  `json:"phone" gorm:"type:varchar(15);not null"`
	DateOfBirth            	time.Time `json:"dateOfBirth" gorm:"type:date NULL"`
	Gender                 	string    `json:"gender" gorm:"type:varchar(10);not null"`
	Password               	string    `json:"password" gorm:"type:varchar(400)"`
	UserType               	string    `json:"userType" gorm:"type:varchar(10)"`
	PasswordLastUpdatedAt  	time.Time `json:"passwordLastUpdatedAt" gorm:"type:timestamp"`
	TicketGenerationStatus 	int       `json:"ticketGenerationStatus" gorm:"type:int; default:200"`
	OTP              		string    `json:"otp" gorm:"type:varchar(400);"`
	OtpSentAt 				time.Time `json:"otpSentAt" gorm:"type:timestamp"`
	InitialOtpSentAt 		time.Time `json:"initialOtpSentAt" gorm:"type:timestamp"`
	OtpCount         		int       `json:"otpCount" gorm:"type:int; default 0"`
	CreatedAt        		time.Time `json:"createdAt" gorm:"autoCreateTime"`
}


// Email string `json:"email" gorm:"uniqueIndex:idx_email_phone;type:varchar(100)"`
// PhoneNumber string `json:"phone" gorm:"uniqueIndex:idx_email_phone;type:varchar(100)"`
// Otp string `json:"otp" gorm:"type:char(64) null; default:NULL`Removed null; from Otp and OtpSentAt fields: default:NULL alone is enough to allow NULL values.

// INSERT INTO your_table (otp_column)
// VALUES (SHA2('123456', 256));

// SELECT CASE WHEN otp_column = SHA2('123456', 256) THEN 'Match' ELSE 'No Match' END AS otp_status
// FROM your_table
// WHERE some_condition;

// import "time"

// func ParseDateOfBirth(dob string) (time.Time, error) {
//     return time.Parse("2006-01-02", dob) // Go’s date layout format
// }
// Use the parsed date in your code where you handle the request:

// go
// Copy code
// user := models.Users{
//     Name:        "John Doe",
//     Email:       "john@example.com",
//     PhoneNumber: "1234567890",
//     Gender:      "male",
// }

// dob, err := ParseDateOfBirth("2024-05-28")
// if err != nil {
//     // handle error
// }
// user.DateOfBirth = dob

// // Save user to the database
// db.Create(&user)
