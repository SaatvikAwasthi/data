package server

import (
	"github.com/gin-gonic/gin"

	"tester/app/api/contract"
	"tester/app/api/router"
	"tester/app/server/config"
	"tester/crosscutting/api/middleware/cors"
	"tester/crosscutting/api/middleware/requestLogger"
)

func setUpApp(cfg config.ServerConfig, factory contract.OperationHandlerFactory) *gin.Engine {
	gin.SetMode(cfg.App.GinMode)
	engine := gin.New()

	engine.
		Use(gin.Recovery()).
		Use(cors.Handle(cfg.App)).
		Use(requestLogger.Handler())

	return router.RegisterRoutes(factory, engine)
}
