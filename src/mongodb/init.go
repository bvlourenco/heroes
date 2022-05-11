package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var dbClient *mongo.Client

func InitMongoConnection(mongoURL, dbName string) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	ctx := context.Background()

	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[DEBUG] Connected to MongoDB!")

	heroCollection = dbClient.Database(dbName).Collection("heroes")
}

func CloseMongoConnection() {
	err := dbClient.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[DEBUG] Connected to MongoDB closed!")
}
