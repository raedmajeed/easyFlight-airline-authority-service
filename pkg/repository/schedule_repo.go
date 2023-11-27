package repository

import (
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"log"
)

func (repo *AdminAirlineRepositoryStruct) CreateSchedules(schedule *dom.Schedule) error {
	result := repo.DB.Create(schedule)
	if result.Error != nil {
		log.Println("unable to create schedule in db at repo folder")
		return result.Error
	}
	dt := schedule.DepartureDateTime.Format("2006-01-02 15:04:05")
	at := schedule.ArrivalDateTime.Format("2006-01-02 15:04:05")
	repo.Convert(dt, at, schedule.ID)
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

func (repo *AdminAirlineRepositoryStruct) FindAllSchedules() []*dom.Schedule {
	var schedules []*dom.Schedule
	repo.DB.Find(&schedules)
	return schedules
}

func (repo *AdminAirlineRepositoryStruct) Convert(depTime, arrTime string, id uint) {
	repo.DB.Model(&dom.Schedule{}).Where("id = ?", id).Update("arrival_date_time", arrTime)
	repo.DB.Model(&dom.Schedule{}).Where("id = ?", id).Update("departure_date_time", depTime)
}
