package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type QuoteResponse struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	quote, err := getQuote(ctx)
	if err != nil {
		log.Printf("Error getting quote: %v\n", err)
		return
	}

	err = saveQuoteToFile(quote.Bid)
	if err != nil {
		log.Printf("Error saving quote to file: %v\n", err)
		return
	}
	fmt.Printf("Quote saved -> Dólar: %.4f\n", quote.Bid)
}

func getQuote(ctx context.Context) (*QuoteResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get quote: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var quote QuoteResponse
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	return &quote, nil
}

func saveQuoteToFile(bid float64) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	content := fmt.Sprintf("Dólar: %.4f", bid)
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}
