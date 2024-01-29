package bookingHandlers

import (
	"context"
	"errors"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"log"
	"time"
)

func (handler *BookingHandler) RegisterSearchFlight(ctx context.Context, p *pb.SearchFlightRequestAdmin) (*pb.SearchFlightResponseAdmin, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed")
	}

	resp, err := handler.svc.SearchFlightInitial(ctx, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (handler *BookingHandler) RegisterSelectFlight(ctx context.Context, p *pb.SelectFlightAdmin) (*pb.CompleteFlightDetails, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		return nil, errors.New("deadline passed")
	}

	log.Println("here ====")
	resp, err := handler.svc.SearchSelectFlight(ctx, p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
