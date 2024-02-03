package api

import (
	cfg "api/internal/config"
	"api/internal/handler"
	"api/internal/service"
	"api/internal/store/pg"
	"fmt"
	"net/http"
	"time"
	"context"
)

type Api struct {
	cfg	*cfg.AppConf
}

func New(cfg *cfg.AppConf) *Api{
	return &Api{cfg}
}

func (app *Api) Run() error{
	ctx := context.TODO()
	dbCtx,cancel := context.WithTimeout(ctx,time.Second * 5)
	defer cancel()
	store, err := pg.New(dbCtx,app.cfg.DBUser,app.cfg.DBPass,app.cfg.DBHost,app.cfg.DBPort,app.cfg.DBName)
	if err != nil {
		return err
	}

	personService := service.NewPersonService(
		store,
		app.cfg.AgifyAPI_URL,
		app.cfg.GenderizeAPI_URL,
		app.cfg.NatoinalizeAPI_URL,
		ctx,
	)
	handler := handler.New(personService)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d",app.cfg.AppPort),
		Handler:        handler.HTTPHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()
}