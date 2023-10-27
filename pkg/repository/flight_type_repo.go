package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
	// pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (repo *AdminAirlineRepositoryStruct) FindFlightTypeByModel(model string) (*dom.FlightTypeModel, error) {
	var flightType dom.FlightTypeModel
	result := repo.DB.Where("flight_model = ?", model).First(&flightType)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of flight type %v", model)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &flightType, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightTypeByID(id int32) (*dom.FlightTypeModel, error) {
	var flightType dom.FlightTypeModel
	result := repo.DB.Where("flight_model = ?", id).First(&flightType)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of flight type %v", id)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &flightType, nil
}

func (repo *AdminAirlineRepositoryStruct) CreateFlightType(p *pb.FlightTypeRequest) (*dom.FlightTypeModel, error) {
	flightType := dom.FlightTypeModel{
		Type:                p.GetType().String(),
		FlightModel:         p.FlightModel,
		Description:         p.Description,
		ManufacturerName:    p.ManufacturerName,
		ManufacturerCountry: p.ManufacturerCountry,
		MaxDistance:         p.MaxDistance,
		CruiseSpeed:         p.CruiseSpeed,
	}
	result := repo.DB.Create(&flightType)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", p.FlightModel)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return &flightType, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAllFlightTypes() ([]dom.FlightTypeModel, error) {
	var flightTypes []dom.FlightTypeModel
	result := repo.DB.Find(&flightTypes)
	if result.Error != nil {
		log.Println("Unable to fetch the flight types")
		return nil, result.Error
	}
	return flightTypes, nil
}

func (repo *AdminAirlineRepositoryStruct) UpdateFlightType(d *dom.FlightTypeModel) (*dom.FlightTypeModel, error) {
	result := repo.DB.Save(&d)
	if result.Error != nil {
		log.Println("Unable to Update the flight types")
		return nil, result.Error
	}
	return d, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlineById(id int32) (*dom.Airline, error) {
	var airline dom.Airline
	result := repo.DB.Where("id = ?", int(id)).First(&airline)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of airline %v", id)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &airline, nil
}

type Layout struct{ Rows [][]bool }

func (repo *AdminAirlineRepositoryStruct) CreateAirlineSeatType(p *pb.AirlineSeatRequest, economyLayout []byte, buisLayout []byte) (*dom.AirlineSeat, error) {
	airlineSeat := dom.AirlineSeat{
		AirlineId:           int(p.AirlineId),
		EconomySeatNumber:   int(p.EconomySeatNo),
		BuisinesSeatNumber:  int(p.BuisinesSeatNo),
		EconomySeatsPerRow:  int(p.EconomySeatsPerRow),
		BuisinesSeatsPerRow: int(p.BuisinesSeatsPerRow),
		EconomySeatLayout:   economyLayout,
		BuisinessSeatLayout: buisLayout,
	}
	result := repo.DB.Create(&airlineSeat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", p.AirlineId)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}

	// * this code fetches me the seats, dont delete
	var ad dom.AirlineSeat
	repo.DB.Where("id = ?", 1).First(&ad)
	l := &Layout{}
	_ = json.Unmarshal(ad.EconomySeatLayout, l)
	fmt.Println(l.Rows[0])

	return &airlineSeat, nil
}
