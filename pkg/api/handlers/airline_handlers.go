package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	k "github.com/raedmajeed/admin-servcie/kafkas"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE COMPANY

func (handler *AdminAirlineHandler) RegisterAirline(ctx context.Context, p *pb.AirlineRequest) (*pb.AirlineResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.RegisterFlight(p)
	if err != nil {
		log.Printf("Unable to create, err: %v", err.Error())
		return nil, err
	}

	k.SetupKafka()
	return utils.ConvertAirlineToResponse(response), nil
}

func (H *AdminAirlineHandler) FetchAllAirlines(context.Context, *pb.EmptyRequest) (*pb.AirlinesResponse, error) {
	return &pb.AirlinesResponse{}, nil
}

func (H *AdminAirlineHandler) FetchAirline(context.Context, *pb.IDRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}

func (H *AdminAirlineHandler) UpdateAirline(context.Context, *pb.AirlineRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}

func (H *AdminAirlineHandler) DeleteAirline(context.Context, *pb.IDRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}
