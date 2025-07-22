package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	appConfig "tester/app/api/config"
	serviceProvideConfig "tester/app/serviceProvider/config"
	mongoConfig "tester/persistance/mongo/config"
)

type ServerConfig struct {
	App          appConfig.App                 `envconfig:"APP"`
	PostProvider serviceProvideConfig.Provider `envconfig:"POST_PROVIDER"`
	Mongo        mongoConfig.Mongo             `envconfig:"MONGO"`
}

func NewServerConfig() ServerConfig {
	cfg := ServerConfig{}
	if err := envconfig.Process("TEST", &cfg); err != nil {
		log.Fatal(err.Error())
	}
	return cfg
}
