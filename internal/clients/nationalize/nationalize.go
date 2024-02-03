package nationalize

import (
	"api/internal/helpers"
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

type Nationalize struct {
	URL string
}

func New(url string) *Nationalize {
	return &Nationalize{
		URL: url,
	}
}

func (n *Nationalize) MakeRequest(ctx context.Context, name string) (string, error) {
	httpSender := httpsender.New(n.URL)
	params := map[string]string{
		"name": name,
	}
	if err := httpSender.SendRequestWithParams(ctx, params); err != nil {
		return "", err
	}
	if httpSender.ResCode < 200 || httpSender.ResCode >= 400 {
		return "", errors.New(httpSender.ResBody)
	}
	response := NationalizeResponse{}
	json.Unmarshal([]byte(httpSender.ResBody), &response)
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
