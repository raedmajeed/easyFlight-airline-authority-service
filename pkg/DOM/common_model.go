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
}
