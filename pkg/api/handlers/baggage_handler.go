package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE BAGGAGE

func (handler *AdminAirlineHandler) RegisterAirlineBaggage(ctx context.Context, p *pb.AirlineBaggageRequest) (*pb.AirlineBaggageResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateAirlineBaggagePolicy(p)
	if err != nil {
		log.Printf("Unable to create airline baggage policy, err: %v", err.Error())
		return nil, err
	}
	airlineBaggagePolicyResponse := utils.ConvertAirlineBaggagePolicyToResponse(response)
	return airlineBaggagePolicyResponse, nil
}

func (handler *AdminAirlineHandler) FetchAllAirlineBaggages(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggagesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAllAirlineBaggages(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) FetchAirlineBaggage(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggageResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAirlineBaggage(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) DeleteAirlineBaggage(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggageResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	err := handler.svc.DeleteAirlineBaggage(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineBaggageResponse{}, nil
}
