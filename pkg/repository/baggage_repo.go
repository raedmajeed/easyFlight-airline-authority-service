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
		FareClass:           string(pb.Class(p.Class)),
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
			log.Printf("Duplicate Key found of flight type %v", id)
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

func (repo *AdminAirlineRepositoryStruct) FetchAllAirlineBaggages(id uint) ([]dom.AirlineBaggage, error) {
	var seats []dom.AirlineBaggage
	if err := repo.DB.Where("airline_id = ?", id).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}
func (repo *AdminAirlineRepositoryStruct) FetchAirlineBaggage(id uint, sid string) (dom.AirlineBaggage, error) {
	var seat dom.AirlineBaggage
	if err := repo.DB.Where("airline_id = ? AND id = ?", id, sid).First(&seat).Error; err != nil {
		return dom.AirlineBaggage{}, err
	}
	return seat, nil
}
func (repo *AdminAirlineRepositoryStruct) DeleteAirlineBaggage(id uint, sid string) error {
	result := repo.DB.Where("airline_id = ?", id).Delete(&dom.AirlineBaggage{}, sid)
	return result.Error
}
