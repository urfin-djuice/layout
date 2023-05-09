package main

import (
	"log"

	"github.com/urfin-djuice/layout/internal/config"
	exampleRepository "github.com/urfin-djuice/layout/internal/repository/example"
	exampleService "github.com/urfin-djuice/layout/internal/service/example"
	"github.com/urfin-djuice/layout/internal/worker/example"
	"github.com/urfin-djuice/layout/internal/worker/loop"
	"github.com/urfin-djuice/layout/pkg/eggs"
)

func main() {
	conf := config.InitConfig()
	repo := exampleRepository.New()
	srv := exampleService.New(conf.Example.Env, conf.Example.Domain, repo)
	exampleWorker := example.New(srv)
	loopWorker := loop.New(conf.Loop.StartFrom)

	app := eggs.New()
	app.Add(exampleWorker)
	app.Add(loopWorker)

	err := app.Run()
	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Exit")
}
