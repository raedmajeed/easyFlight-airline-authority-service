package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) AddFlightToChart(p *pb.FlightChartRequest) (*dom.FlightChart, error) {
	_, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		log.Printf("unable to find airline: %v", p.AirlineEmail)
		return nil, fmt.Errorf("unable to find airline: %v, err: %v", p.AirlineEmail, err.Error())
	}

	flightID := p.FlightFleetId

	fmt.Println("id =======", p.ScheduleId)
	flight, err := svc.repo.FindFlightFleetById(int(flightID))
	if err != nil {
		log.Println("some error")
	}

	fmt.Println(flight, "flight ==============")

	newSchedule, err := svc.repo.FindScheduleByID(int(p.ScheduleId))
	if err != nil {
		return nil, err
	}

	fmt.Println(newSchedule.ID, "id 00000000")

	newDepartureAirport := newSchedule.DepartureAirport

	newDepartureTime := newSchedule.DepartureDateTime

	if err != nil {
		return nil, err
	}

	flightNumber := flight.FlightNumber

	lastChartOfFlight, _ := svc.repo.FindLastArrivedAirport(flightNumber)

	if lastChartOfFlight == nil {
		flightChart := dom.FlightChart{
			FlightNumber: flightNumber,
			FlightID:     flight.ID,
			ScheduleID:   newSchedule.ID,
		}
		err = svc.repo.CreateFlightChart(&flightChart)
		return &flightChart, nil
	}
	//? this preloads the schedule from last flight
	//oldFlight, _ := svc.repo.FindFlightScheduleID(lastChartOfFlight.ID)
	oldFlightSchedule := lastChartOfFlight.Schedule
	oldArrivedAirport := oldFlightSchedule.ArrivalAirport

	//? find the schedule from schedule id

	oldDepartureTime := oldFlightSchedule.ArrivalDateTime

	//* if new departure airport == old arrived airport and new departure time > old arroved time good to go
	if oldArrivedAirport != newDepartureAirport {
		log.Println("the flight is at a different airport, schedule using available flights at the departure airport")
		return nil, errors.New("the flight is at a different airport, schedule using available flights at the departure airport")
	}

	if !oldDepartureTime.Add(time.Hour).Before(newDepartureTime) {
		log.Println("layover time is less than an hour, not possible to schedule flight")
		return nil, errors.New("layover time is less than an hour, not possible to schedule flight")
	}

	if !flight.IsInService {
		log.Println("flight is not in service, use another flight")
		return nil, errors.New("flight is not in service, use another flight")
	}

	if flight.Maintenance {
		log.Println("flight is in maintenance, use another flight")
		return nil, errors.New("flight is in maintenance, use another flight")
	}

	//* creating flight chart here
	flightChart := dom.FlightChart{
		FlightNumber: flightNumber,
		FlightID:     flight.ID,
		ScheduleID:   newSchedule.ID,
	}
	err = svc.repo.CreateFlightChart(&flightChart)
	if err != nil {
		log.Printf("flight chart not created, err: %v", err.Error())
		return nil, err
	}
	// ! also add fare setting here
	// flightSeatData, err := svc.repo.FindFlightSeatByID(int(p.FlightFleetId))
	// if err != nil {
	// 	return nil, err
	// }

	// flightSeat := flightSeatData.Seat

	// if oldFlightSchedule.ArrivalAirport ==

	//! create booked seat table here
	return &flightChart, nil
}
