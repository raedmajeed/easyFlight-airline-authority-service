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
)

func InitApi(cfg *config.ConfigParams, redis *redis.Client) {

	DB, _ := db.NewDBConnect(cfg)
	kfWrite := config.NewKafkaWriterConnect(cfg)
	repo := repository.NewAdminAirlineRepository(DB)
	svc := service.NewAdminAirlineService(repo, redis, cfg, *kfWrite)
	hdl := handlers.NewAdminAirlineHandler(svc)
	bhdl := bookingHandlers.NewBookingHandler(svc)
	go DailyFlightUpdate(svc)
	api.NewServer(cfg, hdl, svc, bhdl)
}
