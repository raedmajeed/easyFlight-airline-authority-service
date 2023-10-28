package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

func (handler *AdminAirlineHandler) RegisterAirportRequest(ctx context.Context, p *pb.Airport) (*pb.AirportResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateAirport(p)
	if err != nil {
		log.Printf("Unable to create airline baggage policy, err: %v", err.Error())
		return nil, err
	}
	airportResponse := utils.ConvertAirportToResponse(response)
	return airportResponse, nil
}
