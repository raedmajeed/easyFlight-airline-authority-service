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

//* METHODS BELOW THIS IS TO HANDLE AIRLINE CANCELLATION

func (handler *AdminAirlineHandler) RegisterAirlineCancellation(ctx context.Context, p *pb.AirlineCancellationRequest) (*pb.AirlineCancellationResponse, error) {
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

	airline_id := md.Get("airline_id")
	id, err := strconv.Atoi(airline_id[0])
	if err != nil {
		log.Println("error converting into int")
		return nil, err
	}

	response, err := handler.svc.CreateAirlineCancellationPolicy(p, id)
	if err != nil {
		log.Printf("Unable to create airline baggage policy, err: %v", err.Error())
		return nil, err
	}
	airlineCancellationPolicyResponse := utils.ConvertAirlineCancellationPolicyToResponse(response)
	return airlineCancellationPolicyResponse, nil
}

func (H *AdminAirlineHandler) FetchAllAirlineCancellations(context.Context, *pb.EmptyRequest) (*pb.AirlineCancellationsResponse, error) {
	return &pb.AirlineCancellationsResponse{}, nil
}

func (H *AdminAirlineHandler) FetchAirlineCancellation(context.Context, *pb.IDRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}

func (H *AdminAirlineHandler) UpdateAirlineCancellation(context.Context, *pb.AirlineCancellationRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}

func (H *AdminAirlineHandler) DeleteAirlineCancellation(context.Context, *pb.IDRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}
