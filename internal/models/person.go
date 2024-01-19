package models

import (
	"errors"
	"regexp"
)

type Person struct {
	Age 		int		`json:"age"`
	ID			string	`json:"ID"`	
	Nationality	string	`json:"nationality"`
	Gender 		string	`json:"gender"`
	Name 		string	`json:"name"`
	Surname 	string	`json:"surname"`
	Patronymic	string	`json:"patronymic"`
}

func (p *Person) Validate() error {
	found, _ := regexp.MatchString("^[a-zA-Z]+$",p.Name)
	if !found {
		return errors.New("Set correct name")
	}
	if p.Age <= 0 {
		return errors.New("Set correct Age")
	} 
	if p.Gender == ""  {
		return errors.New("Set correct Gender")
	}
	if p.Nationality == ""  {
		return errors.New("Set correct Nationality")
	} 
	return nil
} 