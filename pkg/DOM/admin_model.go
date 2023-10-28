package model

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	AirportCode  string  `json:"airport_code" gorm:"unique"`
	AirportName  string  `json:"airport_name" gorm:"unique"`
	City         string  `json:"city"`
	Country      string  `json:"country"`
	Region       string  `json:"region"`
	Latitude     float64 `json:"latitude" gorm:"unique"`
	Longitude    float64 `json:"longitude" gorm:"unique"`
	IATAFCSCode  string  `json:"iata_fcs_code" gorm:"unique"`
	ICAOCode     string  `json:"icao_code" gorm:"unique"`
	Website      string  `json:"website"`
	ContactEmail string  `json:"contact_email"`
	ContactPhone string  `json:"contact_phone"`
	BlackListed  bool    `json:"approved" gorm:"default:false"`
}

type Schedule struct {
	DepartureTime    string `json:"departure_time"`
	ArrivalTime      string `json:"arrival_time"`
	DepartureAirport string `json:"departure_airport"`
	ArrivalAirport   string `json:"arrival_airport"`
	Scheduled        bool   `json:"scheduled" gorm:"default:false"`
}
