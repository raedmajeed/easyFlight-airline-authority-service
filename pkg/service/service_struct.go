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
	kfk   *config.KafkaReadWrite
}

func NewAdminAirlineService(repo inter.AdminAirlineRepostory, redis *redis.Client,
	cfg *config.ConfigParams, kfk *config.KafkaReadWrite) interfaces.AdminAirlineService {

	return &AdminAirlineServiceStruct{
		repo:  repo,
		redis: redis,
		cfg:   cfg,
		kfk:   kfk,
	}
}
