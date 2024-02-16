package repository_test

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	dom "github.com/raedmajeed/admin-servcie/pkg/DOM"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"regexp"
	"testing"
)

func TestFindAirportByAirportCode(t *testing.T) {
	type args struct {
		airportCode string
	}
	tests := []struct {
		name       string
		args       args
		beforeTest func(sqlMock sqlmock.Sqlmock)
		want       *dom.Airport
		wantErr    bool
	}{
		{
			name: "Test 1",
			args: args{
				airportCode: "DEL",
			},
			beforeTest: func(sqlMock sqlmock.Sqlmock) {
				sqlMock.ExpectBegin()
				sqlMock.ExpectQuery(regexp.QuoteMeta(`
						SELECT * FROM flight_booking_airline.airports WHERE airport_code = $1 LIMIT 1		
				`)).WithArgs("DEL").
					WillReturnError(errors.New("NO AIRPORT FOUND"))
				sqlMock.ExpectRollback()
			},
			want:    &dom.Airport{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()
			database, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
			r := repository.AdminAirlineRepositoryStruct{
				DB: database,
			}

			tt.beforeTest(mockSQL)
			got, err := r.FindAirportByAirportCode(tt.args.airportCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAirportByAirportCode(DeL) error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Create() = %v, want %v", got, tt.want)
			}
		})

	}
}
