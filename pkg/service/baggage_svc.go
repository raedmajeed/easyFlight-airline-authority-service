package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// CreateAirlineBaggagePolicy * METHODS TO EVERYTHING AIRLINE BAGGAGE POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineBaggagePolicy(p *pb.AirlineBaggageRequest) (*dom.AirlineBaggage, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", airline.ID)
			return nil, err
		} else {
			log.Printf("Baggage Policy not create of model %v, err: %v", airline.ID, err.Error())
			return nil, err
		}
	}

	airlineBaggagePolicy, err := svc.repo.CreateAirlineBaggagePolicy(p, int(airline.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				airline.ID, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineBaggagePolicy, nil
}
