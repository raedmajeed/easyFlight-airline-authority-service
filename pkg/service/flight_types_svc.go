package service

import (
	"errors"
	"fmt"
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

	if flightType != nil {
		log.Printf("flight type already exists of type %v", flightType.FlightModel)
		return flightType, fmt.Errorf("flight type already exists of type %v", flightType.FlightModel)
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
			log.Printf("no existing record found  of model %v", id)
			return nil, err
		} else {
			log.Printf("flight Type not created of model %v, err: %v", id, err.Error())
			return flightType, err
		}
	}
	return flightType, nil
}

func (svc *AdminAirlineServiceStruct) UpdateFlightType(p *pb.FlightTypeRequest, id int) (*dom.FlightTypeModel, error) {
	existingFlightType, err := svc.repo.FindFlightTypeByID(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found of model %v", id)
		} else {
			log.Printf("Error finding Flight Type of model %v, err: %v", p.FlightModel, err.Error())
		}
		return existingFlightType, err
	}

	if existingFlightType == nil {
		return nil, errors.New("flight type not found")
	}

	updateFields(p, existingFlightType)

	flightType, err := svc.repo.UpdateFlightType(existingFlightType, id)
	if err != nil {
		log.Println("Unable to Update the flight types")
		return nil, err
	}

	return flightType, nil
}

func updateFields(p *pb.FlightTypeRequest, existingFlightType *dom.FlightTypeModel) {
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
	if p.MaxDistance >= 0 {
		existingFlightType.MaxDistance = p.MaxDistance
	}
	if p.CruiseSpeed >= 0 {
		existingFlightType.CruiseSpeed = p.CruiseSpeed
	}
}

func (svc *AdminAirlineServiceStruct) DeleteFlightType(id int) (*dom.FlightTypeModel, error) {
	flightType, err := svc.repo.FindFlightTypeByID(int32(id))
	if err != nil {
		return nil, err
	}
	err = svc.repo.DeleteFlightType(id)
	if err != nil {
		log.Printf("Error deleting Flight Type with ID %v, err: %v", id, err.Error())
		return nil, err
	}
	return flightType, nil
}