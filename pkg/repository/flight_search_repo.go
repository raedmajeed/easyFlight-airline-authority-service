package repository

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"log"
	"time"
)

func (repo *AdminAirlineRepositoryStruct) FindFlightsFromAirport(depAirport string, depTime time.Time) ([]*dom.FlightChart, error) {
	var flights []*dom.FlightChart
	result := repo.DB.Where("departure_airport = ? AND departure_date_time >= ?", depAirport, depTime).Find(&flights)
	if result.Error != nil {
		log.Println("no flights available")
		return nil, result.Error
	}
	return flights, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightsFromDep(depAirport string, depDate string) ([]*dom.FlightChart, error) {
	var flights []*dom.FlightChart
	result := repo.DB.Joins("Schedule").Where("Schedule.departure_airport = ? and Schedule.departure_date = ?", depAirport, depDate).Find(&flights)
	//result := repo.DB.Where("departure_airport = ? AND departure_date = ?", depAirport, depDate).Find(&flights)
	if result.Error != nil {
		log.Println("no flights available")
		return nil, result.Error
	}
	return flights, nil
}
