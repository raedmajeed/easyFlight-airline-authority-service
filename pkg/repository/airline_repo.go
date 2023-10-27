package repository

import (
	"fmt"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (a *AdminAirlineRepositoryStruct) AddFlightType(p *pb.FlightTypeRequest) {
	fmt.Println("REACHED REPO")
	fmt.Println(p.Description)
	// a.DB.Create()
}
