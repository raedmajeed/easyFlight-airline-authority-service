package handlers_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/service"
	"reflect"
	"testing"
)

func TestGetAirport(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		*pb.AirportRequest
	}

	tests := []struct {
		name       string
		args       args
		beforeFunc func(svc *service.MockAdminAirlineService)
		want       *pb.AirportResponse
		wantErr    bool
	}{
		{
			name: "Test 1",
			args: args{
				&pb.AirportRequest{
					AirportCode: "DEL",
				},
			},
			beforeFunc: func(svc *service.MockAdminAirlineService) {
				svc.EXPECT().GetAirport(context.Background(), &pb.AirportRequest{
					AirportCode: "DEL",
				}).Return(&pb.AirportResponse{
					Airport: &pb.Airport{
						AirportCode:  "DEL",
						AirportName:  "Delhi Airport",
						City:         "Delhi",
						Country:      "India",
						Region:       "Asia",
						Latitude:     28.556162,
						Longitude:    77.100635,
						IataFcsCode:  "DEL",
						IcaoCode:     "VIDP",
						Website:      "https://example.com",
						ContactEmail: "contact@example.com",
						ContactPhone: "+1234567890",
					},
				}, nil)
			},
			want: &pb.AirportResponse{
				Airport: &pb.Airport{
					AirportCode:  "DEL",
					AirportName:  "Delhi Airport",
					City:         "Delhi",
					Country:      "India",
					Region:       "Asia",
					Latitude:     28.556162,
					Longitude:    77.100635,
					IataFcsCode:  "DEL",
					IcaoCode:     "VIDP",
					Website:      "https://example.com",
					ContactEmail: "contact@example.com",
					ContactPhone: "+1234567890",
				},
			},
			wantErr: false,
		},
		{
			name: "Test 2",
			args: args{
				&pb.AirportRequest{
					AirportCode: "DEL",
				},
			},
			beforeFunc: func(svc *service.MockAdminAirlineService) {
				svc.EXPECT().GetAirport(context.Background(), &pb.AirportRequest{
					AirportCode: "DEL",
				}).Return(&pb.AirportResponse{}, errors.New("oops"))
			},
			want:    &pb.AirportResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewMockAdminAirlineService(ctrl)
			tt.beforeFunc(svc)

			got, err := svc.GetAirport(context.Background(), tt.args.AirportRequest)

			if (err != nil) != tt.wantErr {
				t.Errorf("services.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error", got, tt.want)
				return
			}
		})
	}
}
