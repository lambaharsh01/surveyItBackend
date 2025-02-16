package structEntities

type TicketingPermissionAccess struct {
	TicketGenerationStatus bool   `json:"ticketGenerationStatus"`
	UserEmail              string `json:"userEmail" binding:"required"`
}

type PaginationStruct struct {
	Page   int `json:"page"`
	Offset int `json:"offset"`
}

// type UsersDetails struct {
// 	Name                   string `json:"name"`
// 	Email                  string `json:"email"`
// 	PhoneNumber            string `json:"phoneNumber"`
// 	Gender                 string `json:"gender"`
// 	UserType               string `json:"userType"`
// 	TicketGenerationStatus int    `json:"ticketGenerationStatus"`
// 	CreatedAt              string `json:"createdAt"`
// }
