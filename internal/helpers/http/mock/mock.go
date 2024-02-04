package httpClientMock

import (
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"api/internal/helpers/http"
	"errors"
	"github.com/goccy/go-json"
)

type HTTPClientMock struct{}

func New(url string) *HTTPClientMock {
	return &HTTPClientMock{}
}

func (h *HTTPClientMock) GetWithParams(ctx context.Context, params map[string]string) (*httpClient.HTTPResponse, error) {
	if !ok {
		return nil, errors.New("Not define name field")
	}
	reponse := &httpClient.HTTPResponse{}
	switch h.baseURL {
	case "agifyTest":
		agifyResp := agify.AgifyResponse{
			1,
			100,
			"test",
		}
		reponse.ResCode = 200
		reponse.ResBody = json.Marshal(agifyResp)
	case "genderizeTest":
		genderizeResp := genderize.GenderizeResponse{
			1,
			1,
			"male",
			"test",
		}
		reponse.ResCode = 200
		reponse.ResBody = json.Marshal(genderizeResp)
	case "nationalizeTest":
		nationalizeResp := nationalize.NationalizeResponse{
			2,
			"test",
			[]nationalize.CountryProbability{
				{
					ID: "RU",
					0.2,
				},
				{
					ID: "US",
					0.1,
				},
			},
		}
		reponse.ResCode = 200
		reponse.ResBody = json.Marshal(nationalizeResp)
	default:
		return nil, errors.New("Error API URL")
	}

	return reponse, nil
}
