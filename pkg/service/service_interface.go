package service

import (
	"github.com/go-redis/redis/v8"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
)

type AdminAirlineService interface {
	//* Methods to Do service operation on flight type
	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
	UpdateFlightType(*pb.FlightTypeRequest, int) (*dom.FlightTypeModel, error)
	// DeleteFlightType()
	GetFlightType(int32) (*dom.FlightTypeModel, error)
	GetAllFlightTypes() ([]dom.FlightTypeModel, error)

	//* Methods to add airline to db
	RegisterFlight(*pb.AirlineRequest) (*dom.Airline, error)

	//*Methods to add airline seats to db
	CreateAirlineSeats(*pb.AirlineSeatRequest, int) (*dom.AirlineSeat, error)

	//*Methods to add airline baggage policy to db
	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest, int) (*dom.AirlineBaggage, error)

	//*Methods to add airline cancellation policy to db
	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest, int) (*dom.AirlineCancellation, error)
}

type AdminAirlineServiceStruct struct {
	repo  repository.AdminAirlineRepostory
	redis *redis.Client
}

func NewAdminAirlineService(repo repository.AdminAirlineRepostory, redis *redis.Client) AdminAirlineService {
	return &AdminAirlineServiceStruct{
		repo:  repo,
		redis: redis,
	}
}
