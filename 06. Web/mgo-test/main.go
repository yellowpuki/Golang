package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	dbHost         string = "mongodb://172.18.104.106:27017"
	dbName         string = "todo_demo"
	collectionName string = "todo"
)

var (
	collection *mongo.Collection
	ctx        = context.TODO()
)

func init() {
	clientOption := options.Client().ApplyURI(dbHost)
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	collection = client.Database(dbName).Collection(collectionName)
}

func main() {

}
