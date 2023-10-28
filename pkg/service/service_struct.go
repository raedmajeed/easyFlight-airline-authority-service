package service

import (
	"github.com/go-redis/redis/v8"
	inter "github.com/raedmajeed/admin-servcie/pkg/repository/interfaces"
	"github.com/raedmajeed/admin-servcie/pkg/service/interfaces"
)

type AdminAirlineServiceStruct struct {
	repo  inter.AdminAirlineRepostory
	redis *redis.Client
}

func NewAdminAirlineService(repo inter.AdminAirlineRepostory, redis *redis.Client) interfaces.AdminAirlineService {
	return &AdminAirlineServiceStruct{
		repo:  repo,
		redis: redis,
	}
}
