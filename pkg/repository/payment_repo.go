package repository

import dom "github.com/raedmajeed/admin-servcie/pkg/DOM"

func (repo *AdminAirlineRepositoryStruct) UpdateBookedSeats(seat dom.BookedSeat, id int) error {
	return repo.DB.Model(&dom.BookedSeat{}).Where("id = ?", id).Updates(seat).Error
}
