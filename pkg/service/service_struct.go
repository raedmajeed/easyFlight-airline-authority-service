package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/raedmajeed/admin-servcie/config"
	inter "github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
)

type AdminAirlineServiceStruct struct {
	repo  inter.AdminAirlineRepostory
	redis *redis.Client
	cfg   *config.ConfigParams
}

func NewAdminAirlineService(repo inter.AdminAirlineRepostory, redis *redis.Client, cfg *config.ConfigParams) interfaces.AdminAirlineService {
	return &AdminAirlineServiceStruct{
		repo:  repo,
		redis: redis,
		cfg:   cfg,
	}
}
