package flights

import "time"

// Flight collects flight data.
// @Description Flight model for API response.
type Flight struct {
	ID          string    `json:"id"`
	Airline     string    `json:"airline"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	Price       float64   `json:"price"`
}

// SearchParams collects parameters for searching flights.
type SearchParams struct {
	Origin      string
	Destination string
	Departure   time.Time
	Arrival     time.Time
}

// CreateFlightReq collects info about flight for request.
type CreateFlightReq struct {
	Airline     string    `json:"airline"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	Price       float64   `json:"price"`
}

// NewFlight creates new flight by passed params
func NewFlight(airline, origin, destination string, departure, arrival time.Time, price float64) *Flight {
	return &Flight{
		Airline:     airline,
		Origin:      origin,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
		Price:       price,
	}
}
