package router

import (
	"github.com/gin-gonic/gin"

	//"tester/api/handler"

	"tester/app/api/contract"
	"tester/app/api/handler"
)

const (
	dataEndpoint = "data"
)

func RegisterRoutes(factory contract.OperationHandlerFactory, engine *gin.Engine) *gin.Engine {
	bankHandler := handler.NewDataHandler(factory)
	engine.GET(dataEndpoint, bankHandler.GetData)
	engine.POST(dataEndpoint, bankHandler.AddData)
	return engine
}
