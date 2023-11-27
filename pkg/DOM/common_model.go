package model

import "time"

type EmailMessage struct {
	Email   string
	Subject string
	Content string
}

type SearchDetails struct {
	DepartureAirport    string
	ArrivalAirport      string
	DepartureDate       string
	ReturnDepartureDate string
	ReturnFlight        bool
	Economy             bool
	MaxStops            string
}

type FlightChartResponse struct {
	DepartureAirport  string
	ArrivalAirport    string
	FlightNumber      string
	DepartureDateTime time.Time
	ArrivalDateTime   time.Time
	AirlineName       string
	EconomyFare       float64
	BusinessFare      float64
}

type FlightDetails struct {
	FlightChartID     uint
	FlightNumber      string
	Airline           string
	DepartureAirport  string    `column:"dep_airport"`
	ArrivalAirport    string    `column:"arr_airport"`
	DepartureDate     string    `column:"dep_date"`
	ArrivalDate       string    `column:"arr_date"`
	DepartureTime     string    `column:"dep_time"`
	ArrivalTime       string    `column:"arr_time"`
	DepartureDateTime time.Time `column:"dep_datetime"`
	ArrivalDateTime   time.Time `column:"arr_datetime"`
}

type Path struct {
	PathId          int
	Flights         []FlightDetails
	NumberOfStops   int
	TotalTravelTime float64
}
