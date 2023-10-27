package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	k "github.com/raedmajeed/admin-servcie/kafkas"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
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
	fmt.Println(airline_id)
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
