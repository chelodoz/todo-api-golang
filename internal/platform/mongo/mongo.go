package mongo

import (
	"context"
	"fmt"
	"time"
	"todo-api-golang/internal/config"
	"todo-api-golang/pkg/logs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDbClient takes mongodb configuration and returns a mongo client
func NewDbClient(config *config.Config, logs *logs.Logs) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.MongoUsername, config.MongoPassword, config.MongoHost, config.MongoPort)
	clientOptions := options.Client().ApplyURI(url).SetRegistry(mongoRegistry)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logs.Logger.Error("Failed to connect to MongoDB")
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Check the connection
	if err = client.Ping(ctx, nil); err != nil {
		logs.Logger.Error("Ping command to MongoDB client failed")
		return nil, err
	}
	logs.Logger.Info("MongoDB client connected")

	return client, nil
}
