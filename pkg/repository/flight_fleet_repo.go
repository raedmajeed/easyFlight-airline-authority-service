package repository

import (
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
)

func (repo *AdminAirlineRepositoryStruct) FindLastFlightInDB() int {
	var flight dom.FlightFleets
	result := repo.DB.Last(&flight)
	if result.Error != nil {
		log.Printf("err: %v", result.Error.Error())
		return -1
	}
	return int(flight.ID)
}

func (repo *AdminAirlineRepositoryStruct) CreateFlightFleet(fl *dom.FlightFleets) error {
	result := repo.DB.Create(&fl)
	if result.Error != nil {
		log.Printf("err: %v", result.Error.Error())
		return result.Error
	}
	return nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightFleetById(id int) (*dom.FlightFleets, error) {
	var flight dom.FlightFleets
	result := repo.DB.Where("id = ?", id).First(&flight)
	if result.Error != nil {
		log.Printf("err: %v", result.Error.Error())
		return nil, result.Error
	}
	return &flight, nil
}

func (repo *AdminAirlineRepositoryStruct) GetFlightFleets(id uint) ([]dom.FlightFleets, error) {
	var seats []dom.FlightFleets
	if err := repo.DB.Where("airline_id = ?", id).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}
func (repo *AdminAirlineRepositoryStruct) GetFlightFleet(id uint, sid string) (dom.FlightFleets, error) {
	var seat dom.FlightFleets
	if err := repo.DB.Where("airline_id = ? AND id = ?", id, sid).First(&seat).Error; err != nil {
		return dom.FlightFleets{}, err
	}
	return seat, nil
}
func (repo *AdminAirlineRepositoryStruct) DeleteFlightFleet(id uint, sid string) error {
	result := repo.DB.Where("airline_id = ?", id).Delete(&dom.FlightFleets{}, sid)
	return result.Error
}
