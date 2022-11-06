package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBClient mongo.Database

func CreateClient() (*DBClient, func() error, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	disconnect := func() error { return client.Disconnect(ctx) }
	return (*DBClient)(client.Database("GoBun")), disconnect, err
}
