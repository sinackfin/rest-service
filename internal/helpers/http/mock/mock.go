package httpClientMock

import (
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"api/internal/helpers/http"
	"errors"
	"github.com/goccy/go-json"
	"context"
)

type HTTPClientMock struct{}

func New() *HTTPClientMock {
	return &HTTPClientMock{}
}

func (h *HTTPClientMock) GetWithParams(ctx context.Context, url string, params map[string]string) (*httpClient.HTTPResponse, error) {
	reponse := &httpClient.HTTPResponse{}
	if params["name"] != "test" {
		return nil,errors.New("Not found")
	}
	switch url {
	case "agifyTest":
		agifyResp := agify.AgifyResponse{
			1,
			100,
			"test",
		}
		marshalBody, err := json.Marshal(agifyResp)
		if err != nil {
			return nil, err
		}
		reponse.ResCode = 200
		reponse.ResBody = string(marshalBody)
	case "genderizeTest":
		genderizeResp := genderize.GenderizeResponse{
			1,
			1,
			"male",
			"test",
		}
		marshalBody, err := json.Marshal(genderizeResp)
		if err != nil {
			return nil, err
		}
		reponse.ResCode = 200
		reponse.ResBody = string(marshalBody)
	case "nationalizeTest":
		nationalizeResp := nationalize.NationalizeResponse{
			2,
			"test",
			[]nationalize.CountryProbability{
				{
					ID: "RU",
					Probability: 0.2,
				},
				{
					ID: "US",
					Probability: 0.1,
				},
			},
		}
		marshalBody, err := json.Marshal(nationalizeResp)
		if err != nil {
			return nil, err
		}
		reponse.ResCode = 200
		reponse.ResBody = string(marshalBody)
	default:
		return nil, errors.New("Error API URL")
	}

	return reponse, nil
}
