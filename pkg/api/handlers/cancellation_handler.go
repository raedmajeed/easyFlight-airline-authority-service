package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE CANCELLATION

func (handler *AdminAirlineHandler) RegisterAirlineCancellation(ctx context.Context, p *pb.AirlineCancellationRequest) (*pb.AirlineCancellationResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateAirlineCancellationPolicy(p)
	if err != nil {
		log.Printf("Unable to create airline baggage policy, err: %v", err.Error())
		return nil, err
	}
	airlineCancellationPolicyResponse := utils.ConvertAirlineCancellationPolicyToResponse(response)
	return airlineCancellationPolicyResponse, nil
}

func (handler *AdminAirlineHandler) FetchAllAirlineCancellations(context.Context, *pb.EmptyRequest) (*pb.AirlineCancellationsResponse, error) {
	return &pb.AirlineCancellationsResponse{}, nil
}

func (handler *AdminAirlineHandler) FetchAirlineCancellation(context.Context, *pb.IDRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}

func (handler *AdminAirlineHandler) UpdateAirlineCancellation(context.Context, *pb.AirlineCancellationRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}

func (handler *AdminAirlineHandler) DeleteAirlineCancellation(context.Context, *pb.IDRequest) (*pb.AirlineCancellationResponse, error) {
	return &pb.AirlineCancellationResponse{}, nil
}
