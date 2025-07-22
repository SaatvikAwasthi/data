package contract

import (
	"context"

	"tester/domain"
	"tester/persistance/mongo/dao"
)

//go:generate mockgen -package=mockContract -source=contract.go -destination=../../../mock/app/operation/contract/contract.go Contract
type RawDataRepository interface {
	Add(context.Context, dao.RawData) error
}

type RawDataProvider interface {
	GetAll(ctx context.Context) ([]dao.RawData, error)
	GetByID(ctx context.Context, id string) (*dao.RawData, error)
	GetByTime(ctx context.Context, createdAt string) (*dao.RawData, error)
}

type PostsServiceProvider interface {
	Fetch(ctx context.Context) (domain.Posts, error)
}
