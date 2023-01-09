package todo

import (
	"net/http"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/todo/note"
	"todo-api-golang/pkg/health"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewApi creates the default configuration for the http server and set up routing.
func NewApi(config config.Config, mongoClient *mongo.Client) *mux.Router {

	noteRepository := note.NewRepository(mongoClient, &config)
	noteService := note.NewService(noteRepository)
	noteHandler := note.NewHandler(noteService)

	return setupRouter(noteHandler)
}

func setupRouter(noteHandler note.Handler) *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.GetAll).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.Create).Methods(http.MethodPost)
	base.HandleFunc("/notes/{noteId}", noteHandler.GetById).Methods(http.MethodGet)
	base.HandleFunc("/notes/{noteId}", noteHandler.Update).Methods(http.MethodPatch)

	return base
}
