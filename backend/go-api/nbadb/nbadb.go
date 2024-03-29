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

func AddAllPlayers(db *sql.DB, players []Player) error {
	for _, player := range players {
		err := AddPlayer(db, player)
		if err != nil {
			return err
		}
	}
	return nil
}

func CountPlayers(db *sql.DB) (int, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM players")
	err := row.Scan(&count)
	return count, err
}

func GetPlayerById(db *sql.DB, id int) (Player, error) {
	var player Player
	row := db.QueryRow("SELECT id, name, position, team FROM players WHERE id = $1", id)
	err := row.Scan(&player.Id, &player.Name, &player.Position, &player.Team)
	return player, err
}

func GetRandomPlayer(db *sql.DB) (Player, error) {
	var player Player
	row := db.QueryRow("SELECT id, name, position, team FROM players ORDER BY RANDOM() LIMIT 1")
	err := row.Scan(&player.Id, &player.Name, &player.Position, &player.Team)
	return player, err
}
