package config

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
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
	DOCKERVERSION    string `mapstructure:"DOCKERVERSION"`
}

func Configuration() (*ConfigParams, error, *redis.Client) {
	cfg := ConfigParams{}
	//if err := godotenv.Load("../../.env"); err != nil {
	//	os.Exit(1)
	//}

	cfg.DBHost = os.Getenv("DBHOST")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.DBUser = os.Getenv("DBUSER")
	cfg.DBPort = os.Getenv("DBPORT")
	cfg.DBPassword = os.Getenv("DBPASSWORD")
	cfg.PORT = os.Getenv("PORT")
	cfg.ADMINPORT = os.Getenv("ADMINPORT")
	cfg.REDISHOST = os.Getenv("REDISHOST")
	cfg.SECRETKEY = os.Getenv("SECRETKEY")
	cfg.BUSINESSSURGE = os.Getenv("BUSINESSSURGE")
	cfg.ADMINBOOKINGPORT = os.Getenv("ADMINBOOKINGPORT")
	cfg.KAFKABROKER = os.Getenv("KAFKABROKER")
	cfg.DOCKERVERSION = os.Getenv("DOCKERVERSION")

	if cfg.DOCKERVERSION == "" {
		log.Fatal("error configuring env values: ", cfg)
	}
	redis2 := connectToRedis(&cfg)
	return &cfg, nil, redis2
}

func connectToRedis(cfg *ConfigParams) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.REDISHOST,
		Password: "",
		DB:       2,
	})
	return client
}

func GetDockerTag() string {
	apiUrl := "http://localhost/info"
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatal("Failed to fetch Docker info:", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body:", err)
	}
	return string(body)
}
