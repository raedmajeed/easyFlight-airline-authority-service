package repository

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) CreateAirlineCancellationPolicy(p *pb.AirlineCancellationRequest, id int) (*dom.AirlineCancellation, error) {
	cancellation := &dom.AirlineCancellation{
		AirlineId:                  id,
		FareClass:                  string(pb.Class(p.Class)),
		CancellationDeadlineBefore: int(p.CancellationDeadlineBeforeHours),
		CancellationPercentage:     int(p.CancellationPercentage),
		Refundable:                 p.Refundable,
	}
	result := repo.DB.Create(&cancellation)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", id)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return cancellation, nil
}

func (repo *AdminAirlineRepositoryStruct) FindAirlineCancellationByid(id int32) (*dom.AirlineCancellation, error) {
	var cancellation dom.AirlineCancellation
	result := repo.DB.Where("id = ?", id).First(&cancellation)
	if result.Error != nil {
		log.Println("Unable to fetch the flight types")
		return nil, result.Error
	}
	return &cancellation, nil
}

func (repo *AdminAirlineRepositoryStruct) FetchAllAirlineCancellations(id uint) ([]dom.AirlineCancellation, error) {
	var seats []dom.AirlineCancellation
	if err := repo.DB.Where("airline_id = ?", id).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}
func (repo *AdminAirlineRepositoryStruct) FetchAirlineCancellation(id uint, sid string) (dom.AirlineCancellation, error) {
	var seat dom.AirlineCancellation
	if err := repo.DB.Where("airline_id = ? AND id = ?", id, sid).First(&seat).Error; err != nil {
		return dom.AirlineCancellation{}, err
	}
	return seat, nil
}
func (repo *AdminAirlineRepositoryStruct) DeleteAirlineCancellation(id uint, sid string) error {
	result := repo.DB.Where("airline_id = ?", id).Delete(&dom.AirlineCancellation{}, sid)
	return result.Error
}
