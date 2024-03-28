package nbadb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=nbaappdb sslmode=disable", user, pass)

	return sql.Open("postgres", connStr)
}
