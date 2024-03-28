package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=postgres sslmode=disable", user, pass)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error reading SQL file:", err)
	}

	queries := strings.Split(string(schema), ";")

	for _, q := range queries {

		_, err = db.Exec(q)
		if err != nil {
			log.Fatalln("Error executing SQL script:", err)
		}

	}

	fmt.Println("Schema created successfully")
}
