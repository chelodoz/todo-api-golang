package todo

import (
	"net/http"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/todo/note"
	"todo-api-golang/internal/trace"
	"todo-api-golang/pkg/health"
	"todo-api-golang/pkg/logs"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type APIServer struct {
	NoteHandler note.Handler
	Router      *mux.Router
	Logger      *logs.Logs
}

// NewApi creates the default configuration for the http server and set up routing.
func NewApi(config *config.Config, mongoClient *mongo.Client, logs *logs.Logs) (*APIServer, error) {

	noteRepository := note.NewRepository(mongoClient, config)
	noteService := note.NewService(noteRepository)
	validator := validator.New()
	if err := validator.RegisterValidation("enum", note.ValidateEnum); err != nil {
		logs.Logger.Error("Failed registering validators for handlers", zap.String("details", err.Error()))
		return nil, err
	}
	noteHandler := note.NewHandler(noteService, validator)

	return &APIServer{
		NoteHandler: noteHandler,
		Router:      setupRouter(noteHandler, logs),
		Logger:      logs,
	}, nil
}

// SetupRoutes create the routes for the todo api
func setupRouter(noteHandler note.Handler, log *logs.Logs) *mux.Router {
	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.Use(trace.ContextIDMiddleware(log))
	base.Use(LogMiddleware(log))

	base.HandleFunc("/health", health.HealthCheck).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.GetAll).Methods(http.MethodGet)
	base.HandleFunc("/notes", noteHandler.Create).Methods(http.MethodPost)
	base.HandleFunc("/notes/{noteId}", noteHandler.GetById).Methods(http.MethodGet)
	base.HandleFunc("/notes/{noteId}", noteHandler.Update).Methods(http.MethodPatch)

	return base
}
