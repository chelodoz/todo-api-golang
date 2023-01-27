package todo

import (
	"log"
	"net/http"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/todo/note"
	"todo-api-golang/pkg/health"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	NoteHandler note.Handler
	Router      *mux.Router
}

// NewApi creates the default configuration for the http server and set up routing.
func NewApi(config *config.Config, mongoClient *mongo.Client) *APIServer {

	noteRepository := note.NewRepository(mongoClient, config)
	noteService := note.NewService(noteRepository)
	validator := validator.New()
	if err := validator.RegisterValidation("enum", note.ValidateEnum); err != nil {
		log.Printf("Failed registering handler validators: %v", err)
	}
	noteHandler := note.NewHandler(noteService, validator)

	return &APIServer{
		NoteHandler: noteHandler,
		Router:      setupRouter(noteHandler),
	}
}

// SetupRoutes create the routes for the todo api
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
