package main

import (
	"log"

	"github.com/doyyan/kubernetes/cmd/app/config"
	"github.com/doyyan/kubernetes/internal/app/adapter/controller"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
	"github.com/sirupsen/logrus"
)

func main() {
	// create a new logger for cross application logging
	logger := logrus.New()
	config, err := config.LoadConfig(".", logger)
	if err = postgresql.CreateDBConnection(config, logger); err != nil {
		logger.Fatal(err)
	}
	c := controller.Controller{logger}
	r := c.Router()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	r.Run(config.ADDRESS)
}
