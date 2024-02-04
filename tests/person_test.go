package tests

import (
	"testing"
	"context"
	"api/internal/store/mock"
	"api/internal/service"
	"api/internal/helpers/http/mock"
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
	"api/internal/models"
)

func TestGetPersonByID(t *testing.T) {
	ctx := context.Background()
	store := mock.New(ctx)
	httpClient := httpClientMock.New()
	personService := service.NewPerson(
		ctx,
		store,
		agify.New("agifyTest", httpClient),
		genderize.New("genderizeTest", httpClient),
		nationalize.New("nationalizeTest", httpClient),
	)
	person := models.Person{
		Name: "test",
	}
	personService.CreatePerson(&person)
	addedPerson,err := personService.GetPerson(person.ID)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if person.Name != addedPerson.Name {
		t.Fatalf("Names not mathced!")
	}
}

func TestUpdatePerson(t *testing.T) {
	ctx := context.Background()
	store := mock.New(ctx)
	httpClient := httpClientMock.New()
	personService := service.NewPerson(
		ctx,
		store,
		agify.New("agifyTest", httpClient),
		genderize.New("genderizeTest", httpClient),
		nationalize.New("nationalizeTest", httpClient),
	)
	person := models.Person{
		Name: "test",
		Age: 200,
		Gender: "female",
		Nationality: "JA",
	}
	personService.CreatePerson(&person)
	err := personService.UpdatePerson(&person)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if person.Age != 100 {
		t.Fatalf("Age not mathced!")
	}
	if person.Gender != "male" {
		t.Fatalf("Gender not mathced!")
	}
	if person.Nationality != "RU" {
		t.Fatalf("Nationality not mathced!")
	}
}
