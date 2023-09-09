package main

import (
	"dashboard/database"
	"dashboard/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	for _, key := range []string{
		"DYNAMOTABLE",
		"REGION",
	} {
		if os.Getenv(key) == "" {
			log.Fatalf("could not find environment variable: %s", key)
		}
	}
}

func main() {
	database.Initialize()
	log.Fatal(server.Start())
}
