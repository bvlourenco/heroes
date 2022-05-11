package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"sinfo.org/heroes/mongodb"
	"sinfo.org/heroes/router"
)

func main() {

	debug := flag.Bool("debug", false, "if true sets program to debug mode")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

    mongoURL, dbName := os.Getenv("MONGO_URL"), os.Getenv("DB_NAME")
	mongodb.InitMongoConnection(mongoURL, dbName)
	defer mongodb.CloseMongoConnection()

    host, port := os.Getenv("HOST"), os.Getenv("PORT")
	router.InitRouter(host, port, *debug)
}
