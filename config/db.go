package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := os.Getenv("DB_URL")

	if connStr == "" {
		log.Println("Using default local DB config")
		connStr = "postgres://postgres:password@localhost:5432/coachdb?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}