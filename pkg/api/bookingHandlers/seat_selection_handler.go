package bookingHandlers

import (
	"context"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"log"
)

func (handler *BookingHandler) RegisterSelectSeat(ctx context.Context, p *pb.SeatRequest) (*pb.SeatResponse, error) {
	response, err := handler.svc.SelectAndBookSeats(ctx, p)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, err
}
