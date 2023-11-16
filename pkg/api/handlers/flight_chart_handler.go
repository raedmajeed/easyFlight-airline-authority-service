package handlers

import (
	"context"
	"errors"
	"fmt"
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
	response, _ := handler.svc.AddFlightToChart(p)
	fmt.Println(response)
	//if err != nil {
	//	log.Printf("unable to schedule the flight, err: %v", err.Error())
	//	return nil, err
	//}
	//flightChartResponse := utils.ConvertFlightChartToResponse(response)
	//return flightChartResponse, nil
	return &pb.FlightChartResponse{}, nil
}
