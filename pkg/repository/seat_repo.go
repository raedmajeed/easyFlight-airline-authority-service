package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

type Layout struct{ Rows [][]bool }

func (repo *AdminAirlineRepositoryStruct) CreateAirlineSeatType(p *pb.AirlineSeatRequest, economyLayout []byte, buisLayout []byte) (*dom.AirlineSeat, error) {
	airlineSeat := dom.AirlineSeat{
		AirlineId:           int(p.AirlineId),
		EconomySeatNumber:   int(p.EconomySeatNo),
		BuisinesSeatNumber:  int(p.BuisinesSeatNo),
		EconomySeatsPerRow:  int(p.EconomySeatsPerRow),
		BuisinesSeatsPerRow: int(p.BuisinesSeatsPerRow),
		EconomySeatLayout:   economyLayout,
		BuisinessSeatLayout: buisLayout,
	}
	result := repo.DB.Create(&airlineSeat)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", p.AirlineId)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}

	// * this code fetches me the seats, dont delete
	var ad dom.AirlineSeat
	repo.DB.Where("id = ?", 1).First(&ad)
	l := &Layout{}
	_ = json.Unmarshal(ad.EconomySeatLayout, l)
	fmt.Println(l.Rows[0])

	return &airlineSeat, nil
}
