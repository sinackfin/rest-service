package pg

import (
	"api/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

type PGStore struct {
	pgxpool *pgxpool.Pool
}

func New(ctx context.Context, dbConnStr string) (*PGStore, error) {
	pgxpool, err := pgxpool.New(ctx, dbConnStr)
	if err != nil {
		return nil, err
	}
	if err = pgxpool.Ping(ctx); err != nil {
		return nil, err
	}
	return &PGStore{
		pgxpool,
	}, nil
}

func (ps *PGStore) CreatePerson(ctx context.Context, person *models.Person) error {
	query := `INSERT INTO person (person_id,age,nationality,gender,name,surname,patronymic)
				VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := ps.pgxpool.Exec(ctx, query, person.ID, person.Age, person.Nationality,
		person.Gender, person.Name, person.Surname, person.Patronymic)
	return err
}

func (ps *PGStore) GetPersonByID(ctx context.Context, id string) (*models.Person, error) {
	query := `SELECT person_id,age,nationality,gender,name,surname,patronymic from person WHERE person_id=$1`
	person := models.Person{}
	err := ps.pgxpool.QueryRow(ctx, query, id).Scan(&person.ID, &person.Age,
		&person.Nationality, &person.Gender, &person.Name, &person.Surname, &person.Patronymic)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
func (ps *PGStore) UpdatePerson(ctx context.Context, person *models.Person) error {
	query := `UPDATE person SET age=$1,nationality=$2,gender=$3,name=$4,surname=$5,patronymic=$6 WHERE person_id=$7`
	_, err := ps.pgxpool.Exec(ctx, query, person.Age, person.Nationality,
		person.Gender, person.Name, person.Surname, person.Patronymic, person.ID)
	return err
}
func (ps *PGStore) DeletePersonByID(ctx context.Context, id string) error {
	query := `DELETE from person WHERE person_id=$1`
	_, err := ps.pgxpool.Exec(ctx, query, id)
	return err
}

func (ps *PGStore) PersonSeeds(ctx context.Context) error {
	const nSeeds = 100_000
	t1 := time.Now()
	log.Info("Seeds start...")
	query := `INSERT INTO person (person_id,age,nationality,gender,name,surname,patronymic)
				VALUES ($1,$2,$3,$4,$5,$6,$7) ON CONFLICT DO NOTHING`
	for i := 0; i < nSeeds; i++ {
		personID := uuid.New().String()
		firstName := uuid.New().String()[:30]
		_, err := ps.pgxpool.Exec(ctx, query, personID, rand.Intn(100), "RU", "male", firstName, "Petrov", "Ivanovich")
		if err != nil {
			return err
		}
	}
	t2 := time.Now()
	log.Info("Seeds done: ", t2.Sub(t1))
	return nil
}
