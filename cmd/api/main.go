package main

import (
	"log"
	"net/http"

	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/api"
	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/db"
	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/config"
)

func main() {
	cfg := config.Load()

	database, err := db.New(cfg.DBUrl())
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(database)

	log.Println("Starting server on :" + cfg.AppPort)

	err = http.ListenAndServe(":"+cfg.AppPort, r)
	if err != nil {
		log.Fatal(err)
	}

	shutdown := tracing.Init()
	defer shutdown(context.Background())
}