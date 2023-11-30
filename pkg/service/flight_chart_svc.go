package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"math"
	"strconv"
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

	// ! also add fare setting here ===
	economyFare, businessFare, err := calculateAndSavePriceFlightChart(svc, newSchedule.ID, flight.ID)
	if err != nil {
		log.Printf("unable to get the economy and business fare, method AddFlightToChart() - service, err: %v", err.Error())
		data := utils.SendAirlineFareSetFailure(airline.Email, flightNumber, newSchedule.DepartureAirport, newDepartureTime, uint(flightID))
		marshal, _ := json.Marshal(&data)
		_ = svc.kfk.EmailWriter.WriteMessages(context.Background(), kafka.Message{
			Value: marshal,
		})
	}

	flightChart := dom.FlightChart{
		FlightNumber: flightNumber,
		FlightID:     flight.ID,
		ScheduleID:   newSchedule.ID,
		EconomyFare:  math.Round(economyFare),
		BusinessFare: math.Round(businessFare),
	}

	err = svc.repo.CreateFlightChart(&flightChart)
	if err != nil {
		log.Printf("flight chart not created, err: %v", err.Error())
		return nil, err
	}

	flightChartResponse := dom.FlightChartResponse{
		DepartureAirport:  newSchedule.DepartureAirport,
		ArrivalAirport:    newSchedule.ArrivalAirport,
		FlightNumber:      flightNumber,
		DepartureDateTime: newSchedule.DepartureDateTime,
		ArrivalDateTime:   newSchedule.ArrivalDateTime,
		AirlineName:       airline.AirlineName,
		EconomyFare:       flightChart.EconomyFare,
		BusinessFare:      flightChart.BusinessFare,
	}
	return &flightChartResponse, nil
}

func calculateAndSavePriceFlightChart(svc *AdminAirlineServiceStruct, scheduleID uint, flightID uint) (float64, float64, error) {
	response, err := svc.repo.FindScheduleByID(int(scheduleID))
	seats, err := svc.repo.FindSeatsByChartID(flightID)
	if err != nil {
		log.Printf("unable to get schedule ID, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return 1, 1, err
	}
	departureDate := response.DepartureDateTime
	departureAirport := response.DepartureAirport
	arrivalAirport := response.ArrivalAirport
	todayDate := time.Now()
	remainingDays := departureDate.Sub(todayDate)
	days := int(remainingDays.Hours() / 24)
	onlyDate := todayDate.Format("2006-01-02")
	businessSurgeFactor, _ := strconv.ParseFloat(svc.cfg.BUSINESSSURGE, 64)

	depResponse, err := svc.repo.FindAirportByAirportCode(departureAirport)
	if err != nil {
		log.Printf("unable to get departure airport, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return 1, 1, err
	}
	ArrResponse, err := svc.repo.FindAirportByAirportCode(arrivalAirport)
	if err != nil {
		log.Printf("unable to get arrival airport, in method  calculateAndSavePrice() - service, err: %v", err.Error())
		return 1, 1, err
	}
	// take the schedule and find how many days left for departure

	DaysLeftPercentage := CalculateCustomPercentage(days)
	// find schedule here and calculate the distance
	distance := DistanceCalculator(depResponse.Latitude, depResponse.Longitude, ArrResponse.Latitude, ArrResponse.Longitude)
	// once I get the distance fetch today's petrol price
	fuelPrice, err := FuelPricedDaily() //DONE
	// check if today's date has any holiday
	holidayPercentage := dom.Holidays(onlyDate)
	// check % of the days
	weekDayPercentage := dom.DaysOFWeek()
	// check how many seats booked, if % > 50 adjust price accordingly
	eFare, bFare := SeatsBookedPercentage(seats)
	// finally once I get all the values add the price
	EconomyFare := (fuelPrice * distance) / 2
	EconomyFare = EconomyFare + ((EconomyFare * DaysLeftPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * holidayPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * weekDayPercentage) / 100)
	EconomyFare = EconomyFare + ((EconomyFare * eFare) / 100)

	BusinessFare := (fuelPrice * distance) / 2
	BusinessFare = BusinessFare + ((BusinessFare * DaysLeftPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * holidayPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * weekDayPercentage) / 100)
	BusinessFare = BusinessFare + ((BusinessFare * bFare) / 100)
	BusinessFare = BusinessFare * float64(businessSurgeFactor)

	return EconomyFare / 10, BusinessFare / 10, nil
}
