package repository

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

type Layout struct{ Rows [][]bool }

func (repo *AdminAirlineRepositoryStruct) CreateAirlineSeatType(id int, p *pb.AirlineSeatRequest, economyLayout []byte, buisLayout []byte) (*dom.AirlineSeat, error) {
	airlineSeat := dom.AirlineSeat{
		AirlineId:           id,
		EconomySeatNumber:   int(p.EconomySeatNo),
		BusinessSeatNumber:  int(p.BuisinesSeatNo),
		EconomySeatsPerRow:  int(p.EconomySeatsPerRow),
		BusinessSeatsPerRow: int(p.BuisinesSeatsPerRow),
		EconomySeatLayout:   economyLayout,
		BusinessSeatLayout:  buisLayout,
	}
	result := repo.DB.Create(&airlineSeat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", id)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return &airlineSeat, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlineSeatByid(id int32) (*dom.AirlineSeat, error) {
	var seats dom.AirlineSeat
	result := repo.DB.Where("id = ?", id).First(&seats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &seats, nil
}

func (repo *AdminAirlineRepositoryStruct) FetchAllAirlineSeats(id uint) ([]dom.AirlineSeat, error) {
	var seats []dom.AirlineSeat
	if err := repo.DB.Where("airline_id = ?", id).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}
func (repo *AdminAirlineRepositoryStruct) FetchAirlineSeat(id uint, sid string) (dom.AirlineSeat, error) {
	var seat dom.AirlineSeat
	if err := repo.DB.Where("airline_id = ? AND id = ?", id, sid).First(&seat).Error; err != nil {
		return dom.AirlineSeat{}, err
	}
	return seat, nil
}
func (repo *AdminAirlineRepositoryStruct) DeleteAirlineSeat(id uint, sid string) error {
	result := repo.DB.Where("airline_id = ?", id).Delete(&dom.AirlineSeat{}, sid)
	return result.Error
}
