package mongo

import (
	"context"
	"fmt"
	"log"
	"todo-api-golang/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongoDb take mongodb url and related to connections
func ConnectMongoDb(config config.Config) (*mongo.Client, error) {

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUsername, config.MongoPassword, config.MongoHost, config.MongoPort)

	clientOptions := options.Client().ApplyURI(url).SetRegistry(mongoRegistry)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	log.Printf("MongoClient connected")

	return client, nil
}
