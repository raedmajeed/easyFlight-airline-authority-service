package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

//* METHODS TO EVERYTHING AIRLINE

func (svc *AdminAirlineServiceStruct) RegisterFlight(p *pb.AirlineRequest) (*dom.Airline, error) {
	airline := &dom.Airline{
		AirlineName:         p.AirlineName,
		CompanyAddress:      p.CompanyAddress,
		PhoneNumber:         p.PhoneNumber,
		Email:               p.Email,
		AirlineCode:         p.AirlineCode,
		AirlineLogoLink:     p.AirlineLogoLink,
		SupportDocumentLink: p.SupportDocumentsLink,
	}
	otp := utils.GenerateOTP()
	otpData := dom.OtpData{
		Otp:     otp,
		Email:   airline.Email,
		Airline: *airline,
	}

	otpJson, err := json.Marshal(&otpData)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	svc.redis.Set(context.Background(), "airline_data", otpJson, time.Second*10000)
	return airline, nil
}
