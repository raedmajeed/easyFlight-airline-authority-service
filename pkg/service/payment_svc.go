package service

import (
	"context"
	"errors"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
)

func (svc *AdminAirlineServiceStruct) AddConfirmedSeatsToBooked(ctx context.Context, request *pb.ConfirmedSeatRequest) (*pb.ConfirmedSeatResponse, error) {
	directFlightIDs := request.FlightChartIdDirect
	returnFlightIDs := request.FlightChartIdIndirect
	err := confirmSeats(svc, directFlightIDs, request.Travellers, request.Economy)
	if err != nil {
		return nil, errors.New("did not update direct flight booked seat")
	}
	if len(returnFlightIDs) == 0 {
		return &pb.ConfirmedSeatResponse{}, err
	}
	err = confirmSeats(svc, returnFlightIDs, request.Travellers, request.Economy)
	if err != nil {
		return nil, errors.New("did not update return booked seat")
	}
	return &pb.ConfirmedSeatResponse{}, err
}

func confirmSeats(svc *AdminAirlineServiceStruct, flights []int32, travellerCount int32, economy bool) error {
	for _, f := range flights {
		bookedSeatResponse, err := svc.repo.FindBookedSeatsByChartID(uint(f))
		if err != nil {
			return err
		}
		if economy {
			bookedSeatResponse.EconomySeatBooked = bookedSeatResponse.EconomySeatTotal + int(travellerCount)
		} else {
			bookedSeatResponse.BusinessSeatBooked = bookedSeatResponse.BusinessSeatTotal + int(travellerCount)
		}
		err = svc.repo.UpdateBookedSeats(*bookedSeatResponse, int(bookedSeatResponse.ID))
		if err != nil {
			return err
		}
	}
	return nil
}
