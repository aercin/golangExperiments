package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	MessageRelay  MessageRelay
	HttpServer    HttpServer
	Postgres      Postgres
	BasketService BasketService
	RabbitMQ      RabbitMQ
}

type HttpServer struct {
	Port int
}

type Postgres struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
}

type RabbitMQ struct {
	BrokerAddress  string
	ProduceTimeout int
	ProduceQueue   string
	ConsumeQueue   string
}

type BasketService struct {
	Address string
}

type MessageRelay struct {
	CycleTime int
}

func NewConfig() *Config {
	environment := os.Getenv("APP_ENV")

	if environment == "" {
		panic("APP_ENV environment variable is not set")
	}

	viper.SetConfigName(fmt.Sprintf("appsettings.%s", environment))
	viper.SetConfigType("json")
	viper.AddConfigPath("../../configs")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct: %v", err))
	}

	return &config
}
