package main

import (
	"esteves/nba-api-server/nbadb"

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

	router := http.NewServeMux()
	router.HandleFunc("GET /hello", handleHello)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println("Starting server on port", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

	defer server.Close()

	log.Println("Shutting down server")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("/hello request from", r.RemoteAddr)
	w.Write([]byte("Hello World"))
}
