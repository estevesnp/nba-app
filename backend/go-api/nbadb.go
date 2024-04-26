package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Player struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Team     string `json:"team"`
}

func (p Player) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Position: %s, Team: %s", p.Id, p.Name, p.Position, p.Team)
}

var (
	host   string
	port   string
	user   string
	pass   string
	dbname string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, setting default values:", err)

		host = "localhost"
		port = "5432"
		user = "postgres"
		pass = "password"
		dbname = "nbaappdb"

		return
	}

	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || pass == "" || dbname == "" {
		log.Fatal("DB_HOST, DB_PORT, DB_USER and DB_PASSWORD must be set in .env file")
	}
}

func OpenDB() (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	connTries := 10
	for i := 0; i < connTries; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		log.Println("Failed to connect to database, retrying in 5 seconds")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	return db, err
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

func GetAllPlayers(db *sql.DB) ([]Player, error) {
	rows, err := db.Query("SELECT id, name, position, team FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.Id, &player.Name, &player.Position, &player.Team)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
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
