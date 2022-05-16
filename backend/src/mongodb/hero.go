package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sinfo.org/heroes/models"
)

var HeroCollection *mongo.Collection

func GetHeroes() ([]models.Hero, error) {
	var heroes []models.Hero

	//filter that matches all documents
	filter := bson.M{}
	cursor, err := HeroCollection.Find(ctx, filter)
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
	result, err := HeroCollection.InsertOne(ctx, hero)
	if err != nil {
		return "", err
	}
	ACK := "Inserted a hero with ID:" + fmt.Sprintf("%v", result.InsertedID.(primitive.ObjectID).Hex())
	fmt.Println("[DEBUG] " + ACK)
	return ACK, nil
}

func DeleteHero(heroId primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": heroId}

	result, err := HeroCollection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}
	ACK := "Deleted " + fmt.Sprintf("%v", result.DeletedCount) + " heroes."
	fmt.Println("[DEBUG] " + ACK)
	return ACK, nil
}

func GetHero(heroId primitive.ObjectID) (*models.Hero, error) {
	var hero *models.Hero
	filter := bson.M{"_id": heroId}

	err := HeroCollection.FindOne(ctx, filter).Decode(&hero)
	if err != nil {
		return nil, err
	}
	return hero, nil
}

func UpdateHero(hero *models.Hero) (string, error) {
	filter := bson.M{"_id": hero.ID}
	update := bson.M{"name": hero.Name}
	result, err := HeroCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return "", err
	}
	ACK := "Matched " + fmt.Sprintf("%v", result.MatchedCount) + " heroes and updated " + fmt.Sprintf("%v", result.ModifiedCount) + " heroes."
	fmt.Println("[DEBUG] " + ACK)
	return ACK, nil
}
