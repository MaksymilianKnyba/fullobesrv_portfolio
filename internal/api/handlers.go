import	{
	"net/http"
	"encoding/json"
	"time"
}

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