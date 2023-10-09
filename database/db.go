package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var client *mongo.Client

func Connect() (*mongo.Database, error) {
	// Return the existing client and database if already initialized
	if client != nil {
		return db, nil
	}

	// Create a new MongoDB client using the mongo.Connect method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	newClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check if the client is properly connected to the MongoDB server
	err = newClient.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Set the global client variable
	client = newClient
	db = client.Database("my-blog")

	return db, nil
}

// Close the MongoDB client when your application exits
func Disconnect() {
	if client != nil {
		_ = client.Disconnect(context.Background())
	}
}
