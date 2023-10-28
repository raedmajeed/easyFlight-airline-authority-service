package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE BAGGAGE POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineBaggagePolicy(p *pb.AirlineBaggageRequest, id int) (*dom.AirlineBaggage, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Baggage Policy not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	airlineBaggagePolicy, err := svc.repo.CreateAirlineBaggagePolicy(p, id)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				id, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineBaggagePolicy, nil
}
