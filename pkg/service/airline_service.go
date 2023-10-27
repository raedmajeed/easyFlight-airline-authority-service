package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

//* METHODS TO EVERYTHING FLIGHT TYPES

func (svc *AdminAirlineServiceStruct) CreateFlightType(p *pb.FlightTypeRequest) (*dom.FlightTypeModel, error) {
	flightType, err := svc.repo.FindFlightTypeByModel(p.FlightModel)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of model %v", p.FlightModel)
		} else {
			log.Printf("Flight Type not create of model %v, err: %v", p.FlightModel, err.Error())
			return flightType, err
		}
	}

	flightType, err = svc.repo.CreateFlightType(p)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				p.FlightModel, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return flightType, nil
}

func (svc *AdminAirlineServiceStruct) GetAllFlightTypes() ([]dom.FlightTypeModel, error) {
	flightTypes, err := svc.repo.FindAllFlightTypes()
	if err != nil {
		log.Printf("Unable to get all the flight types, err: %v", err.Error())
		return nil, err
	}
	return flightTypes, nil
}

func (svc *AdminAirlineServiceStruct) GetFlightType(id int32) (*dom.FlightTypeModel, error) {
	flightType, err := svc.repo.FindFlightTypeByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of model %v", id)
		} else {
			log.Printf("Flight Type not create of model %v, err: %v", id, err.Error())
			return flightType, err
		}
	}
	return flightType, nil
}

func (svc *AdminAirlineServiceStruct) UpdateFlightType(p *pb.FlightTypeRequest, id int) (*dom.FlightTypeModel, error) {
	existingFlightType, err := svc.repo.FindFlightTypeByID(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of model %v", id)
		} else {
			log.Printf("Flight Type not updated of model %v, err: %v", p.FlightModel, err.Error())
			return existingFlightType, err
		}
	}
	if existingFlightType != nil {
		if p.FlightModel != "" {
			existingFlightType.FlightModel = p.FlightModel
		}
		if p.Description != "" {
			existingFlightType.Description = p.Description
		}
		if p.ManufacturerName != "" {
			existingFlightType.ManufacturerName = p.ManufacturerName
		}
		if p.ManufacturerCountry != "" {
			existingFlightType.ManufacturerCountry = p.ManufacturerCountry
		}
		if p.MaxDistance < 0 {
			existingFlightType.MaxDistance = p.MaxDistance
		}
		if p.CruiseSpeed < 0 {
			existingFlightType.CruiseSpeed = p.CruiseSpeed
		}

		flightType, err := svc.repo.UpdateFlightType(existingFlightType)
		if err != nil {
			log.Println("Unable to Update the flight types")
			return flightType, err
		}
	}

	ft, err := svc.CreateFlightType(p)
	return ft, err
}

//* METHODS TO EVERYTHING AIRLINE

func (svc *AdminAirlineServiceStruct) RegisterFlight(p *pb.AirlineRequest) (*dom.Airline, error) {
	airline := &dom.Airline{
		AirlineName:         p.AirlineName,
		CompanyAddress:      p.CompanyAddress,
		PhoneNumber:         p.PhoneNumber,
		Email:               p.Email,
		AirlineCode:         p.AirlineCode,
		AirlineLogoLink:     p.AirlineLogoLink,
		SupportDocumentLink: p.SupportDocumentsLink,
	}
	otp := utils.GenerateOTP()
	otpData := dom.OtpData{
		Otp:     otp,
		Email:   airline.Email,
		Airline: *airline,
	}

	otpJson, err := json.Marshal(&otpData)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	svc.redis.Set(context.Background(), "airline_data", otpJson, time.Minute*10)
	return airline, nil
}

// * METHODS TO EVERYTHING AIRLINE SEATS
type Layout struct{ Rows [][]bool }

func (svc *AdminAirlineServiceStruct) CreateAirlineSeats(p *pb.AirlineSeatRequest, id int) (*dom.AirlineSeat, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Flight Type not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	economySeats := p.EconomySeatNo
	economySeatsPerRow := p.EconomySeatsPerRow
	buisinessSeats := p.BuisinesSeatNo
	buisinessSSeatsPerRow := p.BuisinesSeatsPerRow

	ecoLayout := createEconomySeatsJSONLayout(economySeats, economySeatsPerRow)
	buisLayout := createBuisinessSeatsJSONLayout(buisinessSeats, buisinessSSeatsPerRow)

	economyLayoutJSON, err := json.Marshal(ecoLayout)

	if err != nil {
		log.Println("error parsing economy seat layout")
		return nil, err
	}

	buisinessLayoutJSON, err := json.Marshal(buisLayout)
	if err != nil {
		log.Println("error parsing buisiness seat layout")
		return nil, err
	}
	airlineSeats, err := svc.repo.CreateAirlineSeatType(p, economyLayoutJSON, buisinessLayoutJSON)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Seat layout not created of airline %v, err: %v",
				p.AirlineId, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineSeats, nil
}

func createEconomySeatsJSONLayout(seats, sprow int32) *Layout {
	l := &Layout{}

	for i := 1; i <= int(seats); i += int(sprow) {
		row := make([]bool, sprow)
		for j := 1; j <= int(sprow); j++ {
			row = append(row, false)
		}
		l.Rows = append(l.Rows, row)
	}
	return l
}

func createBuisinessSeatsJSONLayout(seats, sprow int32) *Layout {
	l := &Layout{}
	for i := 1; i <= int(seats); i += int(sprow) {
		row := make([]bool, sprow)
		for j := 1; j <= int(sprow); j++ {
			row = append(row, false)
		}
		l.Rows = append(l.Rows, row)
	}
	return l
}

// * METHODS TO EVERYTHING AIRLINE BAGGAGE POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineBaggagePolicy(p *pb.AirlineBaggageRequest, id int) (*dom.AirlineBaggage, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Baggage Policy not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	airlineBaggagePolicy, err := svc.repo.CreateAirlineBaggagePolicy(p, id)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				id, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineBaggagePolicy, nil
}

// * METHODS TO EVERYTHING AIRLINE CANCELATION POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineCancellationPolicy(p *pb.AirlineCancellationRequest, id int) (*dom.AirlineCancellation, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Cancellation policy not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	airlineCancellationPolicy, err := svc.repo.CreateAirlineCancellationPolicy(p, id)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				id, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineCancellationPolicy, nil
}
