package genderize

import (
	"api/internal/helpers/http"
	"context"
	"errors"
	"github.com/goccy/go-json"
)

type GenderizeAPI struct {
	URL        string
	httpClient httpClient.IHttpClient
}

type GenderizeResponse struct {
	Count       int    `json:"count"`
	Probability int    `json:"probability"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
}

func New(url string, httpClient httpClient.IHttpClient) *GenderizeAPI {
	return &GenderizeAPI{
		url,
		httpClient,
	}
}

func (g *GenderizeAPI) GetGenderByName(ctx context.Context, name string) (string, error) {
	params := map[string]string{
		"name": name,
	}
	resp, err := g.httpClient.GetWithParams(ctx, g.URL, params)
	if err != nil {
		return "", err
	}
	if resp.ResCode < 200 || resp.ResCode >= 400 {
		return "", errors.New(resp.ResBody)
	}
	response := GenderizeResponse{}
	json.Unmarshal([]byte(resp.ResBody), &response)
	return response.Gender, nil
}
