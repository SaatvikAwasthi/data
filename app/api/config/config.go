package config

import (
	"tester/crosscutting/util"
)

type App struct {
	Host           string   `envconfig:"HOST" required:"true"`
	Port           string   `envconfig:"PORT" required:"true"`
	GinMode        string   `envconfig:"GIN_MODE" default:"debug"`
	AllowedOrigins []string `envconfig:"ALLOWED_ORIGINS" default:"*"`
}

func (ac App) GetAddress() string {
	return util.Format("%s:%s", ac.Host, ac.Port)
}
