package query

import (
	"context"

	"tester/app/operation"
	"tester/app/operation/contract"
	"tester/crosscutting/util"
)

type GetRawData struct {
	provider contract.RawDataProvider
}

func NewGetRawDataQuery(prv contract.RawDataProvider) GetRawData {
	return GetRawData{
		provider: prv,
	}
}

func (gq GetRawData) Handle(ctx context.Context, query GetRawDataQueryRequest) (operation.RawDataResponse, error) {
	switch query.Key {
	case Id:
		return gq.getById(ctx, query.Value)
	case CreatedAt:
		return gq.getByTime(ctx, query.Value)
	}

	return operation.RawDataResponse{}, util.Error("unknown query key: %d", query.Key)
}

func (gq GetRawData) getById(ctx context.Context, id string) (operation.RawDataResponse, error) {
	var rawData operation.RawData
	data, err := gq.provider.GetByID(ctx, id)
	if err != nil {
		return operation.RawDataResponse{}, util.Error("failed to get raw data by id: %w", err)
	}

	if data == nil {
		return operation.RawDataResponse{}, util.Error("raw data not found for id: %s", id)
	}

	rawData.ID = data.ID.String()
	rawData.Source = data.Source
	rawData.CreatedAt = data.CreatedAt
	for _, post := range data.Posts {
		rawData.Posts = append(rawData.Posts, operation.Post{
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		})
	}

	return operation.RawDataResponse{Data: []operation.RawData{rawData}}, nil
}

func (gq GetRawData) getByTime(ctx context.Context, time string) (operation.RawDataResponse, error) {
	var rawData operation.RawData
	data, err := gq.provider.GetByTime(ctx, time)
	if err != nil {
		return operation.RawDataResponse{}, util.Error("failed to get raw data by time: %w", err)
	}

	if data == nil {
		return operation.RawDataResponse{}, util.Error("raw data not found for time: %s", time)
	}

	rawData.ID = data.ID.String()
	rawData.Source = data.Source
	rawData.CreatedAt = data.CreatedAt
	for _, post := range data.Posts {
		rawData.Posts = append(rawData.Posts, operation.Post{
			UserId: post.UserId,
			Title:  post.Title,
			Body:   post.Body,
		})
	}

	return operation.RawDataResponse{Data: []operation.RawData{rawData}}, nil
}
