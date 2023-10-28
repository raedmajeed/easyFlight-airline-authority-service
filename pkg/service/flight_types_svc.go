package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
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
