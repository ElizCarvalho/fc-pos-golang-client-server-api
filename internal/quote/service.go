package quote

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ExternalAPIResponse represents the response from the external API
type ExternalAPIResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetQuote(ctx context.Context) (*Quote, error) {
	apiResponse, err := s.callExternalAPI(ctx)
	if err != nil {
		return nil, err
	}

	bid, err := ParseBid(apiResponse.USDBRL.Bid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse bid: %w", err)
	}

	quote := &Quote{
		Bid:       bid,
		Timestamp: time.Now(),
	}

	return quote, nil
}

func (s *Service) callExternalAPI(ctx context.Context) (*ExternalAPIResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call external API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to call external API: %s", resp.Status)
	}

	var apiResponse ExternalAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &apiResponse, nil
}
