package todo

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewTodoRouter(todoHandler TodoHandler) *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.HandleFunc("/todos", todoHandler.GetTodos).Methods(http.MethodGet)
	base.HandleFunc("/todos", todoHandler.CreateTodo).Methods(http.MethodPost)
	base.HandleFunc("/todos/{todoId}", todoHandler.GetTodoById).Methods(http.MethodGet)

	return base
}
