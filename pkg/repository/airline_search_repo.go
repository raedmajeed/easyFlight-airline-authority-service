package repository

import (
	"fmt"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
)

func (repo *AdminAirlineRepositoryStruct) FindFlightByFlightNumber(flightNumber string) (*dom.FlightFleets, error) {
	var flight dom.FlightFleets
	result := repo.DB.Where("flight_number = ?", flightNumber).First(&flight)
	if result.Error != nil {
		fmt.Printf("unable to get the flight for %v", flightNumber)
		return nil, result.Error
	}
	return &flight, nil
}

func (repo *AdminAirlineRepositoryStruct) FindSeatsByChartID(id uint) (*dom.BookedSeat, error) {
	var seat dom.BookedSeat
	result := repo.DB.Where("flight_chart_no = ?", id).Find(&seat).Error
	if result != nil {
		return nil, result
	}
	return &seat, nil
}
