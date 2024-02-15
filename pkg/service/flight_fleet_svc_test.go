package service_test

import (
	"github.com/golang/mock/gomock"
	"github.com/raedmajeed/admin-servcie/config"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
	"github.com/raedmajeed/admin-servcie/pkg/service"
	"testing"
)

func TestCreateFlightFleet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockAdminAirlineRepostory(ctrl)
	svc := service.NewAdminAirlineService(mockRepo, nil, nil, config.KafkaWriter{})

	type args struct {
		p *pb.FlightFleetRequest
	}

	tests := []struct {
		name       string
		args       args
		beforeFunc func(repo repository.MockAdminAirlineRepostory)
		want       dom.FlightFleetResponse
		wantErr    bool
	}{
		{
			name: "Test 1",
			args: args{
				p: &pb.FlightFleetRequest{
					SeatId:               int32(1),
					FlightTypeId:         int32(1),
					AirlineEmail:         "airline@example.com",
					BaggagePolicyId:      int32(1),
					CancellationPolicyId: int32(1),
				},
			},
			beforeFunc: func(repo repository.MockAdminAirlineRepostory) {
				repo.
					EXPECT().
					FindAirlineByEmail("airline@example.com").
					Return(&dom.Airline{
						Email:       "airline@example.com",
						AirlineName: "airline",
						AirlineCode: "123",
					}, nil)
				repo.
					EXPECT().
					FindFlightTypeByID(int32(1)).
					Return(&dom.FlightTypeModel{
						FlightModel: "COMMERCIAL",
						Description: "Large Flight",
					}, nil)

				repo.
					EXPECT().
					FindAirlineSeatByid(int32(1)).
					Return(&dom.AirlineSeat{
						EconomySeatNumber:  1,
						BusinessSeatNumber: 2,
					}, nil)
				repo.
					EXPECT().
					FindAirlineBaggageByid(int32(1)).
					Return(&dom.AirlineBaggage{
						HandAllowedHeight: 1,
						HandAllowedWeight: 2,
					}, nil)
				repo.
					EXPECT().
					FindAirlineCancellationByid(int32(1)).
					Return(&dom.AirlineCancellation{
						CancellationDeadlineBefore: 1,
						CancellationPercentage:     2,
					}, nil)
				repo.
					EXPECT().
					CreateFlightFleet(&dom.FlightFleets{
						AirlineID:            uint(0),
						SeatID:               uint(0),
						FlightTypeID:         uint(0),
						BaggagePolicyID:      uint(0),
						CancellationPolicyID: uint(0),
						FlightNumber:         "123-002",
					}).
					Return(nil)
				repo.
					EXPECT().
					FindLastFlightInDB().
					Return(1)
			},
			want: dom.FlightFleetResponse{
				FlightNumber: "123-002",
				AirlineName:  "airline",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.beforeFunc(*mockRepo)
			got, err := svc.CreateFlightFleet(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("svc.CreateFlightFleet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if got.FlightNumber != tt.want.FlightNumber {
					t.Errorf("svc.CreateFlightFleet() got = %v, want %v", got.FlightNumber, tt.want.FlightNumber)
				}
			}
		})
	}

}
