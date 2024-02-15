package repository

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
)

func (repo *AdminAirlineRepositoryStruct) UpdateBookedSeats(seat dom.BookedSeat, id int) error {
	return repo.DB.Model(&dom.BookedSeat{}).Where("flight_chart_no = ?", seat.FlightChartNo).Updates(seat).Error
}

func (repo *AdminAirlineRepositoryStruct) UpdateEconomyBookedSeat(seat int, seats dom.BookedSeat) error {
	return repo.DB.Model(&dom.BookedSeat{}).Where("flight_chart_no = ?", seats.FlightChartNo).Update("economy_seat_booked", seat).Error
}
func (repo *AdminAirlineRepositoryStruct) UpdateBusinessBookedSeat(seat int, seats dom.BookedSeat) error {
	return repo.DB.Model(&dom.BookedSeat{}).Where("flight_chart_no = ?", seats.FlightChartNo).Update("business_seat_booked", seat).Error
}
