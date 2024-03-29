package service

import (
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"api/internal/models"
	"api/internal/store"
	"context"
	"errors"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type PersonService struct {
	Store          store.IStore
	AgifyAPI       *agify.AgifyApi
	GenderizeAPI   *genderize.GenderizeAPI
	NationalizeAPI *nationalize.NationalizeAPI
	ctx            context.Context
}

func NewPerson(
	ctx context.Context,
	store store.IStore,
	agifyURL *agify.AgifyApi,
	genderURL *genderize.GenderizeAPI,
	nationalityURL *nationalize.NationalizeAPI) *PersonService {
	return &PersonService{
		Store:          store,
		AgifyAPI:       agifyURL,
		GenderizeAPI:   genderURL,
		NationalizeAPI: nationalityURL,
		ctx:            ctx,
	}
}

func (ps *PersonService) ActuatePerson(person *models.Person) error {
	ageCh := make(chan int)
	nationalityCh := make(chan string)
	genderCh := make(chan string)
	name := person.Name
	var age int
	var gender, nationality string
	chCnt := 0
	ctx, cancel := context.WithTimeout(ps.ctx, time.Second*5)

	go func() {
		age, err := ps.AgifyAPI.GetAgeByName(ctx, name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		ageCh <- age
	}()

	go func() {
		gender, err := ps.GenderizeAPI.GetGenderByName(ctx, name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		genderCh <- gender
	}()
	go func() {
		nationality, err := ps.NationalizeAPI.GetNationalityByName(ctx, name)
		if err != nil {
			log.Error(err)
			cancel()
		}
		nationalityCh <- nationality
	}()
loop:
	for {
		select {
		case age = <-ageCh:
			log.Info("age: ", age)
			chCnt++
		case gender = <-genderCh:
			log.Info("gender: ", gender)
			chCnt++
		case nationality = <-nationalityCh:
			log.Info("nationality: ", nationality)
			chCnt++
		case <-ctx.Done():
			return errors.New("Context canceled")
		default:
			if chCnt == 3 {
				break loop
			}
		}
	}
	person.Age = age
	person.Gender = gender
	person.Nationality = nationality

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
	if err := ps.Store.CreatePerson(ps.ctx, person); err != nil {
		return err
	}
	return nil
}

func (ps *PersonService) GetPerson(id string) (*models.Person, error) {

	person, err := ps.Store.GetPersonByID(ps.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, errors.New("Internal error")
	}
	return person, nil
}

func (ps *PersonService) DeletePerson(id string) error {
	if err := ps.Store.DeletePersonByID(ps.ctx, id); err != nil {
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
	if err := ps.Store.UpdatePerson(ps.ctx, updated_person); err != nil {
		return err
	}
	return nil
}
