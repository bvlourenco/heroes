package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sinfo.org/heroes/models"
	"sinfo.org/heroes/mongodb"
)

func heroInDB(objectID primitive.ObjectID) (int, error) {
	_, err := mongodb.GetHero(objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return http.StatusNotFound, err
		} else {
			return http.StatusInternalServerError, err
		}
	}
	return http.StatusOK, nil
}

func getHeroes(c *gin.Context) {
	heroes, err := mongodb.GetHeroes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, heroes)
}

func addHero(c *gin.Context) {
	var hero models.Hero
	err := c.BindJSON(&hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ACK, err := mongodb.AddHero(&hero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, ACK)
}

func deleteHero(c *gin.Context) {
	heroId := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(heroId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, err := heroInDB(objectID)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	ACK, err := mongodb.DeleteHero(objectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, ACK)
}

func getHero(c *gin.Context) {
	heroId := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(heroId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hero, err := mongodb.GetHero(objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, hero)
}

func updateHero(c *gin.Context) {
	var hero models.Hero
	err := c.BindJSON(&hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(hero)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	code, err := heroInDB(hero.ID)
	if err != nil {
		c.JSON(code, gin.H{"error": err.Error()})
		return
	}
	ACK, err := mongodb.UpdateHero(&hero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, ACK)
}
