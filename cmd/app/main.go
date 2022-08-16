package main

import (
	"log"

	"github.com/doyyan/kubernetes/cmd/app/config"
	"github.com/doyyan/kubernetes/internal/app/adapter/controller"
	"github.com/doyyan/kubernetes/internal/app/adapter/kubernetes"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

func main() {
	// create a new logger for cross application logging
	logger := logrus.New()
	ctx := context.Background()
	config, err := config.LoadConfig(ctx, ".", logger)
	if err = postgresql.CreateDBConnection(ctx, config, logger); err != nil {
		logger.Fatal(err)
	}
	kube := kubernetes.Kube{}
	err = kube.SetConfig(ctx, logger)
	if err != nil {
		logger.Fatal(err)
	}
	c := controller.Controller{Context: ctx, Logger: logger}
	r := c.Router()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	r.Run(config.ADDRESS)
}
