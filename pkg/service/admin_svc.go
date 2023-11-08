package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"gorm.io/gorm"
)


func (svc *AdminAirlineServiceStruct) AdminVerifyAirlineRequest(airlineId int) (*dom.Airline, error) {
	airline, err := svc.repo.FindAirlineById(int32(airlineId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("airline not found")
			return nil, gorm.ErrRecordNotFound
		} else {
			log.Printf("db error: %v", err.Error())
			return nil, err
		}
	}

	err = svc.repo.UnlockAirlineAccount(airlineId)
	if err != nil {
		log.Printf("unable to unlock airline err: %v", err.Error())
		return nil, err
	}

	_, err = svc.repo.InitialAirlinePassword(airline)
	if err != nil {
		log.Println("unable to update airline password")
		return nil, err
	}

	//! logic to send email and password to airline password, airline will have the provision to channge pswrd using forgot pswrd
	// email := airline.Email
	// passwrod := airline.AirlineCode

	return airline, nil
}