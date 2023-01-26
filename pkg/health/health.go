package health

import (
	"net/http"
)

// swagger:route GET /health Health Health
// Check health of the api
//
// Check health of the api
//
// responses:
// 200: HealthResponseWrapper

// HealthCheck return a Healthy message in the response
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"Healthy"}`))
}

// Returns Healthy if the api is working
// swagger:response HealthResponseWrapper
type HealthResponseWrapper struct {
	// in: body
	Body struct {
		// example: Healthy
		Status string `json:"status"`
	}
}
