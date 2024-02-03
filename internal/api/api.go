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

	service := service.New(
		store,
		agify.New(app.cfg.AgifyAPI_URL),
		genderize.New(app.cfg.GenderizeAPI_URL),
		nationalize.New(app.cfg.NatoinalizeAPI_URL),
		ctx,
	)
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