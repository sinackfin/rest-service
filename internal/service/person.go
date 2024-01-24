package service

import (
	"api/internal/store"
	"api/internal/config"
	"api/internal/models"
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"time"
	log "github.com/sirupsen/logrus" 
	"errors"
    "github.com/google/uuid"
)

type PersonService struct {
	Store 			store.IStore
	AgifySvc		*agify.Agify
	GenderizeSvc	*genderize.Genderize
	NationalizeSvc	*nationalize.Nationalize
}

func New(cfg *config.ServiceConfig) *PersonService{
	return &PersonService{
		Store: 			cfg.Store,
		AgifySvc:		cfg.Agify,
		GenderizeSvc:	cfg.Genderize,
		NationalizeSvc:	cfg.Nationalize,
	}
}


func (ps *PersonService) ActuatePerson(person *models.Person) error {
	ageCh 			:= make(chan int)
	nationalityCh 	:= make(chan string)
	genderCh 		:= make(chan string)
	name 			:= person.Name
	var age int
	var gender,nationality string

	chCnt := 0
	t := time.NewTimer(5 * time.Second)

	go func(){
		age,_ := ps.AgifySvc.MakeRequest(name)
		ageCh <- age
	}()

	go func(){
		gender,_ := ps.GenderizeSvc.MakeRequest(name)
		genderCh <- gender
	}()
	go func(){
		nationality,_ := ps.NationalizeSvc.MakeRequest(name)
		nationalityCh <- nationality
	}()
loop:
	for {
		select {
			case age = <- ageCh:
				log.Info(age)
				chCnt++
			case gender = <- genderCh:
				log.Info(gender)
				chCnt++
			case nationality = <- nationalityCh:
				log.Info(nationality)
				chCnt++
			case <- t.C:
				return errors.New("Timeout")
			default:
				if chCnt == 3{
					break loop
				}
		}
	}
	person.Age 			= age
	person.Gender 		= gender
	person.Nationality 	= nationality

	return nil
}


func (ps *PersonService) CreatePerson(person *models.Person) error {
	if err := ps.ActuatePerson(person); err != nil {
		return err
	}
	person.ID = uuid.New().String()
	if err := person.Validate(); err != nil {
		return err
	}
	if err := ps.Store.CreatePerson(person); err != nil {
		return err
	}
	return nil
}



func (ps *PersonService) GetPerson(id string) (*models.Person,error) {

	person,err := ps.Store.GetPersonByID(id)
	if err != nil {
		log.Error(err)
		return nil,errors.New("Internal error")
	}
	return person,nil
}

func (ps *PersonService) DeletePerson(id string) error {
	if err := ps.Store.DeletePersonByID(id); err != nil {
		log.Error(err)
		return errors.New("Internal error")
	}
	return nil
}

func (ps *PersonService) UpdatePerson(updated_person *models.Person) error {
	if err := ps.ActuatePerson(updated_person); err != nil {
		return err
	}
	if err := updated_person.Validate(); err != nil {
		return err
	}
	if err := ps.Store.UpdatePerson(updated_person); err != nil {
		return err
	}
	return nil
}