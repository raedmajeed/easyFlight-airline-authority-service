package config

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

type ConfigParams struct {
	DBHost           string `mapstructure:"DBHOST"`
	DBName           string `mapstructure:"DBNAME"`
	DBUser           string `mapstructure:"DBUSER"`
	DBPort           string `mapstructure:"DBPORT"`
	DBPassword       string `mapstructure:"DBPASSWORD"`
	PORT             string `mapstructure:"PORT"`
	ADMINPORT        string `mapstructure:"ADMINPORT"`
	REDISHOST        string `mapstructure:"REDISHOST"`
	SECRETKEY        string `mapstructure:"SECRETKEY"`
	BUSINESSSURGE    string `mapstructure:"BUSINESSSURGE"`
	ADMINBOOKINGPORT string `mapstructure:"ADMINBOOKINGPORT"`
	KAFKABROKER      string `mapstructure:"KAFKABROKER"`
}

//var envs = []string{
//	"DBHOST", "DBNAME", "DBSUER", "DBPORT", "DBPASSWORD", "PORT", "ADMINPORT", "REDISHOST", "SECRETKEY", "BUSINESSSURGE", "ADMINBOOKINGPORT",
//}

func Configuration() (*ConfigParams, error, *redis.Client) {
	var cfg ConfigParams
	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Unable to load env values, err: %v", err.Error())
		return &ConfigParams{}, err, nil
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to unmarshal values, err: %v", err.Error())
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return &cfg, err, nil
	}

	redis2 := connectToRedis(&cfg)
	return &cfg, err, redis2
}

func connectToRedis(cfg *ConfigParams) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "",
		DB:       2,
	})
	return client
}
