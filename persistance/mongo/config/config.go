package config

import (
	"os"
	"time"

	"tester/crosscutting/util"
)

type Mongo struct {
	DataStore string        `envconfig:"DATASTORE" required:"true"`
	Username  string        `envconfig:"USERNAME" required:"true"`
	Password  string        `envconfig:"PASSWORD" required:"true"`
	Host      string        `envconfig:"HOST" default:"localhost"`
	Port      uint          `envconfig:"PORT" default:"27017"`
	ServerURI string        `envconfig:"SERVER_URI"`
	Timeout   time.Duration `envconfig:"TIMEOUT" default:"20s"`
}

func (m Mongo) Server() string {
	if os.Getenv("TEST_APP_MODE") != "PROD" {
		format := "mongodb://%s:%s@%s:%d/%s?authSource=admin"
		return util.Format(format, m.Username, m.Password, m.Host, m.Port, m.DataStore)
	}
	format := "mongodb+srv://%s:%s@%s"
	return util.Format(format, m.Username, m.Password, m.ServerURI)
}
