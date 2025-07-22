package repository

import (
	"tester/persistance/mongo/config"
	"tester/persistance/mongo/connection"
)

type Repositories struct {
	RawData *rawDataRepo
}

func Initialize(cfg config.Mongo) *Repositories {
	handler := connection.NewDBHandler(cfg)
	db, err := handler.GetDB()
	if err != nil {
		panic(err.Error())
	}

	return &Repositories{
		RawData: NewDataRepo(db),
	}
}
