package repository

import (
	"errors"
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
	result := repo.DB.Where("id = ?", id).First(&flightType)
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

func (repo *AdminAirlineRepositoryStruct) UpdateFlightType(flightType *dom.FlightTypeModel, id int) (*dom.FlightTypeModel, error) {
	result := repo.DB.Model(&dom.FlightTypeModel{}).Where("id = ?", id).Updates(flightType)
	if result.Error != nil {
		return nil, result.Error
	}
	return flightType, nil
}

func (repo *AdminAirlineRepositoryStruct) DeleteFlightType(id int) error {
	result := repo.DB.Delete(&dom.FlightTypeModel{}, id)
	return result.Error
}