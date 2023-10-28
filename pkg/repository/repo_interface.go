package repository

// import (
// 	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
// 	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
// 	"gorm.io/gorm"
// )

// type AdminAirlineRepostory interface {
// 	//* Methods to Do repo operation on flight type
// 	CreateFlightType(*pb.FlightTypeRequest) (*dom.FlightTypeModel, error)
// 	FindFlightTypeByModel(model string) (*dom.FlightTypeModel, error)
// 	FindAllFlightTypes() ([]dom.FlightTypeModel, error)
// 	FindFlightTypeByID(id int32) (*dom.FlightTypeModel, error)
// 	UpdateFlightType(*dom.FlightTypeModel) (*dom.FlightTypeModel, error)

// 	//* Methods to Do repo operation on airline type
// 	FindAirlineById(id int32) (*dom.Airline, error)

// 	//* Methods to Do repo operation on airline seats
// 	CreateAirlineSeatType(*pb.AirlineSeatRequest, []byte, []byte) (*dom.AirlineSeat, error)

// 	//* Methods to Do repo operation on airline baggage policy
// 	CreateAirlineBaggagePolicy(*pb.AirlineBaggageRequest, int) (*dom.AirlineBaggage, error)

// 	//* Methods to Do repo operation on airline cancellation policy
// 	CreateAirlineCancellationPolicy(*pb.AirlineCancellationRequest, int) (*dom.AirlineCancellation, error)
// }

// type AdminAirlineRepositoryStruct struct {
// 	DB *gorm.DB
// }

// func NewAdminAirlineRepository(db *gorm.DB) AdminAirlineRepostory {
// 	return &AdminAirlineRepositoryStruct{
// 		DB: db,
// 	}
// }
