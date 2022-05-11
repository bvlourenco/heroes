package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sinfo.org/heroes/models"
	"sinfo.org/heroes/mongodb"
)

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	ACK, err := mongodb.UpdateHero(&hero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, ACK)
}
