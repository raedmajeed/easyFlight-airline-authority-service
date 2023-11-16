package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

//* METHODS BELOW THIS IS TO HANDLE AIRLINE COMPANY

func (handler *AdminAirlineHandler) RegisterAirline(ctx context.Context, p *pb.AirlineRequest) (*pb.OtpResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.RegisterAirlineSvc(p)
	if err != nil {
		log.Printf("Unable to create, err: %v", err.Error())
		return nil, err
	}

	fmt.Println(response.Otp) //! DELETE AFTER TEST

	// k.SetupKafka()
	//! send otp to mail logic here using kafka

	return &pb.OtpResponse{
		Email:          response.Email,
		ExpirationTime: fmt.Sprintf("%v seconds", response.ExpireTime),
	}, nil
}

func (handler *AdminAirlineHandler) VerifyAirlineRegistration(ctx context.Context, p *pb.OTPRequest) (*pb.AirlineResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.VerifyAirlineRequest(p)
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

func (handler *AdminAirlineHandler) FetchAllAirlines(context.Context, *pb.EmptyRequest) (*pb.AirlinesResponse, error) {
	return &pb.AirlinesResponse{}, nil
}

func (handler *AdminAirlineHandler) FetchAirline(context.Context, *pb.IDRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}

func (handler *AdminAirlineHandler) UpdateAirline(context.Context, *pb.AirlineRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}

func (handler *AdminAirlineHandler) DeleteAirline(context.Context, *pb.IDRequest) (*pb.AirlineResponse, error) {
	return &pb.AirlineResponse{}, nil
}
