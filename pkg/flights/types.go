package flights

import "time"

// FlightSearchParameters collects parameters for searching flights.
type FlightSearchParameters struct {
	DepartureDate string `json:"departure_date"`
	ArrivalDate   string `json:"arrival_date"`
	TicketType    string `json:"ticket_type"`
	Baggage       bool   `json:"baggage"`
}

// FlightInfo collects flight info.
type FlightInfo struct {
	FlightID      string `json:"flight_id"`
	BaggageAllow  bool   `json:"baggage_allow"`
	HandLuggage   bool   `json:"hand_luggage"`
	TicketReturns bool   `json:"ticket_returns"`
}

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

// CreateFlightRequest collects info about flight for request.
type CreateFlightRequest struct {
	Airline     string    `json:"airline"`
	Origin      string    `json:"origin"`
	Destination string    `json:"destination"`
	Departure   time.Time `json:"departure"`
	Arrival     time.Time `json:"arrival"`
	Price       float64   `json:"price"`
}

// CreateNewFlight creates new flight by passed params
func CreateNewFlight(airline, origin, destination string, departure, arrival time.Time, price float64) *Flight {
	return &Flight{
		ID:          "FL001",
		Airline:     airline,
		Origin:      origin,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
		Price:       price,
	}
}
