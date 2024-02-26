package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/admin-servcie/config"
	api "github.com/raedmajeed/admin-servcie/pkg/api"
	"github.com/raedmajeed/admin-servcie/pkg/api/bookingHandlers"
	"github.com/raedmajeed/admin-servcie/pkg/api/handlers"
	"github.com/raedmajeed/admin-servcie/pkg/db"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
	"github.com/raedmajeed/admin-servcie/pkg/service"
	"log"
	"os"
)

func InitApi(cfg *config.ConfigParams, redis *redis.Client) {

	DB, err := db.NewDBConnect(cfg)
	if err != nil {
		os.Exit(2)
	}
	log.Println("here")
	kfWrite := config.NewKafkaWriterConnect(cfg)
	repo := repository.NewAdminAirlineRepository(DB)
	svc := service.NewAdminAirlineService(repo, redis, cfg, *kfWrite)
	hdl := handlers.NewAdminAirlineHandler(svc)
	bhdl := bookingHandlers.NewBookingHandler(svc)
	go DailyFlightUpdate(svc)
	api.NewServer(cfg, hdl, svc, bhdl)
}
