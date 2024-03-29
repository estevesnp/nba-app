package nbadb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Player struct {
	Id       int
	Name     string
	Position string
	Team     string
}

func (p Player) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Position: %s, Team: %s", p.Id, p.Name, p.Position, p.Team)
}

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

func AddPlayer(db *sql.DB, player Player) error {
	_, err := db.Exec("INSERT INTO players (id, name, position, team) VALUES ($1, $2, $3, $4)", player.Id, player.Name, player.Position, player.Team)
	return err
}
