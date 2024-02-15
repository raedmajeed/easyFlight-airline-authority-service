package interfaces

import (
	"context"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

type AdminAirlineService interface {
	// CreateFlightType * Methods to Do service operation on flight type
	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
	UpdateFlightType(*pb.FlightTypeRequest, int) (*dom.FlightTypeModel, error)
	DeleteFlightType(id int) (*dom.FlightTypeModel, error)
	GetFlightType(int32) (*dom.FlightTypeModel, error)
	GetAllFlightTypes() ([]dom.FlightTypeModel, error)

	// RegisterAirlineSvc * Methods to add airline to db
	RegisterAirlineSvc(*pb.AirlineRequest) (*dom.RegisterAirlineOtpData, error)
	VerifyAirlineRequest(*pb.OTPRequest) (*dom.Airline, error)
	AdminVerifyAirlineRequest(int) (*dom.Airline, error)

	// CreateAirlineSeats *Methods to add airline seats to db
	CreateAirlineSeats(*pb.AirlineSeatRequest) (*dom.AirlineSeat, error)

	// CreateAirlineBaggagePolicy *Methods to add airline baggage policy to db
	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest) (*dom.AirlineBaggage, error)

	// CreateAirlineCancellationPolicy *Methods to add airline cancellation policy to db
	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest) (*dom.AirlineCancellation, error)

	// CreateAirport *Methods to add airport to db
	CreateAirport(*pb.Airport) (*dom.Airport, error)

	// CreateSchedules *Methods to add schedule to db
	CreateSchedules(*pb.ScheduleRequest) (*dom.Schedule, error)

	// AirlineLogin *Methods to do authentication
	AirlineLogin(*pb.LoginRequest) (string, error)
	AdminLogin(*pb.LoginRequest) (string, error)
	AirlineForgotPassword(*pb.ForgotPasswordRequest) (*dom.OtpData, error)
	VerifyOTP(*pb.OTPRequest) (*dom.LoginResponse, error)
	UpdateAirlinePassword(*pb.ConfirmPasswordRequest, string) (string, error)

	// CreateFlightFleet *Methods to do flight fleet
	CreateFlightFleet(*pb.FlightFleetRequest) (*dom.FlightFleetResponse, error)

	// AddFlightToChart *Methods to do flight chart
	AddFlightToChart(p *pb.FlightChartRequest) (*dom.FlightChartResponse, error)

	SearchFlightInitial(ctx context.Context, p *pb.SearchFlightRequestAdmin) (*pb.SearchFlightResponseAdmin, error)
	SearchFlight(search dom.SearchDetails) ([]dom.Path, []dom.Path, error)
	SearchSelectFlight(context.Context, *pb.SelectFlightAdmin) (*pb.CompleteFlightDetails, error)
	//SearchSelectFlight(ctx context.Context, message kafka.Message)

	SelectAndBookSeats(ctx context.Context, request *pb.SeatRequest) (*pb.SeatResponse, error)
	CalculateDailyFare()
	AddConfirmedSeatsToBooked(context.Context, *pb.ConfirmedSeatRequest) (*pb.ConfirmedSeatResponse, error)

	// to be completed from here
	FetchAllAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error)
	AcceptedAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error)
	RejectedAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error)

	GetAirport(ctx context.Context, p *pb.AirportRequest) (*pb.AirportResponse, error)
	GetAirports(ctx context.Context, p *pb.EmptyRequest) (*pb.AirportsResponse, error)
	DeleteAirport(ctx context.Context, p *pb.AirportRequest) error

	GetSchedules(ctx context.Context, p *pb.EmptyRequest) (*pb.SchedulesResponse, error)

	GetFlightChart(ctx context.Context, p *pb.GetChartRequest) (*pb.FlightChartResponse, error)
	GetFlightCharts(ctx context.Context, p *pb.EmptyRequest) (*pb.FlightChartsResponse, error)

	FetchAllAirlineSeats(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatsResponse, error)
	FetchAirlineSeat(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatResponse, error)
	DeleteAirlineSeat(ctx context.Context, p *pb.FetchRequest) error

	FetchAllAirlineBaggages(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggagesResponse, error)
	FetchAirlineBaggage(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggageResponse, error)
	DeleteAirlineBaggage(ctx context.Context, p *pb.FetchRequest) error

	FetchAllAirlineCancellations(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationsResponse, error)
	FetchAirlineCancellation(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationResponse, error)
	DeleteAirlineCancellation(ctx context.Context, p *pb.FetchRequest) error

	GetFlightFleets(ctx context.Context, p *pb.FetchRequest) (*pb.FlightFleetsResponse, error)
	GetFlightFleet(ctx context.Context, p *pb.FetchRequest) (*pb.FlightFleetResponse, error)
	DeleteFlightFleet(ctx context.Context, p *pb.FetchRequest) error

	GetFlightChartForAirline(ctx context.Context, p *pb.FetchRequest) (*pb.FlightChartsResponse, error)
}
