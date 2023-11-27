package service

import (
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) CreateFlightFleet(p *pb.FlightFleetRequest) (*dom.FlightFleetResponse, error) {
	// test this out using start time and time.since
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		log.Printf("unable to find airline: %v", p.AirlineEmail)
		return nil, fmt.Errorf("unable to find airline: %v, err: %v", p.AirlineEmail, err.Error())
	}

	// change this all methods to group error go routine
	flightType, err := svc.repo.FindFlightTypeByID(p.FlightTypeId)
	if err != nil {
		return nil, fmt.Errorf("flight type not available, err: %v", err.Error())
	}
	airlineSeat, err := svc.repo.FindAirlineSeatByid(p.SeatId)
	if err != nil {
		return nil, fmt.Errorf("airline seat not available, err: %v", err.Error())
	}
	airlineBaggage, err := svc.repo.FindAirlineBaggageByid(p.BaggagePolicyId)
	if err != nil {
		return nil, fmt.Errorf("airline baggage type not available, err: %v", err.Error())
	}
	airlineCancellation, err := svc.repo.FindAirlineCancellationByid(p.CancellationPolicyId)
	if err != nil {
		return nil, fmt.Errorf("airline cancellation type not available, err: %v", err.Error())
	}

	//~ this generates the flight number
	flightNumber := generateFlightNumber(airline.AirlineCode, svc)

	flightFleet := &dom.FlightFleets{
		FlightNumber:         flightNumber,
		AirlineID:            airline.ID,
		SeatID:               uint(airlineSeat.ID),
		FlightTypeID:         uint(flightType.ID),
		BaggagePolicyID:      uint(airlineBaggage.ID),
		CancellationPolicyID: uint(airlineCancellation.ID),
	}

	err = svc.repo.CreateFlightFleet(flightFleet)
	if err != nil {
		log.Printf("unable to create flight fleet, err: %v", err.Error())
		return nil, err
	}

	flightResponse := dom.FlightFleetResponse{
		FlightNumber:       flightNumber,
		FlightTypeModel:    flightType.FlightModel,
		AirlineName:        airline.AirlineName,
		EconomySeatNumber:  airlineSeat.EconomySeatNumber,
		BusinessSeatNumber: airlineSeat.BusinessSeatNumber,
	}

	return &flightResponse, nil
}

func generateFlightNumber(airlineCode string, svc *AdminAirlineServiceStruct) string {
	uniqueNo := findUniqueNo(svc)
	return fmt.Sprintf("%s-%03d", airlineCode, uniqueNo)
}

func findUniqueNo(svc *AdminAirlineServiceStruct) int {
	flightNo := svc.repo.FindLastFlightInDB()
	if flightNo == -1 {
		return 1
	} else {
		return flightNo + 1
	}
}
