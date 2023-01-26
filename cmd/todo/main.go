package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"todo-api-golang/internal/config"
	"todo-api-golang/internal/platform/mongo"
	"todo-api-golang/internal/todo"

	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	config, err := config.LoadConfig("./../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	startHTTPServer(config)
}

func startHTTPServer(config *config.Config) {

	mongoClient, err := mongo.ConnectMongoDb(config)
	if err != nil {
		log.Fatalf("Error starting mongo client: %s\n", err)
	}

	todoApi := todo.NewApi(config, mongoClient)
	todoRouter := todoApi.SetupRouter()

	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}

	// Swagger
	todoRouter.PathPrefix("/swagger/").Handler(http.StripPrefix("/api/v1/swagger/", http.FileServer(http.Dir("./third_party/swagger-ui-4.11.1"))))

	// CORS
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	log := log.New(os.Stdout, "todo-api-golang ", log.LstdFlags)

	// create a new server
	server := http.Server{
		Addr:         config.HTTPServerAddress, // configure the bind address
		Handler:      cors(todoRouter),         // set the default handler
		ErrorLog:     log,                      // set the logger for the server
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		log.Printf("Starting server on port: %v", config.HTTPServerAddress)

		err := http.ListenAndServe(config.HTTPServerAddress, todoRouter)
		if err != nil {
			log.Fatalf("Error starting server: %s\n", err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// block until a signal is received.
	sig := <-ch
	log.Println("Shutdown signal received:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err = server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server shutdown completed")
}
