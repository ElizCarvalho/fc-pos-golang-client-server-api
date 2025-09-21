package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Configurar rota de health check
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I am alive! Let's go!"))
	})

	// Configurar rota de cotacao
	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello! This is the /quote endpoint"))
	})

	// Iniciar servidor
	fmt.Println("Server is running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
