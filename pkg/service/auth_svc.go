package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

func (svc *AdminAirlineServiceStruct) AirlineLogin(p *pb.LoginRequest) (string, error) {
	_, err := svc.repo.FindAirlineByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}

	_, err = svc.repo.FindAirlinePassword(p)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Password Mismatch %v", p.Email)
			return "", fmt.Errorf("password mismatc for user %v", p.Email)
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", fmt.Errorf("password mismatc for user %v", p.Email)
		}
	}

	token, err := utils.GenerateToken(p.Email, p.Role, svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return "", err
	}
	return token, err
}

func (svc *AdminAirlineServiceStruct) AdminLogin(p *pb.LoginRequest) (string, error) {
	_, err := svc.repo.FindAdminByEmail(p)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}

	_, err = svc.repo.FindAdminPassword(p)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Password Mismatch %v", p.Email)
			return "", fmt.Errorf("password mismatch for admin %v", p.Email)
		} else {
			log.Printf("unable to login %v, err: %v", p.Email, err.Error())
			return "", fmt.Errorf("password mismatch for admin %v", p.Email)
		}
	}

	token, err := utils.GenerateToken(p.Email, p.Role, svc.cfg)
	if err != nil {
		log.Printf("unable to generate token for user %v, err: %v", p.Email, err.Error())
		return "", err
	}
	return token, err
}

func (svc *AdminAirlineServiceStruct) AirlineForgotPassword(p *pb.ForgotPasswordRequest) (string, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found og %v", p.Email)
			return "", err
		} else {
			log.Printf("unable to generate otp %v, err: %v", p.Email, err.Error())
			return "", err
		}
	}
	otp := utils.GenerateOTP()

	// send otp to mail here

	forgot_email_id := fmt.Sprintf("forgot_email_%v", airline.ID)

	redis_otp := dom.OtpData{
		Otp:        otp,
		Email:      p.Email,
		ExpireTime: time.Now().Add(time.Second * 60),
	}
	redisJSON, err := json.Marshal(redis_otp)
	if err != nil {
		log.Println("error parsing to json")
		return "", err
	}
	redisStatus := svc.redis.Set(context.Background(), forgot_email_id, redisJSON, time.Second*60)
	if redisStatus.Err() != nil {
		log.Printf("error passing value to redis, err: %v", redisStatus.Err().Error())
		return "", redisStatus.Err()
	}
	return p.Email, nil
}

func (svc *AdminAirlineServiceStruct) VerifyOTP(p *pb.OTPRequest) (*dom.LoginReponse, error) {
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
	forgot_email_id := fmt.Sprintf("forgot_email_%v", airline.ID)
	redisVal, err := svc.redis.Get(context.Background(), forgot_email_id).Result()
	if err != nil {
		log.Println("unable to get value from redis, otp not veriefied")
		return nil, err
	}
	if err := json.Unmarshal([]byte(redisVal), &otpData); err != nil {
		log.Println("unmarshaling jsom failed")
		return nil, err
	}

	if p.Email != otpData.Email && p.Otp != int32(otpData.Otp) {
		log.Printf("error validating otp for user %v", p.Email)
		return nil, fmt.Errorf("error validating otp for user %v, wrong otp entered", p.Email)
	}

	token, err := utils.GenerateToken(p.Email, "airline", svc.cfg)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &dom.LoginReponse{
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

	airline.Password = p.Password
	email, err = svc.repo.UpdateAirlinePassword(airline)
	if err != nil {
		log.Printf("Failed to update password")
		return "", err
	}
	return email, nil
}
