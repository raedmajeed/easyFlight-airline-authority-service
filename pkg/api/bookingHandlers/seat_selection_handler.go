package bookingHandlers

import (
	"context"
	"fmt"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (handler *BookingHandler) AddTravellerSeats(ctx context.Context, p *pb.SeatRequest) (*pb.SeatResponse, error) {
	fmt.Println("WORKING")
	return nil, nil
}
