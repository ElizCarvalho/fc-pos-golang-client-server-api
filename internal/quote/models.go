package quote

import (
	"errors"
	"strconv"
	"time"
)

type Quote struct {
	ID        int       `json:"id"`
	Bid       float64   `json:"bid"`
	Timestamp time.Time `json:"timestamp"`
}

func ParseBid(bidString string) (float64, error) {
	bidConverted, err := strconv.ParseFloat(bidString, 64)
	if err != nil {
		return 0, errors.New("invalid bid format")
	}
	if bidConverted < 0 {
		return 0, errors.New("bid cannot be negative")
	}
	return bidConverted, nil
}

func (q *Quote) Validate() error {
	if q.Bid <= 0 {
		return errors.New("bid cannot be negative")
	}
	return nil
}
