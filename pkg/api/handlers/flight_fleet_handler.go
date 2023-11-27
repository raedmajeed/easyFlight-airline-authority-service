package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

func (handler *AdminAirlineHandler) RegisterFlightFleets(ctx context.Context, p *pb.FlightFleetRequest) (*pb.FlightFleetResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateFlightFleet(p)
	if err != nil {
		log.Printf("unable to add airline flight fleets, err: %v", err.Error())
		return nil, err
	}
	flightFleetsResponse := utils.ConvertFlightFleetToResponse(response)
	return flightFleetsResponse, nil
}
