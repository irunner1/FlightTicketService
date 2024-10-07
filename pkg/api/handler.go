package api

import (
	"encoding/json"
	"net/http"
)

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

// GetFlights handle requests on list of flights.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	response := make([]FlightInfo, 0)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetFlightInfo handle requests for getting info about flight.
func GetFlightInfo(w http.ResponseWriter, r *http.Request) {
	response := FlightInfo{}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// BookTicket handle requests for booking flight.
func BookTicket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

// CheckInOnline handle requests for registration on flight.
func CheckInOnline(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
