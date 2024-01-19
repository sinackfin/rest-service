package httpsender

import (
	"io/ioutil"
	"net/http"
)

type HTTPSender struct {
	ResCode string
	ResBody string
	baseURL	string
}

func New(url string) *HTTPSender {
	return &HTTPSender{
		baseURL: url,
	}
}

func (h *HTTPSender)SendRequestWithParams(params map[string]string) error{
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, h.baseURL, nil)
	
	if err != nil {
		return err
	}

	q := req.URL.Query()
	
	for i,v := range params {
		q.Add(i,v)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	h.ResCode = resp.Status
	h.ResBody = string(responseBody)
	return nil
}