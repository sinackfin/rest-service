package httpClient

import (
	"context"
)

type IHttpClient interface {
	GetWithParams(ctx context.Context, url string, params map[string]string) (*HTTPResponse, error)
}
