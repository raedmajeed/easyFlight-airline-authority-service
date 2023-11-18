package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) AddFlightToChart(p *pb.FlightChartRequest) (*dom.FlightChartResponse, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		log.Printf("unable to find airline: %v", p.AirlineEmail)
		return nil, fmt.Errorf("unable to find airline: %v, err: %v", p.AirlineEmail, err.Error())
	}

	flightID := p.FlightFleetId
	flight, err := svc.repo.FindFlightFleetById(int(flightID))
	if err != nil {
		log.Println("some error")
	}

	// if the particular flight is not owned by the airline it throws an error
	if airline.ID != flight.AirlineID {
		log.Printf("flight of %v is not available in your fleet", flight.FlightNumber)
		return nil, fmt.Errorf("flight of %v is not available in your fleet", flight.FlightNumber)
	}
	newSchedule, err := svc.repo.FindScheduleByID(int(p.ScheduleId))
	if err != nil {
		return nil, err
	}

	newDepartureAirport := newSchedule.DepartureAirport
	newDepartureTime := newSchedule.DepartureDateTime
	if err != nil {
		return nil, err
	}

	flightNumber := flight.FlightNumber
	lastChartOfFlight, _ := svc.repo.FindLastArrivedAirport(flightNumber)
	// if the flight chart dosen't contain any flight of that number id adds directly
	if lastChartOfFlight == nil {
		flightChart := dom.FlightChart{
			FlightNumber: flightNumber,
			FlightID:     flight.ID,
			ScheduleID:   newSchedule.ID,
		}
		err = svc.repo.CreateFlightChart(&flightChart)
		flighChartResponse := dom.FlightChartResponse{
			DepartureAirport:  newSchedule.DepartureAirport,
			ArrivalAirport:    newSchedule.ArrivalAirport,
			FlightNumber:      flightNumber,
			DepartureDateTime: newSchedule.DepartureDateTime,
			ArrivalDateTime:   newSchedule.ArrivalDateTime,
			AirlineName:       airline.AirlineName,
		}
		return &flighChartResponse, nil
	}

	oldFlightSchedule, err := svc.repo.FindScheduleByID(int(lastChartOfFlight.ScheduleID))
	oldArrivedAirport := oldFlightSchedule.ArrivalAirport
	oldArrivalTime := oldFlightSchedule.ArrivalDateTime

	//* if new departure airport == old arrived airport and new departure time > old approved time good to go
	if oldArrivedAirport != newDepartureAirport {
		log.Println("the flight is at a different airport, schedule using available flights at the departure airport")
		return nil, errors.New("the flight is at a different airport, schedule using available flights at the departure airport")
	}

	if !oldArrivalTime.Add(time.Hour).Before(newDepartureTime) {
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

	// ! also add fare setting here ===

	flightChartResponse := dom.FlightChartResponse{
		DepartureAirport:  newSchedule.DepartureAirport,
		ArrivalAirport:    newSchedule.ArrivalAirport,
		FlightNumber:      flightNumber,
		DepartureDateTime: newSchedule.DepartureDateTime,
		ArrivalDateTime:   newSchedule.ArrivalDateTime,
		AirlineName:       airline.AirlineName,
	}
	return &flightChartResponse, nil
}
