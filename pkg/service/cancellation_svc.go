package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE CANCELATION POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineCancellationPolicy(p *pb.AirlineCancellationRequest) (*dom.AirlineCancellation, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", airline.ID)
			return nil, err
		} else {
			log.Printf("Cancellation policy not create of model %v, err: %v", airline.ID, err.Error())
			return nil, err
		}
	}

	airlineCancellationPolicy, err := svc.repo.CreateAirlineCancellationPolicy(p, int(airline.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				airline.ID, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineCancellationPolicy, nil
}
