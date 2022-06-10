package main

import (
	"log"
	"net/http"
	"os"

	"github.com/api-meli-mutants/controllers/routes"
)

func main() {
	router := routes.NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	server := http.ListenAndServe(":"+port, router)
	log.Fatal(server)
}
