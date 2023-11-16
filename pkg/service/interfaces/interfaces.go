package interfaces

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

type AdminAirlineService interface {
	//* Methods to Do service operation on flight type
	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
	UpdateFlightType(*pb.FlightTypeRequest, int) (*dom.FlightTypeModel, error)
	DeleteFlightType(id int) (*dom.FlightTypeModel, error)
	GetFlightType(int32) (*dom.FlightTypeModel, error)
	GetAllFlightTypes() ([]dom.FlightTypeModel, error)

	//* Methods to add airline to db
	RegisterAirlineSvc(*pb.AirlineRequest) (*dom.RegisterAirlineOtpData, error)
	VerifyAirlineRequest(*pb.OTPRequest) (*dom.Airline, error)
	AdminVerifyAirlineRequest(int) (*dom.Airline, error)

	//*Methods to add airline seats to db
	CreateAirlineSeats(*pb.AirlineSeatRequest, int) (*dom.AirlineSeat, error)

	//*Methods to add airline baggage policy to db
	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest, int) (*dom.AirlineBaggage, error)

	//*Methods to add airline cancellation policy to db
	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest, int) (*dom.AirlineCancellation, error)

	//*Methods to add airport to db
	CreateAirport(*pb.Airport) (*dom.Airport, error)

	//*Methods to add schedule to db
	CreateSchedules(*pb.ScheduleRequest) (*dom.Schedule, error)

	//*Methods to do authentication
	AirlineLogin(*pb.LoginRequest) (string, error)
	AdminLogin(*pb.LoginRequest) (string, error)
	AirlineForgotPassword(*pb.ForgotPasswordRequest) (*dom.OtpData, error)
	VerifyOTP(*pb.OTPRequest) (*dom.LoginResponse, error)
	UpdateAirlinePassword(*pb.ConfirmPasswordRequest, string) (string, error)

	//*Methods to do flight fleet
	CreateFlightFleet(*pb.FlightFleetRequest) (*dom.FlightFleetResponse, error)

	//*Methods to do flight chart
	AddFlightToChart(p *pb.FlightChartRequest) (*dom.FlightChart, error)
}
