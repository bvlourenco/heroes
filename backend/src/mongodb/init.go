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
var heroDatabase *mongo.Database

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

	heroDatabase = dbClient.Database(dbName)
	HeroCollection = heroDatabase.Collection("heroes")
}

func DropDatabase() {
	if heroDatabase != nil {
		heroDatabase.Drop(ctx)
	}
}

func CloseMongoConnection() {
	err := dbClient.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[DEBUG] Connected to MongoDB closed!")
}
