package main

import (
	"log"
	"net/http"

	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/api"
)

func main()  {
r := api.NewRouter()
log.Println("Starting API server on :8080")
err := http.ListenAndServe(":8080", r)
if err != nil {
	log.Fatal(err)
	}
}