package store

import (
	"api/internal/models"
	"context"
)

type IStore interface {
	GetPersonByID(ctx context.Context, id string) (*models.Person, error)
	CreatePerson(ctx context.Context, person *models.Person) error
	UpdatePerson(ctx context.Context, person *models.Person) error
	DeletePersonByID(ctx context.Context, id string) error
	PersonSeeds(ctx context.Context) error
}
