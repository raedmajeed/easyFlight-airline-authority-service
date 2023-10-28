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

//* METHODS BELOW THIS IS TO HANDLE AIRLINE BAGGAGE

func (handler *AdminAirlineHandler) RegisterAirlineBaggage(ctx context.Context, p *pb.AirlineBaggageRequest) (*pb.AirlineBaggageResponse, error) {
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

	response, err := handler.svc.CreateAirlineBaggagePolicy(p, id)
	if err != nil {
		log.Printf("Unable to create airline baggage policy, err: %v", err.Error())
		return nil, err
	}
	airlineBaggagePolicyResponse := utils.ConvertAirlineBaggagePolicyToResponse(response)
	return airlineBaggagePolicyResponse, nil
}

func (H *AdminAirlineHandler) FetchAllAirlineBaggages(context.Context, *pb.EmptyRequest) (*pb.AirlineBaggagesResponse, error) {
	return &pb.AirlineBaggagesResponse{}, nil
}

func (H *AdminAirlineHandler) FetchAirlineBaggage(context.Context, *pb.IDRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}

func (H *AdminAirlineHandler) UpdateAirlineBaggage(context.Context, *pb.AirlineBaggageRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}

func (H *AdminAirlineHandler) DeleteAirlineBaggage(context.Context, *pb.IDRequest) (*pb.AirlineBaggageResponse, error) {
	return &pb.AirlineBaggageResponse{}, nil
}
