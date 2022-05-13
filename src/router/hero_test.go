package router

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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

func makeRequest(t *testing.T, method string, path string, payload io.Reader) *http.Response {
	req, err := http.NewRequest(method, path, payload)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	assert.NilError(t, err)
	return res
}

func readResponse(resBody io.ReadCloser) string {
	body, err := ioutil.ReadAll(resBody)
	if err != nil {
		log.Fatal(err)
	}
	resBody.Close()
	return string(body)
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

	res := makeRequest(t, "POST", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusOK)

	resContent := readResponse(res.Body)

	heroDB, err := getHeroDB(heroName)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, resContent, "Inserted a hero with ID:"+heroDB.ID.Hex())
}

func TestAddHeroInvalidType(t *testing.T) {
	mongodb.DropDatabase()

	type InvalidHero struct {
		ID   string `json:"id,omitempty" bson:"_id,omitempty"`
		Name string `json:"name"`
	}

	hero := &InvalidHero{ID: "invalid_id_type", Name: "invalid_hero_id"}

	bytesHero, errMarshal := json.Marshal(hero)
	assert.NilError(t, errMarshal)

	res := makeRequest(t, "POST", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestAddHeroInvalidFormat(t *testing.T) {
	mongodb.DropDatabase()

	type InvalidHero struct {
		ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		OtherName string             `json:"othername" bson:"othername"`
	}

	hero := &InvalidHero{OtherName: "invalid_hero_format"}

	bytesHero, errMarshal := json.Marshal(hero)
	assert.NilError(t, errMarshal)

	res := makeRequest(t, "POST", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestGetHeroes(t *testing.T) {
	mongodb.DropDatabase()
	var heroes []*models.Hero
	hero_test_1 := createHero(t, "hero_test_1")
	hero_test_2 := createHero(t, "hero_test_2")
	heroes = append(heroes, hero_test_1)
	heroes = append(heroes, hero_test_2)
	res := makeRequest(t, "GET", url+"/hero", nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	var heroesRequest []*models.Hero
	json.NewDecoder(res.Body).Decode(&heroesRequest)
	for i := 0; i < len(heroes); i++ {
		compareHero(t, heroesRequest[i], heroes[i])
	}
}

func TestGetHero(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")
	res := makeRequest(t, "GET", url+"/hero/"+hero_test_1.ID.Hex(), nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	var hero *models.Hero
	json.NewDecoder(res.Body).Decode(&hero)
	compareHero(t, hero, hero_test_1)
}

func TestGetHeroInvalidID(t *testing.T) {
	mongodb.DropDatabase()
	res := makeRequest(t, "GET", url+"/hero/"+"invalid_ID", nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestGetHeroNotFound(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := &models.Hero{Name: "hero_test_1"}
	res := makeRequest(t, "GET", url+"/hero/"+hero_test_1.ID.Hex(), nil)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}

func TestDeleteHero(t *testing.T) {
	mongodb.DropDatabase()
	heroName := "hero_test_1"
	hero_test_1 := createHero(t, heroName)

	res := makeRequest(t, "DELETE", url+"/hero/"+hero_test_1.ID.Hex(), nil)
	assert.Equal(t, res.StatusCode, http.StatusOK)

	resContent := readResponse(res.Body)

	_, err := getHeroDB(heroName)
	//When going to DB, it is supposed to return error mongo.ErrNoDocuments
	//because the hero was already deleted from the DB with DELETE request
	if err == nil || (err != nil && err != mongo.ErrNoDocuments) {
		log.Fatal(err)
	}

	assert.Equal(t, resContent, "Deleted 1 heroes.")
}

func TestDeleteHeroInvalidID(t *testing.T) {
	mongodb.DropDatabase()
	heroName := "hero_test_1"
	_ = createHero(t, heroName)

	res := makeRequest(t, "DELETE", url+"/hero/"+"invalid_id", nil)
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	_, err := getHeroDB(heroName)
	if err != nil {
		log.Fatal(err)
	}
}

func TestDeleteHeroNotFound(t *testing.T) {
	mongodb.DropDatabase()
	heroName := "hero_test_1"
	hero_test_1 := &models.Hero{Name: heroName}

	res := makeRequest(t, "DELETE", url+"/hero/"+hero_test_1.ID.Hex(), nil)
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}

func TestUpdateHero(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")
	hero_test_1.Name = "hero_test_1_updated"
	bytesHero, errMarshal := json.Marshal(hero_test_1)

	assert.NilError(t, errMarshal)

	res := makeRequest(t, "PUT", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusOK)

	resContent := readResponse(res.Body)

	heroDB, err := getHeroDB("hero_test_1_updated")
	if err != nil {
		log.Fatal(err)
	}

	compareHero(t, hero_test_1, heroDB)
	assert.Equal(t, resContent, "Matched 1 heroes and updated 1 heroes.")
}

func TestUpdateHeroInvalidFormat(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := createHero(t, "hero_test_1")

	type InvalidHero struct {
		ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		OtherName string             `json:"othername" bson:"othername"`
	}

	hero := &InvalidHero{ID: hero_test_1.ID, OtherName: "invalid_hero_format"}

	bytesHero, errMarshal := json.Marshal(hero)
	assert.NilError(t, errMarshal)

	res := makeRequest(t, "PUT", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
}

func TestUpdateHeroNotFound(t *testing.T) {
	mongodb.DropDatabase()
	hero_test_1 := &models.Hero{Name: "hero_test_1"}
	hero_test_1.Name = "hero_test_1_updated"
	bytesHero, errMarshal := json.Marshal(hero_test_1)

	assert.NilError(t, errMarshal)

	res := makeRequest(t, "PUT", url+"/hero", bytes.NewBuffer(bytesHero))
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}
