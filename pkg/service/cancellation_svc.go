package service

import (
	"context"
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// * METHODS TO EVERYTHING AIRLINE CANCELATION POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineCancellationPolicy(p *pb.AirlineCancellationRequest) (*dom.AirlineCancellation, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", airline.ID)
			return nil, err
		} else {
			log.Printf("Cancellation policy not create of model %v, err: %v", airline.ID, err.Error())
			return nil, err
		}
	}

	airlineCancellationPolicy, err := svc.repo.CreateAirlineCancellationPolicy(p, int(airline.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				airline.ID, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineCancellationPolicy, nil
}

func (svc *AdminAirlineServiceStruct) FetchAllAirlineCancellations(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationsResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAllAirlineCancellations(airline.ID)
	if err != nil {
		return nil, err
	}
	result := ConvertToResponseCancel(resp)
	return result, err
}
func (svc *AdminAirlineServiceStruct) FetchAirlineCancellation(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineCancellationResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAirlineCancellation(airline.ID, p.Id)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineCancellationResponse{
		AirlineCancellation: &pb.AirlineCancellationRequest{
			CancellationPercentage:          int32(resp.CancellationPercentage),
			CancellationDeadlineBeforeHours: uint32(resp.CancellationDeadlineBefore),
			Refundable:                      resp.Refundable,
		},
	}, nil
}
func (svc *AdminAirlineServiceStruct) DeleteAirlineCancellation(ctx context.Context, p *pb.FetchRequest) error {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	err := svc.repo.DeleteAirlineCancellation(airline.ID, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToResponseCancel(data []dom.AirlineCancellation) *pb.AirlineCancellationsResponse {
	var result []*pb.AirlineCancellationRequest
	for _, resp := range data {
		result = append(result, &pb.AirlineCancellationRequest{
			CancellationPercentage:          int32(resp.CancellationPercentage),
			CancellationDeadlineBeforeHours: uint32(resp.CancellationDeadlineBefore),
			Refundable:                      resp.Refundable,
		})
	}
	return &pb.AirlineCancellationsResponse{
		AirlineCancellations: result,
	}
}
