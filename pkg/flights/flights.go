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
	GetFlights() ([]Flight, error)
	GetFlightsByParams(params SearchParams) ([]Flight, error)
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
	flights []Flight
}

// NewFlightService creates new instance of FlightService with predefined flights.
func NewFlightService() FlightService {
	sampleFlights := []Flight{
		{
			ID:          "FL001",
			Airline:     "Global Airlines",
			Origin:      "New York",
			Destination: "London",
			Departure:   time.Date(2023, time.April, 15, 8, 0, 0, 0, time.UTC),
			Arrival:     time.Date(2023, time.April, 15, 20, 0, 0, 0, time.UTC),
			Price:       450.00,
		},
		{
			ID:          "FL002",
			Airline:     "Sky Travelers",
			Origin:      "San Francisco",
			Destination: "Tokyo",
			Departure:   time.Date(2023, time.April, 20, 11, 0, 0, 0, time.UTC),
			Arrival:     time.Date(2023, time.April, 21, 5, 0, 0, 0, time.UTC),
			Price:       800.00,
		},
	}

	return &flightServiceImpl{
		flights: sampleFlights,
	}
}

// GetFlights returns list of flights.
func (s *flightServiceImpl) GetFlights() ([]Flight, error) {
	return s.flights, nil
}

// GetFlights returns list of flights by parameters.
func (s *flightServiceImpl) GetFlightsByParams(params SearchParams) ([]Flight, error) {
	var matchingFlights = []Flight{}
	for _, flight := range s.flights {
		if params.Origin != "" && flight.Origin != params.Origin {
			continue
		}
		if params.Destination != "" && flight.Destination != params.Destination {
			continue
		}
		if !params.Departure.IsZero() && !flight.Departure.Equal(params.Departure) {
			continue
		}
		if !params.Arrival.IsZero() && !flight.Arrival.Equal(params.Arrival) {
			continue
		}
		matchingFlights = append(matchingFlights, flight)
	}
	if len(matchingFlights) == 0 {
		return nil, errors.New("no flights found matching the search criteria")
	}
	return matchingFlights, nil
}

// GetFlightByID gets info about flight by id.
func (s *flightServiceImpl) GetFlightByID(flightID string) (Flight, error) {
	for _, flight := range s.flights {
		if flight.ID == flightID {
			return flight, nil
		}
	}
	return Flight{}, errors.New("flight not found")
}