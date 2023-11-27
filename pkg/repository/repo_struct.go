package repository

import (
	"github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"

	// interfaces "github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type AdminAirlineRepositoryStruct struct {
	DB *gorm.DB
}

func NewAdminAirlineRepository(db *gorm.DB) interfaces.AdminAirlineRepostory {
	return &AdminAirlineRepositoryStruct{
		DB: db,
	}
}
