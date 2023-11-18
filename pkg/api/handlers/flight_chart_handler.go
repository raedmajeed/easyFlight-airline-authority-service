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
