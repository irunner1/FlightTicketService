package api

import (
	"encoding/json"
	"flightticketservice/pkg/booking"
	"flightticketservice/pkg/flights"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

// GetFlights handles requests for getting list of flights.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFlights called")
	flightSer := flights.NewFlightService()
	flights, err := flightSer.GetFlights()

	if err != nil {
		log.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightsByParams handles requests for getting list of flights by parameters.
// GET /flights_by_params?origin=Moscow&destination=Istanbul&departure=2024-04-15T10:00:00Z&arrival=2024-04-15T13:30:00Z
// GET /flights_by_params?origin=New%20York&destination=London&departure=2023-04-15T08:00:00Z&arrival=2023-04-15T20:00:00Z
func GetFlightsByParams(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFlightsByParams called")

	query := r.URL.Query()
	origin := query.Get("origin")
	destination := query.Get("destination")

	departure, err := time.Parse(time.RFC3339, query.Get("departure"))
	if err != nil {
		log.Printf("Error in time format: %v", err)
		http.Error(w, "Invalid departure time format. Use RFC3339.", http.StatusBadRequest)
		return
	}
	arrival, err := time.Parse(time.RFC3339, query.Get("arrival"))
	if err != nil {
		log.Printf("Error in time format: %v", err)
		http.Error(w, "Invalid arrival time format. Use RFC3339.", http.StatusBadRequest)
		return
	}

	searchParams := flights.SearchParams{
		Origin:      origin,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
	}

	flightSer := flights.NewFlightService()
	flights, err := flightSer.GetFlightsByParams(searchParams)

	if err != nil {
		log.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightInfo handles requests for getting flight info.
func GetFlightInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("GetFlightInfo called")

	vars := mux.Vars(r)
	flightID := vars["id"]

	flightSer := flights.NewFlightService()
	flight, err := flightSer.GetFlightByID(flightID)
	if err != nil {
		log.Printf("Error receiving flight: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flight)
}

// BookTicket handles requests for booking flight.
// http://localhost:8010/api/v1/book_ticket?passengerID=passenger1&seatNumber=12A&additionalInfo=V&flightID=flight1
func BookTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("BookTicket called")

	queryParams := r.URL.Query()
	log.Println(queryParams)
	passengerID := queryParams.Get("passengerID")
	seatNumber := queryParams.Get("seatNumber")
	additionalInfo := queryParams.Get("additionalInfo")
	flightID := queryParams.Get("flightID")

	ticketID := generateTicketID()

	if passengerID == "" || seatNumber == "" || flightID == "" {
		log.Printf("Missing required parameters in BookTicket query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	newTicket := booking.Ticket{
		ID:             ticketID,
		FlightID:       flightID,
		PassengerID:    passengerID,
		BookingTime:    time.Now(),
		DepartureTime:  time.Now(),
		ArrivalTime:    time.Now(),
		Status:         "booked",
		SeatNumber:     seatNumber,
		AdditionalInfo: additionalInfo,
	}
	log.Println(newTicket)

	bookingSer := booking.NewBookingService()
	err := bookingSer.BookTicket(&newTicket)

	if err != nil {
		log.Printf("Error in BookTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Ticket booked")
}

// CheckInOnline handles requests for online registration.
func CheckInOnline(w http.ResponseWriter, r *http.Request) {
	log.Println("CheckInOnline called")

	w.WriteHeader(http.StatusOK)
}

// ChangeTicket handles requests for online registration.
func ChangeTicket(w http.ResponseWriter, r *http.Request) {
	log.Println("ChangeTicket called")

	bookingSer := booking.NewBookingService()
	log.Print(bookingSer)
	// _, err := bookingSer.BookTicket()

	// if err != nil {
	//  log.Printf("Error BookTicket: %v", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(flights)
	w.WriteHeader(http.StatusOK)
}

// CancelTicket handles requests for online registration.
func CancelTicket(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// generateTicketID generates unique ID for each ticket.
func generateTicketID() string {
	id := uuid.New()
	return id.String()
}
