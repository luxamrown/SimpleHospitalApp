package main

import (
	"enigmacamp.com/rumah_sakit/config"
	"enigmacamp.com/rumah_sakit/delivery/api"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (a *appServer) initHandlers() {
	a.v1()
}

func (a *appServer) v1() {
	patientApiGroup := a.routerEngine.Group("/patient")
	api.NewPatientApi(patientApiGroup, a.cfg.UseCaseManager.AddPatientUseCase(), a.cfg.UseCaseManager.SetDonePatientUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.routerEngine.Run(a.cfg.ApiConfig.Url)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
