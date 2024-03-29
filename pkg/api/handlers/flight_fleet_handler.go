package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

func (handler *AdminAirlineHandler) RegisterFlightFleets(ctx context.Context, p *pb.FlightFleetRequest) (*pb.FlightFleetResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateFlightFleet(p)
	if err != nil {
		log.Printf("unable to add airline flight fleets, err: %v", err.Error())
		return nil, err
	}
	flightFleetsResponse := utils.ConvertFlightFleetToResponse(response)
	return flightFleetsResponse, nil
}

func (handler *AdminAirlineHandler) GetFlightFleets(ctx context.Context, p *pb.FetchRequest) (*pb.FlightFleetsResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.GetFlightFleets(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) GetFlightFleet(ctx context.Context, p *pb.FetchRequest) (*pb.FlightFleetResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.GetFlightFleet(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) DeleteFlightFleet(ctx context.Context, p *pb.FetchRequest) (*pb.FlightFleetResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	err := handler.svc.DeleteFlightFleet(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.FlightFleetResponse{}, nil
}
