package handlers

import (
	"context"
	"errors"
	"log"
	"strconv"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
)

func (handler *AdminAirlineHandler) AdminVerifyAirline(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlineResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	md, check := metadata.FromIncomingContext(ctx)
	if !check {
		log.Println("some error present")
		return nil, errors.New("airline not verified")
	}

	airlineIds := md.Get("airline_id")
	if len(airlineIds) == 0 {
		log.Println("did not receiver airline id via context")
		return nil, errors.New("did not receiver airline via metadata")
	}

	airlineId, err := strconv.Atoi(airlineIds[0])
	if err != nil {
		log.Println("unable to convert airline id in string to int")
		return nil, errors.New("cannot convert airline id from string to int")
	}
	
	response, err := handler.svc.AdminVerifyAirlineRequest(airlineId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("unable to fetch data from db, but otp is verified")
			return nil, err
		} else {
			log.Printf("Airline not added to db, %v", err.Error())
			return nil, err
		}
	}

	return utils.ConvertAirlineToResponse(response), nil
}
