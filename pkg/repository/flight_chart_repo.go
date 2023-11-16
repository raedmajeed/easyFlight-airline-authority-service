package repository

import (
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
)

func (repo *AdminAirlineRepositoryStruct) FindFlightSeatByID(id int) (*dom.FlightFleets, error) {
	var flight dom.FlightFleets
	result := repo.DB.Preload("Seat").Where("id = ?", id).Find(&flight)
	if result.Error != nil {
		log.Println("unable to fetch seat data from db")
		return nil, result.Error
	}

	return &flight, nil
}

func (repo *AdminAirlineRepositoryStruct) FindLastArrivedAirport(flightNumber string) (*dom.FlightChart, error) {
	var flightChart dom.FlightChart
	result := repo.DB.Preload("Schedule").
		Where("flight_charts.flight_number = ?", flightNumber).
		Order("schedules.arrival_date_time DESC").
		First(&flightChart)
	if result.Error != nil {
		log.Println("unable to fetch data from db")
		return nil, result.Error
	}

	return &flightChart, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightScheduleID(id int) (*dom.FlightChart, error) {
	var flight dom.FlightChart
	result := repo.DB.Preload("Schedule").Find(&flight).Where("id = ?", id)
	if result.Error != nil {
		log.Println("unable to fetch data from db")
		return nil, result.Error
	}

	return &flight, nil
}

func (repo *AdminAirlineRepositoryStruct) CreateFlightChart(flightChart *dom.FlightChart) error {
	result := repo.DB.Create(&flightChart)
	if result.Error != nil {
		log.Printf("unable to create flight schedule err: %v", result.Error.Error())
		return result.Error
	}
	return nil
}
