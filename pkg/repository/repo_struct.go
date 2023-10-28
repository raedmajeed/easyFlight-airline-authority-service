package repository

import (
	"fmt"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"

	// interfaces "github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"
	"gorm.io/gorm"
)

func (a *AdminAirlineRepositoryStruct) AddFlightType(p *pb.FlightTypeRequest) {
	fmt.Println("REACHED REPO")
	fmt.Println(p.Description)
	// a.DB.Create()
}

type AdminAirlineRepositoryStruct struct {
	DB *gorm.DB
}

func NewAdminAirlineRepository(db *gorm.DB) interfaces.AdminAirlineRepostory {
	return &AdminAirlineRepositoryStruct{
		DB: db,
	}
}
