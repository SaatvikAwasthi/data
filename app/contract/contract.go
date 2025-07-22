package contract

import (
	"context"

	"tester/crosscutting/http"
)

//go:generate mockgen -package=mockContract -source=contract.go -destination=../../mock/app/contract/contract.go Contract
type HTTPClient interface {
	Get(ctx context.Context, path string, queryParams, headers map[string]string, result interface{}) http.Response
}
