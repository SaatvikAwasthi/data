package cors

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"tester/app/api/config"
	"tester/crosscutting/constants"
)

func Handle(cfg config.App) gin.HandlerFunc {
	conf := cors.DefaultConfig()
	if cfg.GinMode == gin.DebugMode {
		conf.AllowAllOrigins = true
	} else {
		conf.AllowOrigins = cfg.AllowedOrigins
	}
	conf.AllowCredentials = true
	conf.AllowHeaders = []string{
		constants.Origin,
		constants.Accept,
		constants.ContentTypeHeader,
		constants.Authorization,
		constants.DateUsed,
		constants.XRequestedWith,
	}
	conf.AllowMethods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodOptions,
	}

	return cors.New(conf)
}
