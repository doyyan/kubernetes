package main

import (
	"log"

	"github.com/doyyan/kubernetes/cmd/app/config"
	"github.com/doyyan/kubernetes/internal/app/adapter"
	"github.com/doyyan/kubernetes/internal/app/adapter/postgresql"
)

func main() {
	config, err := config.LoadConfig(".")
	postgresql.CreateDBConnection(config)
	r := adapter.Router()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r.Run(":8080")
}
