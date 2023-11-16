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
