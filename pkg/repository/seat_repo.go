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
		log.Println("Unable to fetch the flight types")
		return nil, result.Error
	}
	return &seats, nil
}
