package service

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (svc *AdminAirlineServiceStruct) CreateAirport(p *pb.Airport) (*dom.Airport, error) {
	airport, err := svc.repo.FindAirportByAirportCode(p.AirportCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of model %v", p.AirportCode)
		} else {
			log.Printf("Airport not created of code %v, err: %v", p.AirportCode, err.Error())
			return airport, err
		}
	}

	airport, err = svc.repo.CreateAirport(p)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				p.AirportCode, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airport, nil
}
