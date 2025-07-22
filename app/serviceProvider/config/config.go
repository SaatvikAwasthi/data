package config

type Provider struct {
	BaseUrl string `json:"base_url,required" envconfig:"BASE_URL" required:"true"`
}
