package connection

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"tester/crosscutting/util"
	"tester/persistance/mongo/config"
)

type dbHandler struct {
	config config.Mongo
	client *mongo.Client
}

func NewDBHandler(config config.Mongo) *dbHandler {
	return &dbHandler{
		config: config,
	}
}

func (dh *dbHandler) GetDB() (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dh.config.Timeout)
	defer cancel()

	if dh.client == nil {
		opts := options.Client()
		opts.ApplyURI(dh.config.Server())
		opts.ConnectTimeout = util.AnyToPtr(dh.config.Timeout)
		client, err := mongo.Connect(opts)
		if err != nil {
			return nil, err
		}
		dh.client = client
	}

	if err := dh.client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return dh.client.Database(dh.config.DataStore), nil
}

func (dh *dbHandler) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), dh.config.Timeout)
	defer cancel()

	if dh.client != nil {
		if err := dh.client.Disconnect(ctx); err != nil {
			return err
		}
		dh.client = nil
	}

	return nil
}
