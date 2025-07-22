package db

import (
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"

	"tester/persistance/mongo/config"
	"tester/persistance/mongo/connection"
)

type testDB struct {
	db     *mongo.Database
	client *mongo.Client
	isTest bool
}

var once sync.Once
var instance *testDB

func InitDB(c config.Mongo) {
	once.Do(func() {
		db, err := connection.NewDBHandler(c).GetDB()
		if err != nil {
			panic(err)
		}
		instance = &testDB{db: db}
	})
}

func GetDB() *mongo.Database {
	return instance.db
}
