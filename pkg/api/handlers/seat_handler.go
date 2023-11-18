package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE SEATS

func (handler *AdminAirlineHandler) RegisterAirlineSeat(ctx context.Context, p *pb.AirlineSeatRequest) (*pb.AirlineSeatResponse, error) {
	log.Println("reached registering seats function at admin service")
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateAirlineSeats(p)
	if err != nil {
		log.Printf("Unable to create airline seats, err: %v", err.Error())
		return nil, err
	}
	return utils.ConvertAirlineSeatsToResponse(response), nil
}

func (handler *AdminAirlineHandler) FetchAllAirlineSeats(context.Context, *pb.EmptyRequest) (*pb.AirlineSeatsResponse, error) {
	return &pb.AirlineSeatsResponse{}, nil
}

func (handler *AdminAirlineHandler) FetchAirlineSeat(context.Context, *pb.IDRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}

func (handler *AdminAirlineHandler) UpdateAirlineSeat(context.Context, *pb.AirlineSeatRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}

func (handler *AdminAirlineHandler) DeleteAirlineSeat(context.Context, *pb.IDRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}
