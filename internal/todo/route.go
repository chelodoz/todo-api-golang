package todo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewTodoRouter(todoHandler TodoHandler) *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.HandleFunc("/health", HealthCheck).Methods(http.MethodGet)
	base.HandleFunc("/todos", todoHandler.GetTodos).Methods(http.MethodGet)
	base.HandleFunc("/todos", todoHandler.CreateTodo).Methods(http.MethodPost)
	base.HandleFunc("/todos/{todoId}", todoHandler.GetTodoById).Methods(http.MethodGet)
	base.HandleFunc("/todos/{todoId}", todoHandler.UpdateTodoById).Methods(http.MethodPatch)

	return base
}

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
