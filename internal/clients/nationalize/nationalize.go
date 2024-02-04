package nationalize

import (
	"api/internal/helpers/http"
	"context"
	"errors"
	"github.com/goccy/go-json"
)

type CountryProbability struct {
	ID          string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type NationalizeResponse struct {
	Count   int                  `json:"count"`
	Name    string               `json:"name"`
	Country []CountryProbability `json:"country"`
}

type NationalizeAPI struct {
	URL        string
	httpClient httpClient.IHttpClient
}

func New(url string, httpClient httpClient.IHttpClient) *NationalizeAPI {
	return &NationalizeAPI{
		url,
		httpClient,
	}
}

func (n *NationalizeAPI) GetNationalityByName(ctx context.Context, name string) (string, error) {
	params := map[string]string{
		"name": name,
	}
	resp, err := n.httpClient.GetWithParams(ctx, n.URL, params)
	if err != nil {
		return "", err
	}
	if resp.ResCode < 200 || resp.ResCode >= 400 {
		return "", errors.New(resp.ResBody)
	}
	response := NationalizeResponse{}
	json.Unmarshal([]byte(resp.ResBody), &response)
	nationality := response.FindMaxProbablity()
	return nationality, nil
}

// В целях упрощения, допущена неточность сравнения вещественных чисел

func (nr *NationalizeResponse) FindMaxProbablity() string {
	_max_probability := 0.0
	nationality := ""
	for _, v := range nr.Country {
		if _max_probability < v.Probability {
			_max_probability = v.Probability
			nationality = v.ID
		}
	}
	return nationality
}
