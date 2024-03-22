package service_test

import (
	"gorm.io/gorm"
	"reflect"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/golang/mock/gomock"
	"github.com/raedmajeed/admin-servcie/config"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	pb "github.com/raedmajeed/admin-servcie/pkg/pb"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
	"github.com/raedmajeed/admin-servcie/pkg/service"
)

func TestCreateFlightType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redis, _ := redismock.NewClientMock()

	type args struct {
		input *dom.FlightTypeModel
	}

	tests := []struct {
		name       string
		args       args
		beforeTest func(airlineRepo *repository.MockAdminAirlineRepostory)
		want       *dom.FlightTypeModel
		wantErr    bool
	}{
		{
			name: "Adding a flight type",
			args: args{
				input: &dom.FlightTypeModel{
					Type:                "1",
					FlightModel:         "FlightModel",
					Description:         "Description",
					ManufacturerName:    "ManufacturerName",
					ManufacturerCountry: "ManufacturerCountry",
					MaxDistance:         100,
					CruiseSpeed:         100,
				},
			},
			beforeTest: func(airlineRepo *repository.MockAdminAirlineRepostory) {
				airlineRepo.EXPECT().FindFlightTypeByModel("FlightModel").Return(nil, gorm.ErrRecordNotFound)
				airlineRepo.EXPECT().CreateFlightType(&pb.FlightTypeRequest{
					FlightModel:         "FlightModel",
					Description:         "Description",
					ManufacturerName:    "ManufacturerName",
					ManufacturerCountry: "ManufacturerCountry",
					MaxDistance:         100,
					CruiseSpeed:         100,
				}).Return(&dom.FlightTypeModel{
					Type:                "1",
					FlightModel:         "FlightModel",
					Description:         "Description",
					ManufacturerName:    "ManufacturerName",
					ManufacturerCountry: "ManufacturerCountry",
					MaxDistance:         100,
					CruiseSpeed:         100,
				}, nil)
			},
			want: &dom.FlightTypeModel{
				Type:                "1",
				FlightModel:         "FlightModel",
				Description:         "Description",
				ManufacturerName:    "ManufacturerName",
				ManufacturerCountry: "ManufacturerCountry",
				MaxDistance:         100,
				CruiseSpeed:         100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repository.NewMockAdminAirlineRepostory(ctrl)
			tt.beforeTest(mockRepo)

			svc := service.NewAdminAirlineService(mockRepo, redis, &config.ConfigParams{}, config.KafkaWriter{})

			got, err := svc.CreateFlightType(&pb.FlightTypeRequest{
				FlightModel:         tt.args.input.FlightModel,
				Description:         tt.args.input.Description,
				ManufacturerName:    tt.args.input.ManufacturerName,
				ManufacturerCountry: tt.args.input.ManufacturerCountry,
				MaxDistance:         tt.args.input.MaxDistance,
				CruiseSpeed:         tt.args.input.CruiseSpeed,
			})

			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFlightType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFlightType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
