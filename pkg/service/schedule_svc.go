package service

import (
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) CreateSchedules(p *pb.ScheduleRequest) (*dom.Schedule, error) {
	schedules, err := svc.repo.CreateSchedules(p)
	if err != nil {
		log.Printf("unable to create schedule, err: %v", err.Error())
		return nil, err
	}
	return schedules, nil
}
