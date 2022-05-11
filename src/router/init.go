package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func InitRouter(host, port string, debug bool) {
	if debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	validate = validator.New()

	router.GET("/ping", ping)
	router.GET("/hero", getHeroes)
	router.POST("/hero", addHero)
	router.DELETE("/hero/:id", deleteHero)
	router.GET("/hero/:id", getHero)
	router.PUT("/hero", updateHero)

	addr := host + ":" + port
	fmt.Println("[DEBUG] Running server on:", addr)
	log.Fatal(router.Run(addr))
}
