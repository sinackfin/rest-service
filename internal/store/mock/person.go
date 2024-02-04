package mock

import (
	"api/internal/models"
	"context"
)

type MockStore struct {
	persons map[string]*models.Person
}

func New(ctx context.Context) *MockStore {
	return &MockStore{
		make(map[string]*models.Person),
	}
}

func (ms *MockStore) CreatePerson(ctx context.Context, person *models.Person) error {
	ms.persons[person.ID] = person
	return nil
}

func (ms *MockStore) GetPersonByID(ctx context.Context, id string) (*models.Person, error) {
	return ms.persons[id], nil
}
func (ms *MockStore) UpdatePerson(ctx context.Context, person *models.Person) error {
	ms.persons[person.ID] = person
	return nil
}
func (ms *MockStore) DeletePersonByID(ctx context.Context, id string) error {
	delete(ms.persons, id)
	return nil
}

func (ms *MockStore) PersonSeeds(ctx context.Context) error {
	return nil
}
