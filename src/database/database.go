package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const database = "ToDoGo"

func ConnectionMongo(collection string) *mongo.Collection {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connect, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb"))
	if err != nil {
		panic(err)
	}
	return connect.Database(database).Collection(collection)
}
