package contract

import (
	"context"

	"tester/app/operation"
	"tester/app/operation/factory"
	"tester/app/operation/query"
)

//go:generate mockgen -package=mockContract -source=contract.go -destination=../../../mock/app/api/contract/contract.go Contract
type OperationHandlerFactory interface {
	QueryHandler(handlers factory.QueryHandlers) interface{}
	CommandHandler(handlers factory.CommandHandlers) interface{}
}

type GetAllRawData interface {
	Handle(ctx context.Context) (operation.RawDataResponse, error)
}

type GetRawData interface {
	Handle(ctx context.Context, query query.GetRawDataQueryRequest) (operation.RawDataResponse, error)
}

type AddRawData interface {
	Handle(ctx context.Context) (operation.AddRawDataResponse, error)
}
