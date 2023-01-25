package todo

import (
	"net/http"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/todo/note"
	"todo-api-golang/pkg/health"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	NoteHandler note.Handler
}

// NewApi creates the default configuration for the http server and set up routing.
func NewApi(config config.Config, mongoClient *mongo.Client) *APIServer {

	noteRepository := note.NewRepository(mongoClient, &config)
	noteService := note.NewService(noteRepository)
	noteHandler := note.NewHandler(noteService)

	return &APIServer{
		NoteHandler: noteHandler,
	}
}

// SetupRoutes create the routes for the todo api
func (api *APIServer) SetupRouter() *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)
	base.HandleFunc("/notes", api.NoteHandler.GetAll).Methods(http.MethodGet)
	base.HandleFunc("/notes", api.NoteHandler.Create).Methods(http.MethodPost)
	base.HandleFunc("/notes/{noteId}", api.NoteHandler.GetById).Methods(http.MethodGet)
	base.HandleFunc("/notes/{noteId}", api.NoteHandler.Update).Methods(http.MethodPatch)

	return base
}
