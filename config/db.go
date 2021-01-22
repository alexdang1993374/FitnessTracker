package config

import (
	"log"
	"os"

	"fitness-tracker/controllers"

	"github.com/go-pg/pg/v9"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "CC1993374",
		Addr:     "localhost:5432",
		Database: "fitness",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	controllers.CreateFitnessTable(db)
	return db
}
