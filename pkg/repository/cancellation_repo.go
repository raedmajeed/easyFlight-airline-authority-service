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
		FareClass:                  int(p.Class),
		CancellationDeadlineBefore: int(p.CancellationDeadlineBeforeHours),
		CancellationPercentage:     int(p.CancellationPercentage),
		Refundable:                 p.Refundable,
	}
	result := repo.DB.Create(&cancellation)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of flight type %v", p.AirlineId)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return cancellation, nil
}
