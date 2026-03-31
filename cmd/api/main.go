package main

import (
	"log"
	"net/http"
	"os"

	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/api"
	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/db"
)

func main()  {
dsn := "postgres://postgres:postgres@db:5432/jobs"
database, err := db.New(dsn)
if err != nil {
	log.Fatal(err)
	}

r := api.NewRouter()
log.Println("Starting API server on :8080")
err := http.ListenAndServe(":8080", r)
if err != nil {
	log.Fatal(err)
	}
}