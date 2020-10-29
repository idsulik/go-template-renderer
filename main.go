package main

import (
	"github.com/idsulik/go-template-renderer/server"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port :" + port)

	log.Fatal(server.New(port))
}
