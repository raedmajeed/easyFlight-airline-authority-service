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
		log.Printf("no flights available from departure airport: %v for time %v", depAirport, depTime)
		return nil, result.Error
	}
	return flights, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightsFromDep(depAirport string, depDate string) ([]*dom.FlightChart, error) {
	var flights []*dom.FlightChart
	result := repo.DB.Joins("JOIN schedules ON schedules.id = flight_charts.schedule_id").
		Where("schedules.departure_airport = ? and schedules.departure_date = ?", depAirport, depDate).
		Find(&flights)
	//result := repo.DB.Where("departure_airport = ? AND departure_date = ?", depAirport, depDate).Find(&flights)
	if result.Error != nil {
		log.Printf("no flights available from departure airport: %v for date %v", depAirport, depDate)
		return nil, result.Error
	}
	return flights, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightScheduleByAirport(airport string, date time.Time, id int) ([]*dom.FlightChart, error) {
	var flights []*dom.FlightChart
	result := repo.DB.Joins("JOIN schedules ON schedules.id = flight_charts.schedule_id").
		Where("schedules.departure_airport = ? and schedules.departure_date_time >= ?", airport, date).
		Find(&flights)
	if result.Error != nil {
		log.Printf("no flights available from departure airport: %v for date %v", airport, date)
		return nil, result.Error
	}
	return flights, nil
}
