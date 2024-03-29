package main

import (
	"esteves/nba-api-server/nbadb"
	"esteves/nba-api-server/scraper"

	"log"
	"net/http"
)

const port string = ":8080"

func main() {

	log.Println("Connecting to database")
	db, err := nbadb.OpenDB()
	if err != nil {
		log.Fatal("Failed connecting to database:", err)
	}
	defer db.Close()

	playerCount, err := nbadb.CountPlayers(db)
	if err != nil {
		log.Fatal("Failed to count players:", err)
	}

	if playerCount == 0 {
		log.Println("Adding players to database")

		players, err := scraper.GetPlayerData()
		if err != nil {
			log.Fatal("Failed to get player data:", err)
		}

		err = nbadb.AddAllPlayers(db, players)
		if err != nil {
			log.Fatal("Failed to add players to database:", err)
		}
	}

	server := getServer()

	log.Println("Starting server on port", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer server.Close()

	log.Println("Shutting down server")
}

func getServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello", handleHello)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	return &server
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("/hello request from", r.RemoteAddr)
	w.Write([]byte("Hello World"))
}
