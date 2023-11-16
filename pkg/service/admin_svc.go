package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/raedmajeed/admin-servcie/pkg/utils"
	"github.com/segmentio/kafka-go"
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

	data := utils.SendAirlinePasswordEmail(airline.Email, airline.AirlineCode)
	byteData, err := json.Marshal(data)

	err = svc.kfk.EmailWriter.WriteMessages(context.Background(), kafka.Message{
		Value: byteData,
	})

	if err != nil {
		return nil, err
	}
	return airline, nil
}
