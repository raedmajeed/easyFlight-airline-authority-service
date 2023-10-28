package model

import (
	"gorm.io/gorm"
)

type FlightTypeModel struct {
	gorm.Model
	Type                string `json:"type" gorm:"not null"`
	FlightModel         string `json:"flight_model" gorm:"not null;unique"`
	Description         string `json:"description" gorm:"not null"`
	ManufacturerName    string `json:"maunufacturer_name" gorm:"not null"`
	ManufacturerCountry string `json:"manufacturer_country" gorm:"not null"`
	MaxDistance         int32  `json:"max_distance" gorm:"not null"`
	CruiseSpeed         int32  `json:"cruise_speed" gorm:"not null"`
}

type Airline struct {
	gorm.Model
	AirlineName         string `json:"name"`
	CompanyAddress      string `json:"company_address"`
	PhoneNumber         string `json:"phone_number"`
	Email               string `json:"email" gorm:"not null;unique"`
	AirlineCode         string `json:"airline_code" gorm:"not null;unique"`
	AirlineLogoLink     string `json:"airline_logo_link"`
	SupportDocumentLink string `json:"support_documents_link"`
}

type AirlineSeat struct {
	gorm.Model
	AirlineId           int    `json:"airline_id"`
	EconomySeatNumber   int    `json:"economy_seat_no"`
	BuisinesSeatNumber  int    `json:"buisines_seat_no"`
	EconomySeatsPerRow  int    `json:"economy_seats_per_row"`
	BuisinesSeatsPerRow int    `json:"buisines_seats_per_row"`
	EconomySeatLayout   []byte `json:"economy_seat_layout"`
	BuisinessSeatLayout []byte `json:"buisiness_seat_layout"`
}

type AirlineBaggage struct {
	gorm.Model
	AirlineId           int    `json:"airline_id"`
	FareClass           int    `json:"class"`
	CabinAllowedWeight  int    `json:"cabin_allowed_weight"`
	CabinAllowedLength  int    `json:"cabin_allowed_length"`
	CabinAllowedBreadth int    `json:"cabin_allowed_breadth"`
	CabinAllowedHeight  int    `json:"cabin_allowed_height"`
	HandAllowedWeight   int    `json:"hand_allowed_weight"`
	HandAllowedLength   int    `json:"hand_allowed_length"`
	HandAllowedBreadth  int    `json:"hand_allowed_breadth"`
	HandAllowedHeight   int    `json:"hand_allowed_height"`
	FeeExtraPerKGCabin  int    `json:"fee_for_extra_kg_cabin"`
	FeeExtraPerKGHand   int    `json:"fee_for_extra_kg_hand"`
	Restrictions        string `json:"restrictions"`
}

type AirlineCancellation struct {
	gorm.Model
	AirlineId                  int  `json:"airline_id"`
	FareClass                  int  `json:"class"`
	CancellationDeadlineBefore int  `json:"cancellation_deadline_before_hours"`
	CancellationPercentage     int  `json:"cancellation_percentage"`
	Refundable                 bool `json:"refundable"`
}

type OtpData struct {
	Otp     int     `json:"otp"`
	Email   string  `json:"email"`
	Airline Airline `json:"airline"`
}

type FlightFleetRequest struct {
	SeatId             int `json:"seat_id" validate:"required,foreign_key:seats"`
	FlightType         int `json:"flight_type" validate:"required,foreign_key:flight_types"`
	BaggagePolicy      int `json:"baggage_policy" validate:"required,foreign_key:baggage_policies"`
	CancellationPolicy int `json:"cancellation_policy" validate:"required,foreign_key:cancellation_policies"`
	FlightCount        int `json:"flight_count" validate:"numeric"`
}
