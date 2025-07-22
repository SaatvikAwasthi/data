package db

import (
	"context"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"

	"tester/persistance/mongo/config"
)

type MongoDBContainer *mongodb.MongoDBContainer

func InitMongoDBContainer(cfg config.Mongo, ctx context.Context) (MongoDBContainer, error) {
	return mongodb.Run(ctx, "mongo:6")
}
