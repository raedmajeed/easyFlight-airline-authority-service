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

func (handler *AdminAirlineHandler) FetchAllAirlineBaggages(context.Context, *pb.EmptyRequest) (*pb.AirlineBaggagesResponse, error) {
	return &pb.AirlineBaggagesResponse{}, nil
}

func (handler *AdminAirlineHandler) FetchAirlineBaggage(context.Context, *pb.IDRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}

func (handler *AdminAirlineHandler) UpdateAirlineBaggage(context.Context, *pb.AirlineBaggageRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}

func (handler *AdminAirlineHandler) DeleteAirlineBaggage(context.Context, *pb.IDRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}
