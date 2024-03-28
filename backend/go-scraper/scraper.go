package main

import (
	"fmt"
	"os"

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

	_ = user
	_ = pass

}
