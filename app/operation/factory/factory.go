package factory

import (
	"tester/app/operation/command"
	query2 "tester/app/operation/query"
	"tester/app/serviceProvider/factory"
	"tester/persistance/mongo/repository"
)

type CommandHandlers uint

const (
	AddRawDataCommandHandler CommandHandlers = iota + 1
)

type QueryHandlers uint

const (
	GetAllRawDataQueryHandler QueryHandlers = iota + 1
	GetRawDataQueryHandler
)

type OperationFactory struct {
	addRawDataCommand  command.AddRawDataCommand
	getAllRawDataQuery query2.GetAllRawData
	getRawDataQuery    query2.GetRawData
}

func Initialize(repositories *repository.Repositories, serviceProvider *factory.ServiceProvider) OperationFactory {
	return OperationFactory{
		addRawDataCommand:  command.NewAddRawDataCommand(repositories.RawData, serviceProvider.Post),
		getAllRawDataQuery: query2.NewGetAllRawDataQuery(repositories.RawData),
		getRawDataQuery:    query2.NewGetRawDataQuery(repositories.RawData),
	}
}

func (factory OperationFactory) CommandHandler(handler CommandHandlers) interface{} {
	switch handler {
	case AddRawDataCommandHandler:
		return factory.addRawDataCommand
	}
	return nil
}

func (factory OperationFactory) QueryHandler(handler QueryHandlers) interface{} {
	switch handler {
	case GetAllRawDataQueryHandler:
		return factory.getAllRawDataQuery
	case GetRawDataQueryHandler:
		return factory.getRawDataQuery
	}
	return nil
}
