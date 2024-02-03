package service

import (
	"api/internal/store"
	"api/internal/models"
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"time"
	log "github.com/sirupsen/logrus" 
	"errors"
    "github.com/google/uuid"
	"context"
)

type PersonService struct {
	Store 			store.IStore
	AgifySvc		*agify.Agify
	GenderizeSvc	*genderize.Genderize
	NationalizeSvc	*nationalize.Nationalize
	ctx				context.Context
}

func New(store store.IStore, agify *agify.Agify, gender *genderize.Genderize, nationality *nationalize.Nationalize, ctx context.Context) *PersonService{
	return &PersonService{
		Store: 			store,
		AgifySvc:		agify,
		GenderizeSvc:	gender,
		NationalizeSvc:	nationality,
		ctx:			ctx,	
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
	ctx,cancel := context.WithTimeout(ps.ctx,time.Second * 5)

	go func(){
		age,err := ps.AgifySvc.MakeRequest(name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		ageCh <- age
	}()

	go func(){
		gender,err := ps.GenderizeSvc.MakeRequest(name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		genderCh <- gender
	}()
	go func(){
		nationality,err := ps.NationalizeSvc.MakeRequest(name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		nationalityCh <- nationality
	}()
loop:
	for {
		select {
			case age = <- ageCh:
				log.Info("age: ",age)
				chCnt++
			case gender = <- genderCh:
				log.Info("gender: ", gender)
				chCnt++
			case nationality = <- nationalityCh:
				log.Info("nationality: ", nationality)
				chCnt++
			case <- ctx.Done():
				return errors.New("Context canceled")
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
	if err := ps.Store.CreatePerson(ps.ctx,person); err != nil {
		return err
	}
	return nil
}



func (ps *PersonService) GetPerson(id string) (*models.Person,error) {

	person,err := ps.Store.GetPersonByID(ps.ctx,id)
	if err != nil {
		log.Error(err)
		return nil,errors.New("Internal error")
	}
	return person,nil
}

func (ps *PersonService) DeletePerson(id string) error {
	if err := ps.Store.DeletePersonByID(ps.ctx,id); err != nil {
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
	if err := ps.Store.UpdatePerson(ps.ctx,updated_person); err != nil {
		return err
	}
	return nil
}