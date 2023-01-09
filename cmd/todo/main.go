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
	config, err := config.LoadConfig("./../../config")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	startHTTPServer(config)
}

func startHTTPServer(config config.Config) {

	mongoClient, err := mongo.ConnectMongoDb(config)
	if err != nil {
		log.Fatalf("Error starting mongo client: %s\n", err)
	}

	router := todo.NewApi(config, mongoClient)

	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}

	// Swagger
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/api/v1/swagger/", http.FileServer(http.Dir("./third_party/swagger-ui-4.11.1"))))

	// CORS
	cors := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	log := log.New(os.Stdout, "todo-api-golang ", log.LstdFlags)

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
			log.Fatalf("Error starting server: %s\n", err)
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
