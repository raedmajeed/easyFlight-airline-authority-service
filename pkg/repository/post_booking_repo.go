package repository

import dom "github.com/raedmajeed/admin-servcie/pkg/DOM"

func (repo *AdminAirlineRepositoryStruct) UpdateEconomyBookedSeats(layout []byte, seat *dom.BookedSeat) error {
	return repo.DB.Model(&seat).Update("economy_seat_layout", layout).Error
}

func (repo *AdminAirlineRepositoryStruct) UpdateBusinessBookedSeats(layout []byte, seat *dom.BookedSeat) error {
	return repo.DB.Model(&seat).Update("business_seat_layout", layout).Error
}

func (repo *AdminAirlineRepositoryStruct) UpdateBusinessSeatNo(no int, seat *dom.BookedSeat) error {
	return repo.DB.Model(&seat).Update("business_seat_booked", no).Error
}

func (repo *AdminAirlineRepositoryStruct) UpdateEconomySeatNo(no int, seat *dom.BookedSeat) error {
	return repo.DB.Model(&seat).Update("economy_seat_booked", no).Error
}

func (repo *AdminAirlineRepositoryStruct) FindBookedSeatsByChartID(id uint) (*dom.BookedSeat, error) {
	var data dom.BookedSeat
	result := repo.DB.Where("flight_chart_no = ?", id).First(&data).Error
	return &data, result
}
