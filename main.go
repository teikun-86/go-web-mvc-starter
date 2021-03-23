package main

import (
	"log"
	"net/http"
	"os"

	"github.com/azizsama/go-web-mvc-starter/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	routes.Start(mux)
	http.ListenAndServe(":"+port, mux)
}
