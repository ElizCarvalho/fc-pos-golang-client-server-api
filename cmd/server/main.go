package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am alive! Let's go!"))
	})

	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello! This is the /quote endpoint"))
	})

	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
