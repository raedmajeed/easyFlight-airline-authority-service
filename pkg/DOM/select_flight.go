package model

type CompleteFlightFacilities struct {
	DirectFlight     FlightFacilities
	ReturnFlight     FlightFacilities
	NumberOfAdults   int
	NumberOfChildren int
	CabinClass       string
	DepartureAirport string
	ArrivalAirport   string
}

type FlightFacilities struct {
	Cancellation Cancellation
	Baggage      Baggage
	FlightPath   Path
	Fare         float64
}

type Cancellation struct {
	CancellationDeadlineBefore int
	CancellationPercentage     int
	Refundable                 bool
}

type Baggage struct {
	CabinAllowedWeight  int
	CabinAllowedLength  int
	CabinAllowedBreadth int
	CabinAllowedHeight  int
	HandAllowedWeight   int
	HandAllowedLength   int
	HandAllowedBreadth  int
	HandAllowedHeight   int
	FeeExtraPerKGCabin  int
	FeeExtraPerKGHand   int
	Restrictions        string
}
