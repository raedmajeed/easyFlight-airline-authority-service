package db

import (
	"fmt"
	"github.com/raedmajeed/admin-servcie/config"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDBConnect(cfg *config.ConfigParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	log.Println("here 12")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Connection to DB %s Failed, Error: %s", cfg.DBName, err)
		return nil, err
	}

	log.Println("here 1")
	err = database.AutoMigrate(
		&dom.FlightTypeModel{},
		&dom.Airline{},
		&dom.AirlineSeat{},
		&dom.AirlineCancellation{},
		&dom.AirlineBaggage{},
		&dom.Airport{},
		&dom.AdminTable{},
		&dom.Schedule{},
		&dom.FlightFleets{},
		&dom.FlightChart{},
		&dom.BookedSeat{},
	)

	if err != nil {
		log.Printf("unable to migrate db, err: %v", err)
		return nil, err
	}

	return database, nil
}
