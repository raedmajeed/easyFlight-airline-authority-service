package handlers

import (
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
)

type AdminAirlineHandler struct {
	// need service here
	// need jwt utils token generatore here
	svc interfaces.AdminAirlineService
	pb.AdminAirlineServer
}

func NewAdminAirlineHandler(svc interfaces.AdminAirlineService) *AdminAirlineHandler {
	return &AdminAirlineHandler{
		svc: svc,
	}
}
