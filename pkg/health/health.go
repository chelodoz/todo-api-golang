package health

import (
	"fmt"
	"net/http"
)

// swagger:route GET /health Health Health
// Check health of the api
//
// Check health of the api
//
// responses:
// 200: helthReponseWrapper

// HealthCheck return a Healthy message in the response
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Healthy")
}

// Returns Healthy if the api is working
// swagger:response helthReponseWrapper
type helthReponseWrapper string
