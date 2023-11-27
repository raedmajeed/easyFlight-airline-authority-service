package handlers

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	utils "github.com/raedmajeed/admin-servcie/pkg/utils"
)

func (handler *AdminAirlineHandler) RegisterScheduleRequest(ctx context.Context, p *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	log.Println("reached registering schedules function at admin service")
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	response, err := handler.svc.CreateSchedules(p)
	if err != nil {
		log.Printf("Unable to create airline seats, err: %v", err.Error())
		return nil, err
	}
	return utils.ConvertSchedulesToResponse(response), nil
}
