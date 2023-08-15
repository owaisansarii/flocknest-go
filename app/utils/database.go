// utils/database.go
package utils

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	dbName     = "flocknest"
	mongoURI   = "mongodb://localhost:27017"
	maxTimeout = 10 * time.Second
	err        error
)

func ConnectDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), maxTimeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	// col := client.Database("flocknest").Collection("users")
	// col.InsertOne(ctx, map[string]string{"name": "piyush"})

	return nil
}

func GetDatabase() (*mongo.Database, error) {
	if client == nil {
		return nil, errors.New("database not connected")
	}
	return client.Database(dbName), nil
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	if client == nil {
		return nil, errors.New("database not connected")
	}

	collection := client.Database(dbName).Collection(collectionName)
	return collection, nil
}

func GetMongoClient() (*mongo.Client, error) {
	if client == nil {
		return nil, errors.New("database not connected")
	}
	return client, nil
}
