package service

import (
	"context"
	"errors"
	"log"

	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"gorm.io/gorm"
)

// CreateAirlineBaggagePolicy * METHODS TO EVERYTHING AIRLINE BAGGAGE POLICY
func (svc *AdminAirlineServiceStruct) CreateAirlineBaggagePolicy(p *pb.AirlineBaggageRequest) (*dom.AirlineBaggage, error) {
	airline, err := svc.repo.FindAirlineByEmail(p.AirlineEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No existing record found  of airline %v", airline.ID)
			return nil, err
		} else {
			log.Printf("Baggage Policy not create of model %v, err: %v", airline.ID, err.Error())
			return nil, err
		}
	}

	airlineBaggagePolicy, err := svc.repo.CreateAirlineBaggagePolicy(p, int(airline.ID))
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			log.Printf("Flight Type not created of model %v, err: %v",
				airline.ID, err.Error())
			return nil, err
		} else {
			return nil, err
		}
	}
	return airlineBaggagePolicy, nil
}

func (svc *AdminAirlineServiceStruct) FetchAllAirlineBaggages(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggagesResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAllAirlineBaggages(airline.ID)
	if err != nil {
		return nil, err
	}
	result := ConvertToResponseBaggage(resp)
	return result, err
}
func (svc *AdminAirlineServiceStruct) FetchAirlineBaggage(ctx context.Context, p *pb.FetchRequest) (*pb.AirlineBaggageResponse, error) {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	resp, err := svc.repo.FetchAirlineBaggage(airline.ID, p.Id)
	if err != nil {
		return nil, err
	}
	return &pb.AirlineBaggageResponse{
		AirlineBaggage: &pb.AirlineBaggageRequest{
			CabinAllowedLength:  int32(resp.CabinAllowedLength),
			CabinAllowedBreadth: int32(resp.CabinAllowedBreadth),
			CabinAllowedWeight:  int32(resp.CabinAllowedWeight),
			CabinAllowedHeight:  int32(resp.CabinAllowedHeight),
			HandAllowedLength:   int32(resp.HandAllowedLength),
			HandAllowedBreadth:  int32(resp.HandAllowedBreadth),
			HandAllowedWeight:   int32(resp.HandAllowedWeight),
			HandAllowedHeight:   int32(resp.HandAllowedHeight),
			FeeForExtraKgHand:   int32(resp.FeeExtraPerKGHand),
			FeeForExtraKgCabin:  int32(resp.FeeExtraPerKGCabin),
		},
	}, nil
}
func (svc *AdminAirlineServiceStruct) DeleteAirlineBaggage(ctx context.Context, p *pb.FetchRequest) error {
	airline, _ := svc.repo.FindAirlineByEmail(p.Email)
	err := svc.repo.DeleteAirlineBaggage(airline.ID, p.Id)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToResponseBaggage(data []dom.AirlineBaggage) *pb.AirlineBaggagesResponse {
	var result []*pb.AirlineBaggageRequest
	for _, resp := range data {
		result = append(result, &pb.AirlineBaggageRequest{
			CabinAllowedLength:  int32(resp.CabinAllowedLength),
			CabinAllowedBreadth: int32(resp.CabinAllowedBreadth),
			CabinAllowedWeight:  int32(resp.CabinAllowedWeight),
			CabinAllowedHeight:  int32(resp.CabinAllowedHeight),
			HandAllowedLength:   int32(resp.HandAllowedLength),
			HandAllowedBreadth:  int32(resp.HandAllowedBreadth),
			HandAllowedWeight:   int32(resp.HandAllowedWeight),
			HandAllowedHeight:   int32(resp.HandAllowedHeight),
			FeeForExtraKgHand:   int32(resp.FeeExtraPerKGHand),
			FeeForExtraKgCabin:  int32(resp.FeeExtraPerKGCabin),
		})
	}
	return &pb.AirlineBaggagesResponse{
		AirlineBaggages: result,
	}
}
