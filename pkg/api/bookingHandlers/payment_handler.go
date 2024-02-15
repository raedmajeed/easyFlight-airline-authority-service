package bookingHandlers

import (
	"context"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (handler *BookingHandler) AddConfirmedSeats(ctx context.Context, request *pb.ConfirmedSeatRequest) (*pb.ConfirmedSeatResponse, error) {
	response, err := handler.svc.AddConfirmedSeatsToBooked(ctx, request)
	return response, err
}
