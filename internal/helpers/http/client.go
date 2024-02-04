package httpClient

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
)

type HTTPResponse struct {
	ResCode int
	ResBody string
}

type HTTPClient struct{}

func New() *HTTPClient {
	return &HTTPClient{}
}

func (h *HTTPClient) GetWithParams(ctx context.Context, url string, queryParams map[string]string) (*HTTPResponse, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	for i, v := range queryParams {
		q.Add(i, v)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &HTTPResponse{
		resp.StatusCode,
		string(responseBody),
	}

	return response, nil
}
