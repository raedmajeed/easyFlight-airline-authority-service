package model

type KafkaPath struct {
	DirectPath       []Path
	ReturnPath       []Path
	DepartureAirport string
	ArrivalAirport   string
}

type SelectRequest struct {
	Token        string
	DirectPathId string
	ReturnPathId string
	Adults       int
	Children     int
	Economy      bool
}

type SearchClaims struct {
	Adults        int
	Children      int
	Economy       bool
	PassengerType string
}

type TemporaryData struct {
	ChartId uint
	Seats   *BookedSeat
}
