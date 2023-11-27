package handlers

import (
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
)

type AdminAirlineHandler struct {
	svc interfaces.AdminAirlineService
	pb.AdminAirlineServer
}

func NewAdminAirlineHandler(svc interfaces.AdminAirlineService) *AdminAirlineHandler {
	return &AdminAirlineHandler{
		svc: svc,
	}
}
