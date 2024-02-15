package utils

import (
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
	"time"
)

func CronJob(svc interfaces.AdminAirlineService) {
	for {
		select {
		case <-time.After(time.Hour * 24):

		}
	}
}
