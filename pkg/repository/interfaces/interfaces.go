package interfaces

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

type AdminAirlineRepostory interface {
	//* Methods to Do repo operation on flight type
	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
	FindFlightTypeByModel(model string) (*dom.FlightTypeModel, error)
	FindAllFlightTypes() ([]dom.FlightTypeModel, error)
	FindFlightTypeByID(id int32) (*dom.FlightTypeModel, error)
	UpdateFlightType(*dom.FlightTypeModel, int) (*dom.FlightTypeModel, error)
	DeleteFlightType(id int) error

	//* Methods to Do repo operation on airline type
	FindAirlineById(id int32) (*dom.Airline, error)
	CreateAirline(airline *dom.Airline) (*dom.Airline, error)

	//* Methods to Do repo operation on admin
	FindAdminByEmail(p *pb.LoginRequest) (*dom.AdminTable, error)

	//* Methods to Do repo operation on airline seats
	CreateAirlineSeatType(*pb.AirlineSeatRequest, []byte, []byte) (*dom.AirlineSeat, error)
	FindAirlineSeatByid(id int32) (*dom.AirlineSeat, error)

	//* Methods to Do repo operation on airline baggage policy
	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest, int) (*dom.AirlineBaggage, error)
	FindAirlineBaggageByid(id int32) (*dom.AirlineBaggage, error)

	//* Methods to Do repo operation on airline cancellation policy
	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest, int) (*dom.AirlineCancellation, error)
	FindAirlineCancellationByid(id int32) (*dom.AirlineCancellation, error)

	//* Methods to Do repo operation on airport
	FindAirportByAirportCode(string) (*dom.Airport, error)
	CreateAirport(*pb.Airport) (*dom.Airport, error)

	//* Methods to Do repo operation on schedules
	CreateSchedules(schedule *dom.Schedule) error
	FindScheduleByID(id int) (*dom.Schedule, error)

	//* Methods to do repo operation on airline
	FindAirlineByEmail(string) (*dom.Airline, error)
	FindAirlinePassword(*pb.LoginRequest) (*dom.Airline, error)
	InitialAirlinePassword(airline *dom.Airline) (string, error)
	UpdateAirlinePassword(airline *dom.Airline) (string, error)
	UnlockAirlineAccount(int) error

	//* Methods to do repo operation on flight fleet
	FindLastFlightInDB() int
	CreateFlightFleet(fl *dom.FlightFleets) error
	FindFlightFleetById(id int) (*dom.FlightFleets, error)

	//* Methods to do repo operation on flight chart
	FindFlightSeatByID(id int) (*dom.FlightFleets, error)
	FindLastArrivedAirport(flightNumber string) (*dom.FlightChart, error)
	FindFlightScheduleID(id int) (*dom.FlightChart, error)
	CreateFlightChart(flightChart *dom.FlightChart) error
}
