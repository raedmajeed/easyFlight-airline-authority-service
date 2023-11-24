package model

import (
	"gorm.io/gorm"
	"time"
)

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
	gorm.Model
	DepartureTime     string    `json:"departure_time"`
	ArrivalTime       string    `json:"arrival_time"`
	DepartureDate     string    `json:"departure_date"`
	ArrivalDate       string    `json:"arrival_date"`
	DepartureAirport  string    `json:"departure_airport"`
	ArrivalAirport    string    `json:"arrival_airport"`
	DepartureDateTime time.Time `json:"departure_date_time"`
	ArrivalDateTime   time.Time `json:"arrival_date_time"`
	Scheduled         bool      `json:"scheduled" gorm:"default:false"`
}

type AdminTable struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type BookedSeat struct {
	gorm.Model
	FlightChartNo      int
	AirlineId          int    `json:"airline_id"`
	EconomySeatTotal   int    `json:"economy_seat_total_no"`
	BusinessSeatTotal  int    `json:"business_seat_total"`
	EconomySeatBooked  int    `json:"economy_seat_no"`
	BusinessSeatBooked int    `json:"business_seat_no"`
	EconomySeatLayout  []byte `json:"economy_seat_layout"`
	BusinessSeatLayout []byte `json:"business_seat_layout"`
}

type Status int

const (
	CONFIRMED Status = iota
	DELAYED
	SCHEDULED
)

type FlightChart struct {
	gorm.Model
	FlightNumber string       `gorm:"not null"`
	FlightID     uint         `gorm:"not null"`
	Flight       FlightFleets `gorm:"foreignKey:FlightID"`
	Status       Status       `gorm:"default:0"`
	ScheduleID   uint         `gorm:"not null"`
	Schedule     Schedule     `gorm:"foreignKey:ScheduleID"`
	EconomyFare  float64
	BusinessFare float64
}

type CombinedChartScheduleFleet struct {
	FlightChart
	Schedule
	FlightFleets
}
