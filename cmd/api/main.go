package main

import (
	"log"

	"github.com/raedmajeed/admin-servcie/config"
	"github.com/raedmajeed/admin-servcie/pkg/di"
)

func main() {
	cfg, err, redis := config.Configuration()
	kafkaMethods := config.NewKafkaConnect()
	if err != nil {
		log.Printf("unable to load env values, err: %v", err.Error())
		return
	}
	server, err := di.InitApi(cfg, redis, kafkaMethods)
	if err != nil {
		log.Fatalf("Server not starter due to error: %v", err.Error())
		return
	}
	if err = server.ServerStart(); err != nil {
		log.Fatalf("Server not starter due to error: %v", err.Error())
	}
}
