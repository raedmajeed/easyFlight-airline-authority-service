package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"google.golang.org/grpc/metadata"
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
		log.Printf("Unable to login %v of email == %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	if token == "" {
		log.Printf("Unable to login %v of email ++ %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	return utils.ConvertLoginRequestToResponse(token, p), nil
}

func (handler *AdminAirlineHandler) RegisterForgotPasswordRequest(ctx context.Context, p *pb.ForgotPasswordRequest) (*pb.OtpResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.AirlineForgotPassword(p)
	if err != nil {
		log.Printf("unable to send otp, please try again")
	}

	if err != nil {
		log.Println("unable to parse duration:", err.Error())
    return nil, err
	}

	return &pb.OtpResponse{
		Email:          response.Email,
		ExpirationTime: fmt.Sprintf("%v seconds", response.ExpireTime),
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

	md, check := metadata.FromIncomingContext(ctx)
	if !check {
		log.Println("unable to read metadata from context")
		return nil, errors.New("unable to read metadata from context")
	}

	emails := md.Get("registered_email")
	email := emails[0]
	response, err := handler.svc.UpdateAirlinePassword(p, email)
	if err != nil {
		log.Printf("unable to send otp, please try again")
		return nil, err
	}

	return &pb.EmailResponse{
		Email: response,
	}, nil
}
