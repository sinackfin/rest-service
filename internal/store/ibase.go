package store

import (
	"api/internal/models"
)

type IStore interface {
	GetPersonByID(id string) (*models.Person,error)
	CreatePerson(*models.Person)error
	UpdatePerson(*models.Person) error 
	DeletePersonByID(id string)error
}