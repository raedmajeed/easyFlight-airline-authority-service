package repository

import (
	"errors"
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) FindAirlineById(id int32) (*dom.Airline, error) {
	var airline dom.Airline
	result := repo.DB.Where("id = ?", int(id)).First(&airline)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of airline %v", id)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &airline, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlineByEmail(email string) (*dom.Airline, error) {
	var airline dom.Airline
	result := repo.DB.Where("email = ?", email).First(&airline)
	fmt.Println(result)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of airline %v", email)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &airline, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlinePassword(p *pb.LoginRequest) (*dom.Airline, error) {
	var user dom.Airline
	result := repo.DB.Where("email = ? AND password = ?", p.Email, p.Password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Password Mismatch %v", p.Email)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &user, nil
}

func (repo *AdminAirlineRepositoryStruct) InitialAirlinePassword(airline *dom.Airline) (string, error) {
	pswrd, err := utils.HashPassword(airline.AirlineCode)
	if err != nil {
		return "", err
	}
	result := repo.DB.Model(&dom.Airline{}).Where("id = ?", airline.ID).Update("Password", pswrd)
	if result.Error != nil {
		log.Println("Unable to Update the airline password")
		return "", result.Error
	}
	return airline.Email, nil
}

func (repo *AdminAirlineRepositoryStruct) UpdateAirlinePassword(airline *dom.Airline) (string, error) {
	result := repo.DB.Model(&dom.Airline{}).Where("id = ?", airline.ID).Update("Password", airline.Password)
	if result.Error != nil {
		log.Println("Unable to Update the airline password")
		return "", result.Error
	}
	return airline.Email, nil
}

func (repo *AdminAirlineRepositoryStruct) CreateAirline(airline *dom.Airline) (*dom.Airline, error) {
	result := repo.DB.Create(airline)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Println("duplicate key found")
			return nil, result.Error
		} else {
			return nil, result.Error
		}
	}
	return airline, nil
}

func (repo *AdminAirlineRepositoryStruct) UnlockAirlineAccount(airlineID int) error {
	return repo.DB.Model(&dom.Airline{}).Where("id = ?", airlineID).Update("IsAccountLocked", false).Error
}
