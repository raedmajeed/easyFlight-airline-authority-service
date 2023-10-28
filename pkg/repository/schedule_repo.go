package repository

import (
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (repo *AdminAirlineRepositoryStruct) CreateSchedules(p *pb.ScheduleRequest) (*dom.Schedule, error) {
	schedule := &dom.Schedule{
		DepartureAirport: p.DepartureAirport,
		ArrivalAirport:   p.ArrivalAirport,
		DepartureTime:    p.DepartureTime,
		ArrivalTime:      p.ArrivalTime,
	}

	result := repo.DB.Create(&schedule)
	if result.Error != nil {
		log.Println("unable to create schedule in db at repo folder")
		return nil, result.Error
	}
	return schedule, nil
}
