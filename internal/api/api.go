package api

import (
	cfg "api/internal/config"
	"api/internal/handler"
	"api/internal/service"
	"api/internal/store/pg"
	"api/internal/clients/agify"
	"api/internal/clients/nationalize"
	"api/internal/clients/genderize"
	"fmt"
	"net/http"
	"time"
	"context"
)

type Api struct {
	cfg	*cfg.Config
}

func New(cfg *cfg.Config) *Api{
	return &Api{cfg}
}

func (app *Api) Run() error{
	ctx := context.Background()
	store, err := pg.New(ctx,app.cfg.PgConf)
	
	if err != nil {
		return err
	}

	serviceCfg := &cfg.ServiceConfig{
		store,
		nationalize.New(app.cfg.NatoinalizeAPI_URL),
		genderize.New(app.cfg.GenderizeAPI_URL),
		agify.New(app.cfg.AgifyAPI_URL),
	}

	service := service.New(serviceCfg)
	handler := handler.New(service)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d",app.cfg.AppPort),
		Handler:        handler.HTTPHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}