package handlers

import (
	"context"
	"errors"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (handler *AdminAirlineHandler) RegisterFlightChart(ctx context.Context, p *pb.FlightChartRequest) (*pb.FlightChartResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.AddFlightToChart(p)
	if err != nil {
		log.Printf("unable to schedule the flight, err: %v", err.Error())
		return nil, err
	}
	flightChartResponse := utils.ConvertFlightChartToResponse(response)
	return flightChartResponse, nil
}

func (handler *AdminAirlineHandler) GetFlightChart(ctx context.Context, p *pb.GetChartRequest) (*pb.FlightChartResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.GetFlightChart(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (handler *AdminAirlineHandler) GetFlightCharts(ctx context.Context, p *pb.EmptyRequest) (*pb.FlightChartsResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.GetFlightCharts(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) GetFlightChartForAirline(ctx context.Context, p *pb.FetchRequest) (*pb.FlightChartsResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.GetFlightChartForAirline(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
