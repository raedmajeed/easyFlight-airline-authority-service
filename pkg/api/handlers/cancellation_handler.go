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

func (handler *AdminAirlineHandler) FetchAllAirlineCancellations(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationsResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAllAirlineCancellations(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *AdminAirlineHandler) FetchAirlineCancellation(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAirlineCancellation(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *AdminAirlineHandler) DeleteAirlineCancellation(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	err := handler.svc.DeleteAirlineCancellation(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineCancellationResponse{}, nil
}
