package factory

import (
	"tester/app/serviceProvider/config"
	"tester/app/serviceProvider/provider"
	httpConfig "tester/crosscutting/http/config"
)

type ServiceProvider struct {
	Post *provider.PostsServiceProvider
}

func Initialize(postConfig config.Provider, httpConfig httpConfig.HTTPConfig) *ServiceProvider {
	return &ServiceProvider{
		Post: provider.New(postConfig, httpConfig),
	}
}
