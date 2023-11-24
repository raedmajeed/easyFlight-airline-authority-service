package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

//* METHODS TO EVERYTHING AIRLINE

func (svc *AdminAirlineServiceStruct) RegisterAirlineSvc(p *pb.AirlineRequest) (*dom.RegisterAirlineOtpData, error) {
	//go
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
	data := utils.WriteMessageToEmail(fmt.Sprintf("%v", otp), p.Email)
	byteData, err := json.Marshal(data)
	err = svc.kfk.EmailWriter.WriteMessages(context.Background(), kafka.Message{
		Value: byteData,
	})

	otpData := &dom.RegisterAirlineOtpData{
		Otp:        otp,
		Email:      airline.Email,
		ExpireTime: time.Now().Add(time.Minute * 2),
		Airline:    *airline,
	}

	otpJson, err := json.Marshal(&otpData)
	if err != nil {
		log.Printf("error parsing JSON, err: %v", err.Error())
		return nil, err
	}

	register_airline := fmt.Sprintf("register_airline_%v", p.Email)
	svc.redis.Set(context.Background(), register_airline, otpJson, time.Minute*2)

	return otpData, nil
}

func (svc *AdminAirlineServiceStruct) VerifyAirlineRequest(p *pb.OTPRequest) (*dom.Airline, error) {
	register_airline := fmt.Sprintf("register_airline_%v", p.Email)
	redisVal := svc.redis.Get(context.Background(), register_airline)

	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis err: %v", redisVal.Err().Error())
		return nil, redisVal.Err()
	}

	var otpData dom.RegisterAirlineOtpData
	err := json.Unmarshal([]byte(redisVal.Val()), &otpData)
	if err != nil {
		log.Println("unable to unmarshal json")
		return nil, err
	}

	if otpData.ExpireTime.Before(time.Now()) {
		log.Println("otp expired, try again later")
		return nil, errors.New("otp expired, try again later")
	}

	if otpData.Email != p.Email || otpData.Otp != int(p.Otp) {
		log.Printf("otp not verified for user %v", otpData.Email)
		return nil, errors.New("otp not verified")
	}

	_, err = svc.repo.FindAirlineByEmail(otpData.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Existing record found  of airline %v", p.Email)
		return nil, errors.New("airline already exists")
	}

	airline, err := svc.repo.CreateAirline(&otpData.Airline)
	if err != nil {
		return nil, err
	}
	return airline, nil
}
