package main

import (
	"log"
	"net/http"
)

const port string = ":8080"

func main() {

	router := http.NewServeMux()
	router.HandleFunc("GET /hello", handleHello)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Starting server on port %s\n", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Println(err)
	}

	log.Println("Shutting down server...")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("/hello request from %s\n", r.RemoteAddr)
	w.Write([]byte("Hello World"))
}
