package structEntities

type AddTicketPayloadStruct struct {
	BusNumber    string `json:"busNumber" binding:"required"`
	BusInitials  string `json:"busInitials" binding:"required"`
	BusColor     string `json:"busColor" binding:"required"`
	BusRoute     string `json:"busRoute" binding:"required"`
	StartingStop string `json:"startingStop" binding:"required"`
	EndStop      string `json:"endStop" binding:"required"`
	TicketAmount int    `json:"ticketAmount" binding:"required"`
	TicketCount  int    `json:"ticketCount" binding:"required"`
	Discount     float64    `json:"discount"`
	Longitude    string `json:"longitude" binding:"required"`
	Latitude     string `json:"latitude" binding:"required"`
}

type ColorsPayloadStruct struct {
	Colors []string `json:"colors" binding:"required"`
}

type InitialsPayloadStruct struct {
	Initials []string `json:"initials" binding:"required"`
}

type StopsPayloadStruct struct {
	Stops []string `json:"stops" binding:"required"`
}

type RoutesPayloadStruct struct {
	Routes []RouteStruct `json:"routes" binding:"required"`
}

type RouteStruct struct {
	Route     string `json:"route" binding:"required"`
	TerminalA string `json:"terminalA" binding:"required"`
	TerminalB string `json:"terminalB" binding:"required"`
}
