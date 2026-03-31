package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/maksymilianKnyba/fullobesrv_portfolio/internal/db"
)

func NewRouter(database *db.DB) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Get("/health", HealthHandler)
	r.Get("/ready", ReadyHandler(database))
	r.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		MetricsHandler().ServeHTTP(w, r)
	})

	return r
}