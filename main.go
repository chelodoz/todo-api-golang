package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sample-golang-api/handler"
	"sample-golang-api/repository"
	"sample-golang-api/service"
	"sample-golang-api/util"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	startHTTPServer(config)
}

func startHTTPServer(config util.Config) {

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router := mux.NewRouter()
	base := router.PathPrefix("/api/v1").Subrouter()

	base.HandleFunc("/todos", todoHandler.GetTodos).Methods(http.MethodGet)
	base.HandleFunc("/todos", todoHandler.CreateTodo).Methods(http.MethodPost)
	base.HandleFunc("/todos/{todoId}", todoHandler.GetTodoById).Methods(http.MethodGet)

	// Swagger
	// router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./docs/swagger-ui-4.11.1"))))

	// CORS
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	log := log.New(os.Stdout, "sample-golang-api ", log.LstdFlags)
	// create a new server
	server := http.Server{
		Addr:         config.HTTPServerAddress, // configure the bind address
		Handler:      cors(router),             // set the default handler
		ErrorLog:     log,                      // set the logger for the server
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Printf("Starting server on port: %v", config.HTTPServerAddress)

		err := http.ListenAndServe(config.HTTPServerAddress, router)
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-ch
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
