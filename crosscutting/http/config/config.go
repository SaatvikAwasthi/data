package config

import (
	"net/http"
)

type HTTPConfig struct {
	BaseURL             string
	TimeoutInSeconds    int
	RetryCount          int
	RetryableErrorCodes []int
}

func NewHTTPConfig() *HTTPConfig {
	return &HTTPConfig{}
}

func (c HTTPConfig) Default() HTTPConfig {
	return HTTPConfig{
		TimeoutInSeconds: 30,
		RetryCount:       3,
		RetryableErrorCodes: []int{
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
			http.StatusRequestTimeout,
			http.StatusTooManyRequests,
		},
	}
}

func (c HTTPConfig) WithBaseURL(baseURL string) HTTPConfig {
	c.BaseURL = baseURL
	return c
}
