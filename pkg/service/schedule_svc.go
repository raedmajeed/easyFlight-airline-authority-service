package service

import (
	"fmt"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) CreateSchedules(p *pb.ScheduleRequest) (*dom.Schedule, error) {
	//! implement go routines here
	if !checkAirportExists(p.DepartureAirport, svc) && !checkAirportExists(p.ArrivalAirport, svc) {
		log.Printf("depature and airport airport does not exist airport code: %v %v", p.DepartureAirport, p.ArrivalAirport)
		return nil, fmt.Errorf("depature and arrival airport does not exist airport code: %v %v", p.DepartureAirport, p.ArrivalAirport)
	}
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

	fmt.Println(departureTime, arrivalTime) //~ delete after
	if arrivalTime.Before(departureTime) {
		log.Println("Arrival time is behind departure time")
		return nil, fmt.Errorf("Arrival time is behind departure time")
	}

	schedules, err := svc.repo.CreateSchedules(p)
	if err != nil {
		log.Printf("unable to create schedule, err: %v", err.Error())
		return nil, err
	}
	return schedules, nil
}

func checkAirportExists(airport string, svc *AdminAirlineServiceStruct) bool {
	_, err := svc.repo.FindAirportByAirportCode(airport)
	if err != nil {
		return false
	}
	return true
}

func convertScheduleTimeToGo(date, dtime string) (time.Time, error) {
	concTime := date + " " + dtime
	t, err := time.Parse( "02/01/2006 15:04", concTime)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
