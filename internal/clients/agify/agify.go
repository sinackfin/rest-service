package agify

import (
	"api/internal/helpers/http"
	"context"
	"errors"
	"github.com/goccy/go-json"
)

type AgifyApi struct {
	URL        string
	httpClient httpClient.IHttpClient
}

type AgifyResponse struct {
	Count int    `json:"count"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
}

func New(url string, httpClient httpClient.IHttpClient) *AgifyApi {
	return &AgifyApi{
		url,
		httpClient,
	}
}

func (a *AgifyApi) GetAgeByName(ctx context.Context, name string) (int, error) {
	params := map[string]string{
		"name": name,
	}
	resp, err := a.httpClient.GetWithParams(ctx, a.URL, params)
	if err != nil {
		return 0, err
	}
	if resp.ResCode < 200 || resp.ResCode >= 400 {
		return 0, errors.New(resp.ResBody)
	}
	response := AgifyResponse{}
	json.Unmarshal([]byte(resp.ResBody), &response)
	return response.Age, nil
}
