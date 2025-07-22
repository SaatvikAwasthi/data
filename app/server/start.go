package server

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"tester/app/operation/command"
	operationFactory "tester/app/operation/factory"
	"tester/app/server/config"
	serviceProviderFactory "tester/app/serviceProvider/factory"
	httpConfig "tester/crosscutting/http/config"
	"tester/crosscutting/util"
	mongoRepo "tester/persistance/mongo/repository"
)

func Init() {
	srvCfg := config.NewServerConfig()
	_ = setUp(srvCfg).Run(srvCfg.App.GetAddress())
}

func setUp(cfg config.ServerConfig) *gin.Engine {
	defer util.RecoverPanic()
	clientConfig := httpConfig.NewHTTPConfig().Default()
	serviceProvider := serviceProviderFactory.Initialize(cfg.PostProvider, clientConfig)
	mongoRepository := mongoRepo.Initialize(cfg.Mongo)
	factory := operationFactory.Initialize(mongoRepository, serviceProvider)
	//fetchData(context.Background(), factory)
	return setUpApp(cfg, factory)
}

func fetchData(ctx context.Context, factory operationFactory.OperationFactory) {
	defer util.RecoverPanic()
	cmd := factory.CommandHandler(operationFactory.AddRawDataCommandHandler)
	resp, err := cmd.(command.AddRawDataCommand).Handle(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Command response: %v", resp)
}
