package bookingHandlers

import (
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
)

type BookingHandler struct {
	svc interfaces.AdminAirlineService
	pb.AdminServiceServer
}

func NewBookingHandler(svc interfaces.AdminAirlineService) *BookingHandler {
	return &BookingHandler{
		svc: svc,
	}
}
