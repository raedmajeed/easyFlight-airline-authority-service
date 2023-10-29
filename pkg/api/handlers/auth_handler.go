package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
)

func (handler *AdminAirlineHandler) RegisterLoginRequest(ctx context.Context, p *pb.LoginRequest) (*pb.LoginResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}
	var token string
	var err error
	if p.Role == "airline" {
		token, err = handler.svc.AirlineLogin(p)
	} else {
		token, err = handler.svc.AdminLogin(p)
	}
	if err != nil {
		log.Printf("Unable to login %v of email %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	if token == "" {
		log.Printf("Unable to login %v of email %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	return utils.ConvertLoginRequestToResponse(token, p), nil
}

func (handler *AdminAirlineHandler) RegisterForgotPasswordRequest(ctx context.Context, p *pb.ForgotPasswordRequest) (*pb.EmailResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.AirlineForgotPassword(p)
	if err != nil {
		log.Printf("unable to send otp, please try again")
	}
	return &pb.EmailResponse{
		Email: response,
	}, nil
}

func (handler *AdminAirlineHandler) RegisterVerifyOTPRequest(ctx context.Context, p *pb.OTPRequest) (*pb.LoginResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.VerifyOTP(p)
	if err != nil {
		log.Printf("unable to send otp, please try again")
		return nil, err
	}

	return &pb.LoginResponse{
		Email: response.Email,
		Token: response.Token,
	}, nil
}

func (handler *AdminAirlineHandler) RegisterConfirmPasswordRequest(ctx context.Context, p *pb.ConfirmPasswordRequest) (*pb.EmailResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	registered_email := ctx.Value("registered_email")
	response, err := handler.svc.UpdateAirlinePassword(p, fmt.Sprintf("%d", registered_email))
	if err != nil {
		log.Printf("unable to send otp, please try again")
		return nil, err
	}

	return &pb.EmailResponse{
		Email: response,
	}, nil
}
