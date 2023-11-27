package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/admin-servcie/config"
	api "github.com/raedmajeed/admin-servcie/pkg/api"
	pkg "github.com/raedmajeed/admin-servcie/pkg/api"
	"github.com/raedmajeed/admin-servcie/pkg/api/handlers"
	"github.com/raedmajeed/admin-servcie/pkg/db"
	"github.com/raedmajeed/admin-servcie/pkg/repository"
	"github.com/raedmajeed/admin-servcie/pkg/service"
)

func InitApi(cfg *config.ConfigParams, redis *redis.Client) (*pkg.Server, error) {
	// db connection
	DB, err := db.NewDBConnect(cfg)
	if err != nil {
		return nil, err
	}
	kfWrite := config.NewKafkaWriterConnect()
	repo := repository.NewAdminAirlineRepository(DB)
	svc := service.NewAdminAirlineService(repo, redis, cfg, *kfWrite)
	hdl := handlers.NewAdminAirlineHandler(svc)
	server, err := api.NewServer(cfg, hdl, svc)
	if err != nil {
		return nil, err
	}
	return server, nil
}
