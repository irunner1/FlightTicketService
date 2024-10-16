package flights

import (
	"errors"
	"time"
)

// FlightService interface for working with flights.
type FlightService interface {
	GetFlights() ([]Flight, error)
	GetFlightsByParams(params SearchParams) ([]Flight, error)
	GetFlightByID(flightID string) (Flight, error)
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

// @Summary Get list of flights
// @Description get flights
// @Tags flights
// @Accept  json
// @Produce  json
// @Success 200 {array} Flight
// @Router /api/v1/flights [get]
func (s *flightServiceImpl) GetFlights() ([]Flight, error) {
	return s.flights, nil
}

// @Summary Search flights by parameters
// @Description Retrieves a list of flights filtered by the provided search parameters.
// @Tags flights
// @Accept json
// @Produce json
// @Param origin query string false "Origin location of the flight"
// @Param destination query string false "Destination location of the flight"
// @Param departure query string false "Departure date and time"
// @Param arrival query string false "Arrival date and time"
// @Success 200 {array} Flight
// @Failure 404 "No flights found matching the search criteria"
// @Router /api/v1/flights/search [get]
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

// @Summary Get flight by ID
// @Description Gets flight details for a specific flight ID.
// @Tags flights
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the flight"
// @Success 200 {object} Flight
// @Failure 404 "Flight not found"
// @Router /api/v1/flights/{id} [get]
func (s *flightServiceImpl) GetFlightByID(flightID string) (Flight, error) {
	for _, flight := range s.flights {
		if flight.ID == flightID {
			return flight, nil
		}
	}
	return Flight{}, errors.New("flight not found")
}
