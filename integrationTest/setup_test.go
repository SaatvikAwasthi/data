package integrationTest

import (
	"context"
	"log"
	"os"
	"testing"

	appConfig "tester/app/api/config"
	"tester/app/server/config"
	"tester/integrationTest/db"
	"tester/integrationTest/mockServer"
	mongoConfig "tester/persistance/mongo/config"
)

var TestServer *mockServer.MockServer
var Conf config.ServerConfig

func TestMain(m *testing.M) {
	ctx := context.Background()

	os.Setenv("DOCKER_HOST", "unix:///Users/saatvikAwasthi/.colima/docker.sock")

	Conf = config.ServerConfig{
		App: appConfig.App{
			Host:           "localhost",
			Port:           "13001",
			GinMode:        "debug",
			AllowedOrigins: []string{"*"},
		},
		Mongo: mongoConfig.Mongo{
			DataStore: "mydb",
			Username:  "admin	",
			Password:  "password",
			Host:      "localhost",
			Port:      27017,
			Timeout:   0,
		},
	}

	setUpDBContainer(ctx)
	TestServer = mockServer.NewMockServer().Init()
	Conf.PostProvider.BaseUrl = TestServer.GetURL()
	db.InitDB(Conf.Mongo)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func setUpDBContainer(ctx context.Context) db.MongoDBContainer {
	c, err := db.InitMongoDBContainer(Conf.Mongo, ctx)
	if err != nil {
		log.Fatalln("Failed to create container:", err)
		return c
	}
	return c
}
