package agify

import (
	"api/internal/helpers"
	"github.com/goccy/go-json"
	"errors"
	"context"
)

type AgifyResponse struct {
	Count int		`json:"count"`
	Age	  int		`json:"age"`
	Name  string 	`json:"name"`
}

type Agify struct {
	URL		string
}

func New(url string) *Agify {
	return &Agify {
		URL:	url,
	}
}

func (a *Agify) MakeRequest(ctx context.Context, name string) (int,error){
	httpSender := httpsender.New(a.URL)
	params := map[string]string{
		"name": name,
	}
	if err := httpSender.SendRequestWithParams(ctx,params); err != nil {
		return 0,err
	}
	if httpSender.ResCode < 200 || httpSender.ResCode >= 400 {
		return 0,errors.New(httpSender.ResBody)
	}
	response := AgifyResponse{}
    json.Unmarshal([]byte(httpSender.ResBody), &response)
	return response.Age,nil
}
