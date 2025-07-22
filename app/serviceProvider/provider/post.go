package provider

import (
	"context"
	"errors"

	"tester/app/contract"
	"tester/app/serviceProvider/config"
	"tester/crosscutting/http"
	httpConfig "tester/crosscutting/http/config"
	"tester/crosscutting/util"
	"tester/domain"
)

const (
	postsEndpoint = "/posts"
)

type PostsServiceProvider struct {
	httpClient contract.HTTPClient
}

func New(conf config.Provider, httpConfig httpConfig.HTTPConfig) *PostsServiceProvider {
	return &PostsServiceProvider{
		httpClient: http.NewClient(httpConfig.WithBaseURL(conf.BaseUrl)),
	}
}

func (dsp *PostsServiceProvider) Fetch(ctx context.Context) (domain.Posts, error) {
	var response domain.Posts
	res := dsp.httpClient.Get(ctx, postsEndpoint, nil, nil, &response)
	if res.Err != nil {
		return nil, res.Err
	}
	if res.HTTPCode != 200 {
		return nil, util.NewError("failed to fetch data %d", res.HTTPCode)
	}
	if response == nil || len(response) == 0 {
		return nil, errors.New("no data found")
	}
	return response, nil
}
