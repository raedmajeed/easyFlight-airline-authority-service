package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE CANCELATION POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineCancellationPolicy(p *pb.AirlineCancellationRequest, id int) (*dom.AirlineCancellation, error) {
	_, err := svc.repo.FindAirlineById(int32(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", p.AirlineId)
			return nil, err
		} else {
			log.Printf("Cancellation policy not create of model %v, err: %v", p.AirlineId, err.Error())
			return nil, err
		}
	}

	airlineCancellationPolicy, err := svc.repo.CreateAirlineCancellationPolicy(p, id)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				id, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineCancellationPolicy, nil
}
