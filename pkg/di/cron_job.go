package di

import (
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"log"
	"time"
)

func DailyFlightUpdate(svc interfaces.AdminAirlineService) {
	for {
		select {
		case <-time.After(time.Hour * 24):
			day := time.Now().Weekday()
			log.Println("calculating fare for today: ", day)
			svc.CalculateDailyFare()
		}
	}
}
