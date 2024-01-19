package genderize

import (
	"api/internal/helpers"
	"encoding/json"
)

type Genderize struct {
	URL		string
}

type GenderizeResponse struct {
	Count 			int		`json:"count"`
	Probability 	int		`json:"probability"`
	Gender			string	`json:"gender"`
	Name			string	`json:"name"`
}

func New(url string) *Genderize {
	return &Genderize {
		URL:	url,
	}
}

func (g *Genderize) MakeRequest(name string) (string,error){
	httpSender := httpsender.New(g.URL)
	params := map[string]string{
		"name": name,
	}
	if err := httpSender.SendRequestWithParams(params); err != nil {
		return "",err
	}
	response := GenderizeResponse{}
    json.Unmarshal([]byte(httpSender.ResBody), &response)
	return response.Gender,nil
}