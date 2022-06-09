package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	server := http.ListenAndServe(":"+port, router)
	log.Fatal(server)
}
