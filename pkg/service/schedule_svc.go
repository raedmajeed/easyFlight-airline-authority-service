package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) CreateSchedules(p *pb.ScheduleRequest) (*dom.Schedule, error) {
	//! implement go routines here
	if !checkAirportExists(p.DepartureAirport, svc) {
		log.Printf("depature airport does not exist airport code: %v", p.DepartureAirport)
		return nil, fmt.Errorf("depature airport does not exist airport code: %v", p.DepartureAirport)
	}
	if !checkAirportExists(p.ArrivalAirport, svc) {
		log.Printf("arrival airport does not exist airport code: %v", p.ArrivalAirport)
		return nil, fmt.Errorf("arrival airport does not exist airport code: %v", p.ArrivalAirport)
	}

	departureTime, err := convertScheduleTimeToGo(p.DepartureDate, p.DepartureTime)
	if err != nil {
		log.Printf("Error converting departure time: %v", err)
		return nil, err
	}

	arrivalTime, err := convertScheduleTimeToGo(p.ArrivalDate, p.ArrivalTime)
	if err != nil {
		log.Printf("Error converting arrival time: %v", err)
		return nil, err
	}

	if arrivalTime.Before(departureTime) {
		log.Println("Arrival time is behind departure time")
		return nil, errors.New("arrival time is behind departure time")
	}

	scheduled := dom.Schedule{
		DepartureTime:     p.DepartureTime,
		ArrivalTime:       p.ArrivalTime,
		DepartureDate:     p.DepartureDate,
		ArrivalDate:       p.ArrivalDate,
		DepartureAirport:  p.DepartureAirport,
		ArrivalAirport:    p.ArrivalAirport,
		DepartureDateTime: departureTime,
		ArrivalDateTime:   arrivalTime,
	}

	err = svc.repo.CreateSchedules(&scheduled)
	if err != nil {
		log.Printf("unable to create schedule, err: %v", err.Error())
		return nil, err
	}
	return &scheduled, nil
}

func checkAirportExists(airport string, svc *AdminAirlineServiceStruct) bool {
	_, err := svc.repo.FindAirportByAirportCode(airport)
	return err == nil
}

func convertScheduleTimeToGo(date, dtime string) (time.Time, error) {
	format := "02/01/2006 15:04"
	dateTimeStringDep := date + " " + dtime
	tt, _ := time.Parse(format, dateTimeStringDep)
	return tt, nil
}
