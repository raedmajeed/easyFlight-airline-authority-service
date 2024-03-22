package main

import (
	"log"

	"github.com/raedmajeed/admin-servcie/config"
	"github.com/raedmajeed/admin-servcie/pkg/di"
)

func main() {
	version := config.GetDockerTag()
	log.Println("connecting to docker image version: ", version)
	cfg, err, redis := config.Configuration()
	if err != nil {
		log.Printf("unable to load env values, err: %v", err.Error())
		return
	}
	di.InitApi(cfg, redis)
}
