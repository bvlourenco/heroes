package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sinfo.org/heroes/models"
)

var heroCollection *mongo.Collection

func GetHeroes() ([]models.Hero, error) {
	var heroes []models.Hero

	//filter that matches all documents
	filter := bson.M{}
	cursor, err := heroCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var hero models.Hero
		err := cursor.Decode(&hero)
		if err != nil {
			return nil, err
		}

		heroes = append(heroes, hero)
	}
	return heroes, nil
}

func AddHero(hero *models.Hero) (string, error) {
	result, err := heroCollection.InsertOne(ctx, hero)
	if err != nil {
		return "", err
	}
	ACK := "[DEBUG] Inserted a hero with ID:" + fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex())
	fmt.Println(ACK)
	return ACK, nil
}

func DeleteHero(heroId primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": heroId}

	result, err := heroCollection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	ACK := "[DEBUG] Deleted " + fmt.Sprintf("%v", result.DeletedCount) + " heroes."
	fmt.Println(ACK)
	return ACK, nil
}

func GetHero(heroId primitive.ObjectID) (*models.Hero, error) {
	var hero *models.Hero
	filter := bson.M{"_id": heroId}

	err := heroCollection.FindOne(ctx, filter).Decode(&hero)
	if err != nil {
		return nil, err
	}
	return hero, nil
}

func UpdateHero(hero *models.Hero) (string, error) {
	filter := bson.M{"_id": hero.ID}
	update := bson.M{"name": hero.Name}
	result, err := heroCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return "", err
	}
	ACK := "[DEBUG] Matched " + fmt.Sprintf("%v", result.MatchedCount) + " heroes and updated " + fmt.Sprintf("%v", result.ModifiedCount) + " heroes."
	fmt.Println(ACK)
	return ACK, nil
}
