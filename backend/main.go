package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/db"
	"backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using Docker env vars")
	}

	db.Init()
	r := routes.Router()

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
