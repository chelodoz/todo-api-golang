package mongo

import (
	"context"
	"fmt"
	"log"
	"time"
	"todo-api-golang/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDb take mongodb url and related to connections
func ConnectMongoDb(config *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUsername, config.MongoPassword, config.MongoHost, config.MongoPort)
	clientOptions := options.Client().ApplyURI(url).SetRegistry(mongoRegistry)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	log.Printf("MongoClient connected")

	return client, nil
}
