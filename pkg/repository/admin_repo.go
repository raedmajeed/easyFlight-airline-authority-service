package repository

import (
	"errors"
	"fmt"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

func (repo *AdminAirlineRepositoryStruct) FindAdminByEmail(p *pb.LoginRequest) (*dom.AdminTable, error) {
	var admin dom.AdminTable
	result := repo.DB.Where("email = ?", p.Email).First(&admin)
	fmt.Println(result)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Record not found of admin %v", p.Email)
			return nil, gorm.ErrRecordNotFound
		} else {
			return nil, result.Error
		}
	}
	return &admin, nil
}
