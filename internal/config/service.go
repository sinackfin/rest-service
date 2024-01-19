package config

import (
	"api/internal/store"
	"api/internal/clients/agify"
	"api/internal/clients/genderize"
	"api/internal/clients/nationalize"
)

type ServiceConfig struct{
	Store				store.IStore
	Nationalize			*nationalize.Nationalize
	Genderize			*genderize.Genderize
	Agify				*agify.Agify
}