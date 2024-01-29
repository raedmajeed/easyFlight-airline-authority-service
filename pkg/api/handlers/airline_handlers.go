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
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.RegisterAirlineSvc(p)
	if err != nil {
		log.Printf("Unable to create, err: %v", err.Error())
		return nil, err
	}

	fmt.Println(response.Otp) //! DELETE AFTER TEST

	return &pb.OtpResponse{
		Email:          response.Email,
		ExpirationTime: fmt.Sprintf("%v seconds", response.ExpireTime),
	}, nil
}

func (handler *AdminAirlineHandler) VerifyAirlineRegistration(ctx context.Context, p *pb.OTPRequest) (*pb.AirlineResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.VerifyAirlineRequest(p)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else {
			return nil, err
		}
	}

	return utils.ConvertAirlineToResponse(response), nil
}

func (handler *AdminAirlineHandler) FetchAllAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.FetchAllAirlines(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) GetAcceptedAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	log.Println("reached here 2")
	response, err := handler.svc.AcceptedAirlines(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (handler *AdminAirlineHandler) GetRejectedAirlines(ctx context.Context, p *pb.EmptyRequest) (*pb.AirlinesResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	response, err := handler.svc.RejectedAirlines(ctx, p)
	if err != nil {
		return nil, err
	}
	return response, nil
}
