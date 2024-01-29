package handlers

import (
	"context"
	"errors"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE SEATS

func (handler *AdminAirlineHandler) RegisterAirlineSeat(ctx context.Context, p *pb.AirlineSeatRequest) (*pb.AirlineSeatResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateAirlineSeats(p)
	if err != nil {
		return nil, err
	}
	return utils.ConvertAirlineSeatsToResponse(response), nil
}

func (handler *AdminAirlineHandler) FetchAllAirlineSeats(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatsResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAllAirlineSeats(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *AdminAirlineHandler) FetchAirlineSeat(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAirlineSeat(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *AdminAirlineHandler) DeleteAirlineSeat(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineSeatResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	err := handler.svc.DeleteAirlineSeat(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineSeatResponse{}, nil
}
