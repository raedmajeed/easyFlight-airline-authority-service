package interfaces

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"time"
)

type AdminAirlineRepostory interface {
	// CreateFlightType * Methods to Do repo operation on flight type
	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
	FindFlightTypeByModel(model string) (*dom.FlightTypeModel, error)
	FindAllFlightTypes() ([]dom.FlightTypeModel, error)
	FindFlightTypeByID(id int32) (*dom.FlightTypeModel, error)
	UpdateFlightType(*dom.FlightTypeModel, int) (*dom.FlightTypeModel, error)
	DeleteFlightType(id int) error

	// FindAirlineById * Methods to Do repo operation on airline type
	FindAirlineById(id int32) (*dom.Airline, error)
	CreateAirline(airline *dom.Airline) (*dom.Airline, error)

	// FindAdminByEmail * Methods to Do repo operation on admin
	FindAdminByEmail(p *pb.LoginRequest) (*dom.AdminTable, error)

	// CreateAirlineSeatType * Methods to Do repo operation on airline seats
	CreateAirlineSeatType(int, *pb.AirlineSeatRequest, []byte, []byte) (*dom.AirlineSeat, error)
	FindAirlineSeatByid(id int32) (*dom.AirlineSeat, error)

	// CreateAirlineBaggagePolicy * Methods to Do repo operation on airline baggage policy
	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest, int) (*dom.AirlineBaggage, error)
	FindAirlineBaggageByid(id int32) (*dom.AirlineBaggage, error)

	// CreateAirlineCancellationPolicy * Methods to Do repo operation on airline cancellation policy
	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest, int) (*dom.AirlineCancellation, error)
	FindAirlineCancellationByid(id int32) (*dom.AirlineCancellation, error)

	// FindAirportByAirportCode * Methods to Do repo operation on airport
	FindAirportByAirportCode(string) (*dom.Airport, error)
	CreateAirport(*pb.Airport) (*dom.Airport, error)

	// CreateSchedules * Methods to Do repo operation on schedules
	CreateSchedules(schedule *dom.Schedule) error
	FindScheduleByID(id int) (*dom.Schedule, error)
	FindAllSchedules() []*dom.Schedule
	Convert(d, a string, id uint)

	// FindAirlineByEmail * Methods to do repo operation on airline
	FindAirlineByEmail(string) (*dom.Airline, error)
	FindAirlinePassword(*pb.LoginRequest) (*dom.Airline, error)
	InitialAirlinePassword(airline *dom.Airline) (string, error)
	UpdateAirlinePassword(airline *dom.Airline) (string, error)
	UnlockAirlineAccount(int) error

	// FindLastFlightInDB * Methods to do repo operation on flight fleet
	FindLastFlightInDB() int
	CreateFlightFleet(fl *dom.FlightFleets) error
	FindFlightFleetById(id int) (*dom.FlightFleets, error)

	// FindFlightSeatByID * Methods to do repo operation on flight chart
	FindFlightSeatByID(id int) (*dom.FlightFleets, error)
	FindLastArrivedAirport(flightNumber string) (*dom.FlightChart, error)
	FindFlightScheduleID(id int) (*dom.FlightChart, error)
	CreateFlightChart(flightChart *dom.FlightChart) error

	// FindFlightsFromDep FindFlightsFromAirport * Methods to do repo operation on flight search
	FindFlightsFromDep(depAirport string, depDate string) ([]*dom.FlightChart, error)
	FindFlightsFromAirport(depAirport string, depTime time.Time) ([]*dom.FlightChart, error)
	FindFlightScheduleByAirport(airport string, date time.Time, id int) ([]*dom.FlightChart, error)

	FindFlightByFlightNumber(string) (*dom.FlightFleets, error)
	FindSeatsByChartID(uint) (*dom.BookedSeat, error)

	FindBookedSeatsByChartID(uint) (*dom.BookedSeat, error)
	UpdateEconomyBookedSeats([]byte, *dom.BookedSeat) error
	UpdateBusinessBookedSeats([]byte, *dom.BookedSeat) error

	UpdateBusinessSeatNo(int, *dom.BookedSeat) error
	UpdateEconomySeatNo(int, *dom.BookedSeat) error

	FindAllBookedSeats() ([]dom.BookedSeat, error)
	FindFlightChartById(id int) (dom.FlightChart, error)
	UpdateFlightChart(chart dom.FlightChart) error

	UpdateBookedSeats(seat dom.BookedSeat, int2 int) error
	UpdateEconomyBookedSeat(seat int, seats dom.BookedSeat) error
	UpdateBusinessBookedSeat(seat int, seats dom.BookedSeat) error

	FindAllAirlines() ([]dom.Airline, error)
	FindAllAcceptedAirlines() ([]dom.Airline, error)
	FindAllRejectedAirlines() ([]dom.Airline, error)

	FindAllAirports() ([]dom.Airport, error)
	DeleteAirportByCode(string2 string) error

	FindFlightChart(string2 string, string3 string) (dom.FlightChart, error)
	FindAllFlightChart() ([]dom.FlightChart, error)

	FetchAllAirlineSeats(string2 uint) ([]dom.AirlineSeat, error)
	FetchAirlineSeat(string2 uint, string3 string) (dom.AirlineSeat, error)
	DeleteAirlineSeat(string2 uint, string3 string) error

	FetchAllAirlineBaggages(string2 uint) ([]dom.AirlineBaggage, error)
	FetchAirlineBaggage(string2 uint, string3 string) (dom.AirlineBaggage, error)
	DeleteAirlineBaggage(string2 uint, string3 string) error

	FetchAllAirlineCancellations(string2 uint) ([]dom.AirlineCancellation, error)
	FetchAirlineCancellation(string2 uint, string3 string) (dom.AirlineCancellation, error)
	DeleteAirlineCancellation(string2 uint, string3 string) error

	GetFlightFleets(string2 uint) ([]dom.FlightFleets, error)
	GetFlightFleet(string2 uint, string3 string) (dom.FlightFleets, error)
	DeleteFlightFleet(string2 uint, string3 string) error
	GetFlightChartForAirline(string2 uint) ([]dom.FlightChart, error)
}
