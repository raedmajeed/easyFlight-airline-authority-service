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
	result := repo.DB.Joins("JOIN schedules ON flight_charts.schedule_id = schedules.id").
		Where("flight_charts.flight_number = ?", flightNumber).
		Select("flight_charts.id, flight_number, flight_id, status, schedule_id").
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

func (repo *AdminAirlineRepositoryStruct) FindAllBookedSeats() ([]dom.BookedSeat, error) {
	var seats []dom.BookedSeat
	err := repo.DB.Find(&seats).Error
	return seats, err
}

func (repo *AdminAirlineRepositoryStruct) FindFlightChartById(id int) (dom.FlightChart, error) {
	var chart dom.FlightChart
	err := repo.DB.Where("id = ?", id).First(&chart).Error
	return chart, err
}

func (repo *AdminAirlineRepositoryStruct) UpdateFlightChart(chart dom.FlightChart) error {
	err := repo.DB.Model(&dom.FlightChart{}).Where("id = ?", chart.ID).Updates(chart).Error
	return err
}

func (repo *AdminAirlineRepositoryStruct) GetFlightChartForAirline(string2 uint) ([]dom.FlightChart, error) {
	//var seats []dom.FlightFleets
	//if err := repo.DB.Where("airline_id = ?", id).Find(&seats).Error; err != nil {
	//	return nil, err
	//}
	return nil, nil
}

func (repo *AdminAirlineRepositoryStruct) FindFlightChart(dep string, arr string) (dom.FlightChart, error) {
	var chart dom.FlightChart
	if err := repo.DB.Where("departure_airport = ? and arrival_airport = ?", dep, arr).First(&chart).Error; err != nil {
		return dom.FlightChart{}, err
	}
	return chart, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAllFlightChart() ([]dom.FlightChart, error) {
	var charts []dom.FlightChart
	if err := repo.DB.Find(&charts).Error; err != nil {
		return nil, err
	}
	return charts, nil
}
