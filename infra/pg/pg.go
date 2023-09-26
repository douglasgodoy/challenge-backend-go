package infra

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DbClient *sql.DB

func InitializePg() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPass)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("INITIALIZE DB")

	DbClient = db

	log.Print("RUN MIGRATIONS")
	runMigrations()
	log.Print("MIGRATIONS FINISHED")

}
