package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ElizCarvalho/fc-pos-golang-client-server-api/internal/database"
	"github.com/ElizCarvalho/fc-pos-golang-client-server-api/internal/quote"
)

type Server struct {
	quoteService *quote.Service
	quoteRepo    *quote.Repository
	db           *sql.DB
}

func NewServer() (*Server, error) {
	log.Println("Connecting to database...")
	db, err := database.NewConnection()
	if err != nil {
		return nil, err
	}

	log.Println("Creating table...")
	err = database.CreateTable(db)
	if err != nil {
		return nil, err
	}

	quoteRepo := quote.NewRepository(db)
	quoteService := quote.NewService(quoteRepo)

	log.Println("Server initialized successfully")
	return &Server{
		quoteService: quoteService,
		quoteRepo:    quoteRepo,
	}, nil
}

func (s *Server) SetupRoutes() {
	http.HandleFunc("/healthcheck", s.healthCheckHandler)
	http.HandleFunc("/cotacao", s.cotacaoHandler)
}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am alive! Let's go!"))
}

func (s *Server) cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	// Context with 200ms timeout for external API call
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	quote, err := s.quoteService.GetQuote(ctx)
	if err != nil {
		log.Printf("Error getting quote: %v\n", err)
		http.Error(w, "Error getting quote", http.StatusInternalServerError)
		return
	}

	// Context with 10ms timeout for database save
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	err = s.quoteRepo.SaveQuote(ctxDB, quote.Bid)
	if err != nil {
		log.Printf("Error saving quote: %v\n", err)
	}

	response := map[string]float64{
		"bid": quote.Bid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) Start(port string) error {
	fmt.Println("Server is running on port", port)
	return http.ListenAndServe(":"+port, nil)
}

func (s *Server) Close() error {
	return s.db.Close()
}
