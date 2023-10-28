package repository

import (
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) FindAirportByAirportCode(airportCode string) (*dom.Airport, error) {
	var airport dom.Airport
	result := repo.DB.Where("airport_code = ?", airportCode).First(&airport)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of airport %v", airportCode)
			return nil, gorm.ErrRecordNotFound
		} else {
			return &airport, nil
		}
	}
	return &airport, nil
}

func (repo *AdminAirlineRepositoryStruct) CreateAirport(p *pb.Airport) (*dom.Airport, error) {
	airport := dom.Airport{
		AirportCode:  p.AirportCode,
		AirportName:  p.AirportName,
		City:         p.City,
		Country:      p.Country,
		Region:       p.Region,
		Latitude:     p.Latitude,
		Longitude:    p.Longitude,
		IATAFCSCode:  p.IataFcsCode,
		ICAOCode:     p.IcaoCode,
		Website:      p.Website,
		ContactEmail: p.ContactEmail,
		ContactPhone: p.ContactPhone,
	}
	result := repo.DB.Create(&airport)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			log.Printf("Duplicate Key found of Airport %v", p.AirportCode)
			return nil, gorm.ErrDuplicatedKey
		} else {
			return nil, result.Error
		}
	}
	return &airport, nil
}
