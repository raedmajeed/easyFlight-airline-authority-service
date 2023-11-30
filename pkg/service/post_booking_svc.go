package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"log"
	"strconv"
)

type Layout struct{ Rows [][]bool }

func (svc *AdminAirlineServiceStruct) SelectAndBookSeats(ctx context.Context, request *pb.SeatRequest) (*pb.SeatResponse, error) {
	pnrNumber := request.PNR
	seatArray := request.SeatArray
	flightChart := request.FlightChartId

	response, err := svc.repo.FindBookedSeatsByChartID(uint(flightChart))
	if err != nil {
		log.Printf("error finding booked seats, err: %v", err.Error())
		return nil, err
	}

	economy := true
	seatLayout := response.EconomySeatLayout
	if !request.Economy {
		economy = false
		seatLayout = response.BusinessSeatLayout
	}

	var layout Layout
	err = json.Unmarshal(seatLayout, &layout)
	if err != nil {
		log.Printf("unable to unmarshal json, err: %v", err.Error())
		return nil, err
	}

	seatsInRows := len(layout.Rows[0])
	columns := len(layout.Rows)

	var seatNos []string
	for _, seat := range seatArray {
		s := fmt.Sprintf("%v", seat)
		rowS := string(s[0])
		columnS := string(s[1])
		row, _ := strconv.Atoi(rowS)
		column, _ := strconv.Atoi(columnS)
		r := string(rune(row + 65))
		err = checkRowColumn(row, seatsInRows, column, columns)
		if err != nil {
			return nil, err
		}
		err = checkSeatBooked(layout.Rows[row][column], r+columnS)
		if err != nil {
			return nil, err
		}
		layout.Rows[row][column] = true
		seatNos = append(seatNos, r+columnS)
	}

	marshal, err := json.Marshal(layout)
	if economy {
		response.EconomySeatLayout = marshal
		err = svc.repo.UpdateEconomyBookedSeats(response.EconomySeatLayout, response)
		ecoSeat := response.EconomySeatBooked + len(seatArray)
		err = svc.repo.UpdateEconomySeatNo(ecoSeat, response)
	} else {
		response.BusinessSeatLayout = marshal
		err = svc.repo.UpdateBusinessBookedSeats(response.BusinessSeatLayout, response)
		busSeat := response.BusinessSeatBooked + len(seatArray)
		err = svc.repo.UpdateBusinessSeatNo(busSeat, response)
	}
	if err != nil {
		log.Printf("unable to update booked seats, err: %v", err.Error())
		return nil, err
	}

	return &pb.SeatResponse{
		PNR:     pnrNumber,
		SeatNos: seatNos,
	}, nil

}

func checkSeatBooked(l bool, s string) error {
	if l == true {
		return fmt.Errorf("seat %v already booked, please try with another seat", s)
	}
	return nil
}

func checkRowColumn(r, rs, c, cs int) error {
	if r > rs {
		return errors.New("row is greater than rows in flight")
	}
	if c > cs {
		return errors.New("column is greater than columns in flight")
	}
	return nil
}
