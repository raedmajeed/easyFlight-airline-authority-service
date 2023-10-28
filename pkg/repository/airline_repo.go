package repository

import (
	"errors"
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) FindAirlineById(id int32) (*dom.Airline, error) {
	var airline dom.Airline
	result := repo.DB.Where("id = ?", int(id)).First(&airline)
	fmt.Println(result)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of airline %v", id)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &airline, nil
}
