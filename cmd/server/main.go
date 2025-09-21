package main

import (
	"log"

	"github.com/ElizCarvalho/fc-pos-golang-client-server-api/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal("Error creating server: ", err)
	}
	defer srv.Close()

	srv.SetupRoutes()
	log.Fatal(srv.Start("8080"))
}
