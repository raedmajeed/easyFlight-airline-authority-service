package repository

import (
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
)

func (repo *AdminAirlineRepositoryStruct) CreateSchedules(schedule *dom.Schedule) error {
	result := repo.DB.Create(schedule)
	if result.Error != nil {
		log.Println("unable to create schedule in db at repo folder")
		return result.Error
	}
	return nil
}

func (repo *AdminAirlineRepositoryStruct) FindScheduleByID(id int) (*dom.Schedule, error) {
	var schedule dom.Schedule
	result := repo.DB.Where("id = ?", uint(id)).Find(&schedule)
	if result.Error != nil {
		log.Println("unable to create schedule in db at repo folder")
		return nil, result.Error
	}
	return &schedule, nil
}
