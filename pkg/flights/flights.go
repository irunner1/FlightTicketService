package flights

import (
	"errors"
	"time"
)

// Flight collects flight data.
type Flight struct {
	ID          string    `json:"id"`
	Airline     string    `json:"airline"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	Price       float64   `json:"price"`
}

// FlightService interface for working with flights.
type FlightService interface {
	GetFlights(params SearchParams) ([]Flight, error)
	GetFlightByID(flightID string) (Flight, error)
}

// SearchParams collects parameters for searching flights.
type SearchParams struct {
	Origin      string
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

// flightServiceImpl structure implements interface FlightService.
type flightServiceImpl struct {
}

// NewFlightService creates new instance of FlightService.
func NewFlightService() FlightService {
	return &flightServiceImpl{}
}

// GetFlights returns list of flights.
func (s *flightServiceImpl) GetFlights(params SearchParams) ([]Flight, error) {
	return []Flight{}, nil
}

// GetFlightByID gets info about flight by id.
func (s *flightServiceImpl) GetFlightByID(flightID string) (Flight, error) {
	if flightID == "" {
		return Flight{}, errors.New("flight ID cannot be empty")
	}
	return Flight{
		ID:          flightID,
		Airline:     "Example Airline",
		Origin:      "City A",
		Destination: "City B",
		Departure:   time.Now(),
		Arrival:     time.Now().Add(2 * time.Hour),
		Price:       199.99,
	}, nil
}
