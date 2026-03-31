package api

import	(
	"net/http"
	"encoding/json"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status: "OK",
		Timestamp: time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func ReadyHandler(databese *db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		err := databese.Ping(ctx)
		if err != nil {
			http.Error(w, "Database not ready", http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ready"))
	}
}