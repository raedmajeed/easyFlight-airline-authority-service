package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE SEATS

func (svc *AdminAirlineServiceStruct) CreateAirlineSeats(p *pb.AirlineSeatRequest) (*dom.AirlineSeat, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineEmail)
			return nil, err
		} else {
			log.Printf("Flight Type not create of model %v, err: %v", p.AirlineEmail, err.Error())
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
	airlineSeats, err := svc.repo.CreateAirlineSeatType(int(airline.ID), p, economyLayoutJSON, buisinessLayoutJSON)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Seat layout not created of airline %v, err: %v",
				p.AirlineEmail, err.Error())
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

func (svc *AdminAirlineServiceStruct) FetchAllAirlineSeats(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatsResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAllAirlineSeats(airline.ID)
	if err != nil {
		return nil, err
	}
	result := ConvertToResponse(resp)
	return result, err
}
func (svc *AdminAirlineServiceStruct) FetchAirlineSeat(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAirlineSeat(airline.ID, p.Id)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineSeatResponse{
		AirlineSeat: &pb.AirlineSeatRequest{
			EconomySeatNo:       int32(resp.BusinessSeatNumber),
			BuisinesSeatNo:      int32(resp.EconomySeatNumber),
			EconomySeatsPerRow:  int32(resp.EconomySeatsPerRow),
			BuisinesSeatsPerRow: int32(resp.BusinessSeatsPerRow),
		},
	}, err
}
func (svc *AdminAirlineServiceStruct) DeleteAirlineSeat(ctx context.Context, p *pb.FetchRequest) error {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	err := svc.repo.DeleteAirlineSeat(airline.ID, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToResponse(data []dom.AirlineSeat) *pb.AirlineSeatsResponse {
	var result []*pb.AirlineSeatRequest
	for _, d := range data {
		result = append(result, &pb.AirlineSeatRequest{
			EconomySeatNo:       int32(d.EconomySeatNumber),
			BuisinesSeatNo:      int32(d.BusinessSeatNumber),
			EconomySeatsPerRow:  int32(d.EconomySeatsPerRow),
			BuisinesSeatsPerRow: int32(d.BusinessSeatsPerRow),
		})
	}
	return &pb.AirlineSeatsResponse{
		AirlineSeats: result,
	}
}
