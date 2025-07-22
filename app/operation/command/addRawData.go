package command

import (
	"context"
	"time"

	"tester/app/operation"
	"tester/app/operation/contract"
	"tester/persistance/mongo/dao"
)

type AddRawDataCommand struct {
	repository      contract.RawDataRepository
	serviceProvider contract.PostsServiceProvider
}

func NewAddRawDataCommand(repo contract.RawDataRepository, srvPrv contract.PostsServiceProvider) AddRawDataCommand {
	return AddRawDataCommand{
		repository:      repo,
		serviceProvider: srvPrv,
	}
}

func (arc AddRawDataCommand) Handle(ctx context.Context) (operation.AddRawDataResponse, error) {
	var rawPosts []dao.Post
	posts, err := arc.serviceProvider.Fetch(ctx)
	if err != nil {
		return operation.AddRawDataResponse{}, err
	}

	for _, post := range posts {
		rawPosts = append(rawPosts, dao.Post{
			UserId: post.UserId,
			Body:   post.Body,
			Title:  post.Title,
		})
	}
	rawData := dao.RawData{
		Source:    "https://jsonplaceholder.typicode.com/posts",
		CreatedAt: time.Now().UTC().String(),
		Posts:     rawPosts,
	}

	if err = arc.repository.Add(ctx, rawData); err != nil {
		return operation.AddRawDataResponse{}, err
	}

	return operation.AddRawDataResponse{Message: "Successfully fetched and stored posts in db"}, nil
}
