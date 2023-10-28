package handlers

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE SEATS

func (handler *AdminAirlineHandler) RegisterAirlineSeat(ctx context.Context, p *pb.AirlineSeatRequest) (*pb.AirlineSeatResponse, error) {
	log.Println("reached registering seats function at admin service")
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("not able to get the metadata from url")
		return nil, errors.New("Not able to get the metada from context")
	}

	airline_id := md.Get("airline_id")[0]
	id, err := strconv.Atoi(airline_id)
	if err != nil {
		log.Println("error converting into int")
		return nil, err
	}

	response, err := handler.svc.CreateAirlineSeats(p, id)
	if err != nil {
		log.Printf("Unable to create airline seats, err: %v", err.Error())
		return nil, err
	}
	return utils.ConvertAirlineSeatsToResponse(response), nil
}

func (H *AdminAirlineHandler) FetchAllAirlineSeats(context.Context, *pb.EmptyRequest) (*pb.AirlineSeatsResponse, error) {
	return &pb.AirlineSeatsResponse{}, nil
}

func (H *AdminAirlineHandler) FetchAirlineSeat(context.Context, *pb.IDRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}

func (H *AdminAirlineHandler) UpdateAirlineSeat(context.Context, *pb.AirlineSeatRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}

func (H *AdminAirlineHandler) DeleteAirlineSeat(context.Context, *pb.IDRequest) (*pb.AirlineSeatResponse, error) {
	return &pb.AirlineSeatResponse{}, nil
}
