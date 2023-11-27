package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

func (svc *AdminAirlineServiceStruct) AirlineLogin(p *pb.LoginRequest) (string, error) {

	//! what is the airline has login token and when logged in gets the account locked. Handle that case

	airline, err := svc.repo.FindAirlineByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}
	if airline.IsAccountLocked {
		log.Printf("airline account of user %v is locked, please contact dgca for further enquiry", airline.Email)
		return "", fmt.Errorf("airline account of user %v is locked, please contact dgca for further enquiry", airline.Email)
	}

	check := utils.CheckPasswordMatch([]byte(airline.Password), p.Password)
	if !check {
		log.Printf("password mismatch for user %v", p.Email)
		return "", fmt.Errorf("password mismatch for user %v", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, p.Role, svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return "", err
	}
	return token, err
}

func (svc *AdminAirlineServiceStruct) AdminLogin(p *pb.LoginRequest) (string, error) {
	admin, err := svc.repo.FindAdminByEmail(p)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}

	check := utils.CheckPasswordMatch([]byte(admin.Password), p.Password)
	if !check {
		log.Printf("password mismatch for user %v", p.Email)
		return "", fmt.Errorf("password mismatch for user %v", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, p.Role, svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return "", err
	}
	return token, err
}

func (svc *AdminAirlineServiceStruct) AirlineForgotPassword(p *pb.ForgotPasswordRequest) (*dom.OtpData, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return nil, err
		} else {
			log.Printf("unable to generate otp %v, err: %v", p.Email, err.Error())
			return nil, err
		}
	}
	otp := utils.GenerateOTP()
	data := utils.WriteMessageToEmail(fmt.Sprintf("%v", otp), p.Email)
	byteData, err := json.Marshal(data)

	fmt.Println(otp)
	err = svc.kfk.EmailWriter.WriteMessages(context.Background(), kafka.Message{
		Value: byteData,
	})

	if err != nil {
		return nil, err
	}

	expireTime := time.Now().Add(time.Second * 60)
	forgotEmailId := fmt.Sprintf("forgot_password_%v", airline.ID)

	redisOtp := dom.OtpData{
		Otp:        otp,
		Email:      p.Email,
		ExpireTime: expireTime,
	}
	redisJSON, err := json.Marshal(redisOtp)
	if err != nil {
		log.Println("error parsing to json")
		return nil, err
	}
	svc.redis.Set(context.Background(), forgotEmailId, redisJSON, time.Second*60)
	return &dom.OtpData{
		Email:      p.Email,
		ExpireTime: expireTime,
	}, nil
}

func (svc *AdminAirlineServiceStruct) VerifyOTP(p *pb.OTPRequest) (*dom.LoginResponse, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return nil, err
		} else {
			log.Printf("unable to generate otp %v, err: %v", p.Email, err.Error())
			return nil, err
		}
	}

	var otpData dom.OtpData
	forgot_email_id := fmt.Sprintf("forgot_password_%v", airline.ID)
	redisVal := svc.redis.Get(context.Background(), forgot_email_id)
	if redisVal.Err() != nil {
		log.Printf("unable to get value from redis, otp not verified %v", redisVal.Err().Error())
		return nil, err
	}
	if err := json.Unmarshal([]byte(redisVal.Val()), &otpData); err != nil {
		log.Println("unmarshaling jsom failed")
		return nil, err
	}

	if p.Email != otpData.Email || p.Otp != int32(otpData.Otp) {
		log.Printf("error validating otp for user %v", p.Email)
		return nil, fmt.Errorf("error validating otp for user %v, wrong otp entered", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, "airline", svc.cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &dom.LoginResponse{
		Email: p.Email,
		Token: token,
	}, nil
}

func (svc *AdminAirlineServiceStruct) UpdateAirlinePassword(p *pb.ConfirmPasswordRequest, email string) (string, error) {
	airline, err := svc.repo.FindAirlineByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", email)
			return "", err
		} else {
			log.Printf("unable to generate otp %v, err: %v", email, err.Error())
			return "", err
		}
	}

	if p.ConfirmPassword != p.Password {
		log.Println("paswword dosn't match")
		return "", errors.New("password does not match")
	}

	hashedPass, err := utils.HashPassword(p.Password)
	if err != nil {
		log.Println("error hashing password")
		return "", err
	}
	airline.Password = string(hashedPass)
	email, err = svc.repo.UpdateAirlinePassword(airline)

	if err != nil {
		log.Printf("Failed to update password")
		return "", err
	}
	return email, nil
}
