package agify

import (
	"api/internal/helpers"
	"github.com/goccy/go-json"
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

func (a *Agify) MakeRequest(name string) (int,error){
	httpSender := httpsender.New(a.URL)
	params := map[string]string{
		"name": name,
	}
	if err := httpSender.SendRequestWithParams(params); err != nil {
		return 0,err
	}
	response := AgifyResponse{}
    json.Unmarshal([]byte(httpSender.ResBody), &response)
	return response.Age,nil
}
