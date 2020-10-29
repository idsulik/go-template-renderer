package main

import (
	"github.com/idsulik/go-template-renderer/server"
	"log"
	"os"
	"strconv"
)

func main() {
	port := 8080

	if portStr, found := os.LookupEnv("PORT"); found {
		var err error
		port, err = strconv.Atoi(portStr)

		if err != nil {
			panic(err)
		}
	}

	log.Fatal(server.New(port))
}