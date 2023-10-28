package service

import (
	"encoding/json"
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE SEATS
type Layout struct{ Rows [][]bool }

func (svc *AdminAirlineServiceStruct) CreateAirlineSeats(p *pb.AirlineSeatRequest, id int) (*dom.AirlineSeat, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Flight Type not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	economySeats := p.EconomySeatNo
	economySeatsPerRow := p.EconomySeatsPerRow
	buisinessSeats := p.BuisinesSeatNo
	buisinessSSeatsPerRow := p.BuisinesSeatsPerRow

	ecoLayout := createEconomySeatsJSONLayout(economySeats, economySeatsPerRow)
	buisLayout := createBuisinessSeatsJSONLayout(buisinessSeats, buisinessSSeatsPerRow)

	economyLayoutJSON, err := json.Marshal(ecoLayout)

	if err != nil {
		log.Println("error parsing economy seat layout")
		return nil, err
	}

	buisinessLayoutJSON, err := json.Marshal(buisLayout)
	if err != nil {
		log.Println("error parsing buisiness seat layout")
		return nil, err
	}
	airlineSeats, err := svc.repo.CreateAirlineSeatType(p, economyLayoutJSON, buisinessLayoutJSON)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Seat layout not created of airline %v, err: %v",
				p.AirlineId, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineSeats, nil
}

func createEconomySeatsJSONLayout(seats, sprow int32) *Layout {
	l := &Layout{}

	for i := 1; i <= int(seats); i += int(sprow) {
		row := make([]bool, sprow)
		for j := 1; j <= int(sprow); j++ {
			row = append(row, false)
		}
		l.Rows = append(l.Rows, row)
	}
	return l
}

func createBuisinessSeatsJSONLayout(seats, sprow int32) *Layout {
	l := &Layout{}
	for i := 1; i <= int(seats); i += int(sprow) {
		row := make([]bool, sprow)
		for j := 1; j <= int(sprow); j++ {
			row = append(row, false)
		}
		l.Rows = append(l.Rows, row)
	}
	return l
}
