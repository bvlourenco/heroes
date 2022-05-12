package router

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gotest.tools/assert"
	"sinfo.org/heroes/models"
	"sinfo.org/heroes/mongodb"
)

var url string

func insertHeroDB(t *testing.T, hero *models.Hero) primitive.ObjectID {
	res, err := mongodb.HeroCollection.InsertOne(context.Background(), hero)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID.(primitive.ObjectID)
}

func compareHero(t *testing.T, hero1, hero2 *models.Hero) {
	assert.Equal(t, hero1.ID, hero2.ID)
	assert.Equal(t, hero1.Name, hero2.Name)
}

func createHero(t *testing.T, name string) *models.Hero {
	hero_test := &models.Hero{Name: name}
	hero_test_id := insertHeroDB(t, hero_test)
	hero_test.ID = hero_test_id
	return hero_test
}

func getHeroDB(name string) (*models.Hero, error) {
	var hero *models.Hero
	err := mongodb.HeroCollection.FindOne(context.Background(), bson.M{"name": name}).Decode(&hero)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	return hero, nil
}

func setupTests() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	mongoURL, dbName := os.Getenv("MONGO_URL"), os.Getenv("DB_NAME")
	mongodb.InitMongoConnection(mongoURL, dbName)

	host, port := os.Getenv("HOST"), os.Getenv("PORT")
	url = "http://" + host + ":" + port
}

func TestMain(m *testing.M) {
	setupTests()
	os.Exit(m.Run())
	mongodb.CloseMongoConnection()
}

func TestAddHero(t *testing.T) {
	mongodb.DropDatabase()
	heroName := "hero_test"
	hero := &models.Hero{Name: heroName}
	bytesHero, errMarshal := json.Marshal(hero)

	assert.NilError(t, errMarshal)

	res, err := http.Post(url+"/hero", "application/json", bytes.NewBuffer(bytesHero))
	assert.NilError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	heroDB, err := getHeroDB(heroName)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, string(body), "Inserted a hero with ID:"+heroDB.ID.Hex())
}

func TestGetHeroes(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")
	hero_test_2 := createHero(t, "hero_test_2")
	res, err := http.Get(url + "/hero")
	assert.NilError(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	var heroes []*models.Hero
	json.NewDecoder(res.Body).Decode(&heroes)
	compareHero(t, heroes[0], hero_test_1)
	compareHero(t, heroes[1], hero_test_2)
}

func TestGetHero(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")
	res, err := http.Get(url + "/hero/" + hero_test_1.ID.Hex())
	assert.NilError(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	var hero *models.Hero
	json.NewDecoder(res.Body).Decode(&hero)
	compareHero(t, hero, hero_test_1)
}

func TestDeleteHero(t *testing.T) {
	mongodb.DropDatabase()
	heroName := "hero_test_1"
	hero_test_1 := createHero(t, heroName)

	req, err := http.NewRequest("DELETE", url+"/hero/"+hero_test_1.ID.Hex(), nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_, err = getHeroDB(heroName)
	if err == nil || (err != nil && err != mongo.ErrNoDocuments) {
		log.Fatal(err)
	}

	assert.Equal(t, string(body), "Deleted 1 heroes.")
}

func TestUpdateHero(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")
	hero_test_1.Name = "hero_test_1_updated"
	bytesHero, errMarshal := json.Marshal(hero_test_1)

	assert.NilError(t, errMarshal)

	req, err := http.NewRequest("PUT", url+"/hero", bytes.NewBuffer(bytesHero))
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	heroDB, err := getHeroDB("hero_test_1_updated")
	if err != nil {
		log.Fatal(err)
	}

	compareHero(t, hero_test_1, heroDB)
	assert.Equal(t, string(body), "Matched 1 heroes and updated 1 heroes.")
}
