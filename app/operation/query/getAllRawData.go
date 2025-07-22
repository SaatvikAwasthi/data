package query

import (
	"context"

	"tester/app/operation"
	"tester/app/operation/contract"
)

type GetAllRawData struct {
	provider contract.RawDataProvider
}

func NewGetAllRawDataQuery(prv contract.RawDataProvider) GetAllRawData {
	return GetAllRawData{
		provider: prv,
	}
}

func (query GetAllRawData) Handle(ctx context.Context) (operation.RawDataResponse, error) {
	var result []operation.RawData

	rawData, err := query.provider.GetAll(ctx)
	if err != nil {
		return operation.RawDataResponse{}, err
	}

	for _, data := range rawData {
		var posts []operation.Post
		for _, post := range data.Posts {
			posts = append(posts, operation.Post{
				UserId: post.UserId,
				Title:  post.Title,
				Body:   post.Body,
			})
		}

		result = append(result, operation.RawData{
			ID:        data.ID.String(),
			Source:    data.Source,
			CreatedAt: data.CreatedAt,
			Posts:     posts,
		})
	}

	return operation.RawDataResponse{Data: result}, nil
}
