package genderize

import (
	"api/internal/helpers"
	"context"
	"errors"
	"github.com/goccy/go-json"
)

type Genderize struct {
	URL string
}

type GenderizeResponse struct {
	Count       int    `json:"count"`
	Probability int    `json:"probability"`
	Gender      string `json:"gender"`
	Name        string `json:"name"`
}

func New(url string) *Genderize {
	return &Genderize{
		URL: url,
	}
}

func (g *Genderize) MakeRequest(ctx context.Context, name string) (string, error) {
	httpSender := httpsender.New(g.URL)
	params := map[string]string{
		"name": name,
	}
	if err := httpSender.SendRequestWithParams(ctx, params); err != nil {
		return "", err
	}
	if httpSender.ResCode < 200 || httpSender.ResCode >= 400 {
		return "", errors.New(httpSender.ResBody)
	}
	response := GenderizeResponse{}
	json.Unmarshal([]byte(httpSender.ResBody), &response)
	return response.Gender, nil
}
