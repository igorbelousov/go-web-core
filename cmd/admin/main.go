package main

import (
	"fmt"
	"log"

	"github.com/igorbelousov/go-web-core/foundation/database"
	"github.com/igorbelousov/go-web-core/internal/data/schema"
)

func main() {
	migrate()
}

func migrate() {

	dbConfig := database.Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       "0.0.0.0",
		Name:       "postgres",
		DisableTLS: true,
	}

	db, err := database.Open(dbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := schema.Migrate(db); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("migrate comlete")

	if err := schema.Seed(db); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("seed data comlete")
}
