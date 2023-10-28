package model

import "gorm.io/gorm"

type AirportRequest struct {
	gorm.Model
	AirportCode  string  `json:"airport_code" validate:"max=5,required,aplha"`
	AirportName  string  `json:"airport_name" validate:"max=50,required,alpha"`
	City         string  `json:"city" validate:"max=20,required,alpha"`
	Country      string  `json:"country" validate:"max=20,required,alpha"`
	Region       string  `json:"region" validate:"max=50,required,alpha"`
	Latitude     float64 `json:"latitude" valdiate:"required,numeric,latitude,latitude_format=ddmmss"`
	Longitude    float64 `json:"longitude" validate:"required,numeric,longitude,longitude_format=ddmmss"`
	IATAFCSCode  string  `json:"iata_fcs_code" validate:"required,alpha,len=3"`
	ICAOCode     string  `json:"icao_code" validate:"required,alpha,len=4"`
	Website      string  `json:"website" validate:"required,url"`
	ContactEmail string  `json:"contact_email" validate:"required,email"`
	ContactPhone string  `json:"contact_phone" validate:"required,phone"`
	BlackListed  string  `json:"black_listed" validate:"required,boolean"`
}

type ScheduleRequest struct {
	DepartureTime    string `json:"departure_time" validate:"required,datetime"`
	ArrivalTime      string `json:"arrival_time" validate:"required,datetime"`
	DepartureAirport string `json:"departure_airport" validate:"required,alpha,len=3"`
	ArrivalAirport   string `json:"arrival_airport" validate:"required,alpha,len=3"`
	Scheduled        string `json:"scheduled" validate:"required,boolean"`
}
