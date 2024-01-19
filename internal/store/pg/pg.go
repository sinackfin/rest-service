package pg

import (
	"api/internal/config"
	"api/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"fmt"
)

type PGStore struct {
	cfg 	*config.PGStoreConfg
	pgxpool *pgxpool.Pool
}

func New(ctx context.Context,cfg *config.PGStoreConfg) (*PGStore,error) {
	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",cfg.DB_USER,cfg.DB_PASS,cfg.DB_HOST,cfg.DB_PORT,cfg.DB_DBNAME)
	pgxpool, err :=  pgxpool.New(ctx, dbConnStr)
	if err != nil {
		return nil, err
	}
	if err = pgxpool.Ping(ctx); err != nil {
		return nil,err
	}
	return &PGStore {
		cfg,
		pgxpool,
	},nil
}


func (ps *PGStore) CreatePerson(person *models.Person) error{
	query := `INSERT INTO person (person_id,age,nationality,gender,name,surname,patronymic)
				VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_,err := ps.pgxpool.Query(context.Background(),query,person.ID,person.Age,person.Nationality,
					person.Gender,person.Name,person.Surname,person.Patronymic)
	return err
}

func (ps *PGStore) GetPersonByID(id string) (*models.Person,error) {
	query := `SELECT person_id,age,nationality,gender,name,surname,patronymic from person WHERE person_id=$1`
	person := models.Person{}
	err := ps.pgxpool.QueryRow(context.Background(), query, id).Scan(&person.ID,&person.Age,
		&person.Nationality,&person.Gender,&person.Name,&person.Surname,&person.Patronymic)
	if err != nil {
		return nil,err
	}
	return &person,nil
}
func (ps *PGStore) UpdatePerson(person *models.Person) error {
	query := `UPDATE person SET age=$1,nationality=$2,gender=$3,name=$4,surname=$5,patronymic=$6 WHERE person_id=$7`
	_,err := ps.pgxpool.Query(context.Background(),query,person.Age,person.Nationality,
	person.Gender,person.Name,person.Surname,person.Patronymic,person.ID)
	return err
}
func (ps *PGStore) DeletePersonByID(id string) error {
	query := `DELETE from person WHERE person_id=$1`
	_,err := ps.pgxpool.Query(context.Background(),query,id)
	return err
}
