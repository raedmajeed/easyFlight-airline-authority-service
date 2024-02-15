package service_test

//
//import (
//	"github.com/golang/mock/gomock"
//	"github.com/raedmajeed/admin-servcie/config"
//	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
//	"github.com/raedmajeed/admin-servcie/pkg/repository"
//	"github.com/raedmajeed/admin-servcie/pkg/service"
//	"testing"
//)
//
//func TestAdminVerifyAirlineRequest(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := repository.NewMockAdminAirlineRepostory(ctrl)
//	svc := service.NewAdminAirlineService(mockRepo, nil, &config.ConfigParams{}, config.KafkaWriter{})
//
//	beforeFunc := func(repo *repository.MockAdminAirlineRepostory) {
//		repo.EXPECT().FindAirlineById(int32(1)).Return(
//			&dom.Airline{
//				Email:          "raed786@gmail.com",
//				AirlineName:    "RAED",
//				CompanyAddress: "RAED",
//				PhoneNumber:    "7902498141",
//				AirlineCode:    "IX99",
//			}, nil,
//		)
//		repo.EXPECT().UnlockAirlineAccount(1).Return(nil)
//		repo.EXPECT().InitialAirlinePassword(&dom.Airline{
//			Email:          "raed786@gmail.com",
//			AirlineName:    "RAED",
//			CompanyAddress: "RAED",
//			PhoneNumber:    "7902498141",
//			AirlineCode:    "IX99",
//		})
//	}
//
//	type args struct {
//		id int
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    dom.Airline
//		wantErr bool
//	}{
//		{
//			name:    "Test Case 1",
//			args:    args{id: 1},
//			want:    dom.Airline{Email: "raed786@gmail.com", AirlineName: "RAED", CompanyAddress: "RAED", PhoneNumber: "7902498141"},
//			wantErr: false,
//		},
//	}
//
//	beforeFunc(mockRepo)
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := svc.AdminVerifyAirlineRequest(tt.args.id)
//
//			if (err != nil) != tt.wantErr {
//				t.Errorf("svc.AdminVerifyAirlineRequest() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//
//			if got.Email != tt.want.Email || got.AirlineName != tt.want.AirlineName || got.CompanyAddress != tt.want.CompanyAddress || got.PhoneNumber != tt.want.PhoneNumber {
//				t.Errorf("svc.AdminVerifyAirlineRequest() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
