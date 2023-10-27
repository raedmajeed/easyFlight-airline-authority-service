package handlers

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
)

type AdminAirlineHandler struct {
	// need service here
	// need jwt utils token generatore here
	svc service.AdminAirlineService
	pb.AdminAirlineServer
}

func NewAdminAirlineHandler(svc service.AdminAirlineService) *AdminAirlineHandler {
	return &AdminAirlineHandler{
		svc: svc,
	}
}

//* METHODS BELOW THIS IS TO HANDLE FLIGHT TYPES

func (handler *AdminAirlineHandler) RegisterFlightType(ctx context.Context, p *pb.FlightTypeRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	response, err := handler.svc.CreateFlightType(p)
	if err != nil {
		log.Printf("Unable to create, err: %v", err.Error())
		return nil, err
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}

func (handler *AdminAirlineHandler) FetchAllFlightTypes(ctx context.Context, p *pb.EmptyRequest) (*pb.FlightTypesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	response, err := handler.svc.GetAllFlightTypes()
	if err != nil {
		log.Println("Unable to fetch flight types some error in service")
		return nil, errors.New("Unable to fetch flight types")
	}
	flightTypesResponse := utils.ConvertFlightModelsToResponse(response)

	return flightTypesResponse, nil
}

func (handler *AdminAirlineHandler) FetchFlightType(ctx context.Context, p *pb.IDRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	response, err := handler.svc.GetFlightType(p.GetId())
	if err != nil {
		log.Println("Unable to fetch flight types some error in service")
		return nil, errors.New("Unable to fetch flight types")
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}

func (handler *AdminAirlineHandler) UpdateFlightType(ctx context.Context, p *pb.FlightTypeRequest) (*pb.FlightTypeResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting")
		return nil, errors.New("Deadline Passed, returning")
	}

	var flight_type_id []string
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		flight_type_id = md.Get("flight_type_id")
	}

	flight_id, err := strconv.Atoi(flight_type_id[0])
	if err != nil {
		log.Printf("No id found or error converting string")
		return nil, errors.New("No id found or error converting string")
	}

	response, err := handler.svc.UpdateFlightType(p, flight_id)
	if err != nil {
		log.Println("Unable to fetch flight types some error in service")
		return nil, errors.New("Unable to fetch flight types")
	}
	flightTypeResponse := utils.ConvertFlightModelToResponse(response)
	return flightTypeResponse, nil
}

func (handler *AdminAirlineHandler) DeleteFlightType(context.Context, *pb.IDRequest) (*pb.FlightTypeResponse, error) {
	return &pb.FlightTypeResponse{}, nil
}
