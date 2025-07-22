package http

import (
	"context"
	"log"
	"slices"
	"time"

	"github.com/go-resty/resty/v2"

	"tester/crosscutting/constants"
	"tester/crosscutting/http/config"
)

type Response struct {
	HTTPCode int
	Err      error
}

type Client struct {
	client *resty.Client
}

func NewClient(conf config.HTTPConfig) Client {
	client := resty.New()
	client.SetBaseURL(conf.BaseURL)
	if conf.TimeoutInSeconds > 0 {
		client.SetTimeout(time.Duration(conf.TimeoutInSeconds) * time.Second)
	}
	if conf.RetryCount > 0 {
		client.SetRetryCount(conf.RetryCount)
	}
	if conf.RetryCount > 0 && len(conf.RetryableErrorCodes) > 0 {
		client.AddRetryCondition(
			func(r *resty.Response, err error) bool {
				isRetryableError := slices.Contains(conf.RetryableErrorCodes, r.StatusCode())
				if isRetryableError {
					log.Printf("Http is retrying as error code is received %d: \n", r.StatusCode())
				}
				return isRetryableError
			},
		)
	}

	return Client{
		client: client,
	}
}

func (c Client) Get(ctx context.Context, path string, queryParams, headers map[string]string, result interface{}) Response {
	req := c.client.R().ForceContentType(constants.ApplicationJson)

	if queryParams != nil {
		req.SetQueryParams(queryParams)
	}
	if headers != nil {
		req.SetHeaders(headers)
	}
	if result != nil {
		req.SetResult(result)
	}

	res, err := req.Get(path)
	if res != nil {
		return Response{
			HTTPCode: res.StatusCode(),
			Err:      err,
		}
	}

	return Response{Err: err}
}
