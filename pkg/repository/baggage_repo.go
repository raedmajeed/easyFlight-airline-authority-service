package repository

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) CreateAirlineBaggagePolicy(p *pb.AirlineBaggageRequest, id int) (*dom.AirlineBaggage, error) {
	baggage := &dom.AirlineBaggage{
		AirlineId:           id,
		FareClass:           int(p.Class),
		CabinAllowedWeight:  int(p.CabinAllowedWeight),
		CabinAllowedLength:  int(p.CabinAllowedLength),
		CabinAllowedBreadth: int(p.CabinAllowedBreadth),
		CabinAllowedHeight:  int(p.CabinAllowedHeight),
		HandAllowedWeight:   int(p.HandAllowedWeight),
		HandAllowedLength:   int(p.HandAllowedLength),
		HandAllowedBreadth:  int(p.HandAllowedBreadth),
		HandAllowedHeight:   int(p.HandAllowedHeight),
		FeeExtraPerKGCabin:  int(p.FeeForExtraKgCabin),
		FeeExtraPerKGHand:   int(p.FeeForExtraKgHand),
		Restrictions:        p.Restrictions,
	}

	result := repo.DB.Create(&baggage)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", p.AirlineId)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return baggage, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlineBaggageByid(id int32) (*dom.AirlineBaggage, error) {
	var baggage dom.AirlineBaggage
	result := repo.DB.Where("id = ?", id).First(&baggage)
	if result.Error != nil {
		log.Println("Unable to fetch the flight types")
		return nil, result.Error
	}
	return &baggage, nil
}
