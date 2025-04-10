package main

import (
	"encoding/json"
	t "flightticketservice/pkg/booking"
	f "flightticketservice/pkg/flights"
	p "flightticketservice/pkg/passenger"
	"net/http"
	"os"
	"time"

	"flightticketservice/utils"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

// Flights

// handleGetFlights handles requests for getting list of flights.
func (s *APIServer) handleGetFlights(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetFlights called")

	flights, err := s.flights.GetFlights()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, flights)
}

// handleGetFlightByParams handles requests for getting list of flights by parameters.
func (s *APIServer) handleGetFlightByParams(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetFlightsByParams called")

	query := r.URL.Query()
	origin := query.Get("origin")
	destination := query.Get("destination")

	departure, err := time.Parse(time.RFC3339, query.Get("departure"))
	if err != nil {
		utils.ErrorLog.Printf("Error in time format: %v", err)
		http.Error(w, "Invalid departure time format. Use RFC3339.", http.StatusBadRequest)
		return
	}
	arrival, err := time.Parse(time.RFC3339, query.Get("arrival"))
	if err != nil {
		utils.ErrorLog.Printf("Error in time format: %v", err)
		http.Error(w, "Invalid arrival time format. Use RFC3339.", http.StatusBadRequest)
		return
	}

	searchParams := f.SearchParams{
		Origin:      origin,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
	}

	flights, err := s.flights.GetFlightsByParams(searchParams)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, flights)
}

// handleGetFlightByID handles requests for getting flight info.
func (s *APIServer) handleGetFlightByID(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetFlightInfo called")

	vars := mux.Vars(r)
	flightID := vars["id"]

	flight, err := s.flights.GetFlightByID(flightID)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flight: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, flight)
}

// handleCreateFlight handles requests for creating flight.
func (s *APIServer) handleCreateFlight(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CreateFlight called")

	createFlightReq := new(f.CreateFlightReq)
	if err := json.NewDecoder(r.Body).Decode(createFlightReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode flight data")
		return
	}

	newFlight := f.NewFlight(
		createFlightReq.Airline,
		createFlightReq.Origin,
		createFlightReq.Destination,
		createFlightReq.Departure,
		createFlightReq.Arrival,
		createFlightReq.Price,
	)

	utils.InfoLog.Println("new flight: ", newFlight, " created")

	if err := s.flights.CreateFlight(newFlight); err != nil {
		utils.ErrorLog.Printf("Error in CreateFlight: %v", err)
		return
	}

	WriteJSON(w, http.StatusCreated, "Flight created")
}

// handleUpdateFlight handles requests for updating flight.
func (s *APIServer) handleUpdateFlight(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("UpdateFlight called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	if passengerID == "" {
		utils.ErrorLog.Printf("Missing required parameters in UpdateFlight query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	createFlightReq := new(f.CreateFlightReq)
	if err := json.NewDecoder(r.Body).Decode(createFlightReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode flight data")
		return
	}

	newFlight := f.NewFlight(
		createFlightReq.Airline,
		createFlightReq.Origin,
		createFlightReq.Destination,
		createFlightReq.Departure,
		createFlightReq.Arrival,
		createFlightReq.Price,
	)

	utils.InfoLog.Println("Flight: ", newFlight, " updated")

	if err := s.flights.UpdateFlight(passengerID, newFlight); err != nil {
		utils.ErrorLog.Printf("Error in UpdateFlight: %v", err)
		return
	}

	WriteJSON(w, http.StatusOK, "Flight updated")
}

// handleDeleteFlight handles requests for deleting flight.
func (s *APIServer) handleDeleteFlight(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("DeleteFlight called")

	vars := mux.Vars(r)
	flightID := vars["id"]

	if flightID == "" {
		utils.ErrorLog.Printf("flight id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	if err := s.flights.DeleteFlight(flightID); err != nil {
		utils.ErrorLog.Printf("Error in DeleteFlight: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Flight deleted")
}

// Tickets

// handleBookTicket handles requests for booking flight.
func (s *APIServer) handleBookTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("BookTicket called")

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)
	ticketID := queryParams.Get("ticketID")
	passengerID := queryParams.Get("passengerID")
	additionalInfo := queryParams.Get("additionalInfo")

	if passengerID == "" || ticketID == "" {
		utils.ErrorLog.Printf("Missing required parameters in BookTicket query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	err := s.tickets.BookTicket(ticketID, passengerID, additionalInfo)

	if err != nil {
		utils.ErrorLog.Printf("Error in BookTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket booked")
}

// handleCheckInOnline handles requests for online registration.
func (s *APIServer) handleCheckInOnline(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CheckInOnline called")

	WriteJSON(w, http.StatusOK, "")
}

// handleChangeTicket handles requests for changing tickets.
func (s *APIServer) handleChangeTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("ChangeTicket called")

	vars := mux.Vars(r)
	ticketID := vars["id"]
	flightID := r.URL.Query().Get("flightID")

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)

	if ticketID == "" {
		utils.ErrorLog.Printf("Missing required parameters in ChangeTicket query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	err := s.tickets.ChangeFlight(ticketID, flightID)

	if err != nil {
		utils.ErrorLog.Printf("Error in ChangeTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket changed")
}

// handleCancelTicket handles requests for ticket cancellation.
func (s *APIServer) handleCancelTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CancelTicket called")

	vars := mux.Vars(r)
	ticketID := vars["id"]

	if ticketID == "" {
		utils.ErrorLog.Printf("ticket id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	err := s.tickets.CancelTicket(ticketID)

	if err != nil {
		utils.ErrorLog.Printf("Error in CancelTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket Cancelled")
}

// handleGetTickets handles requests for getting list of tickets.
func (s *APIServer) handleGetTickets(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetTickets called")

	tickets, err := s.tickets.GetTickets()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving tickets: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, tickets)
}

// handleGetTicketByID handles requests for getting ticket info.
func (s *APIServer) handleGetTicketByID(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetTicketInfo called")

	vars := mux.Vars(r)
	ticketID := vars["id"]

	tickets, err := s.tickets.GetTicketByID(ticketID)
	if err != nil {
		utils.ErrorLog.Printf("Error receiving tickets: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, tickets)
}

// handleUpdateTicket handles requests for updating ticket info
func (s *APIServer) handleUpdateTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("UpdateTicket called")

	vars := mux.Vars(r)
	ticketID := vars["id"]

	if ticketID == "" {
		utils.ErrorLog.Printf("Missing required parameters in UpdateTicket query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	createTicketReq := new(t.CreateTicketReq)
	if err := json.NewDecoder(r.Body).Decode(createTicketReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode ticket data")
		return
	}

	newTicket := t.CreateNewTicket(
		createTicketReq.FlightID,
		createTicketReq.PassengerID,
		"updated", // status
		createTicketReq.SeatNumber,
		createTicketReq.AdditionalInfo,
		createTicketReq.DepartureTime,
		createTicketReq.ArrivalTime,
	)

	utils.InfoLog.Println("Flight: ", newTicket, " updated")

	if err := s.tickets.UpdateTicket(ticketID, newTicket); err != nil {
		utils.ErrorLog.Printf("Error in UpdateTicket: %v", err)
		return
	}

	WriteJSON(w, http.StatusOK, "Flight updated")
}

// handleCreateTicket handles requests for updating ticket info
func (s *APIServer) handleCreateTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CreateTicket called")

	createTicketReq := new(t.CreateTicketReq)
	if err := json.NewDecoder(r.Body).Decode(createTicketReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode passenger data")
		return
	}

	newTicket := t.CreateNewTicket(
		createTicketReq.FlightID,
		createTicketReq.PassengerID,
		"created", // status
		createTicketReq.SeatNumber,
		createTicketReq.AdditionalInfo,
		createTicketReq.DepartureTime,
		createTicketReq.ArrivalTime,
	)

	utils.InfoLog.Println("New ticket: ", newTicket, " created")

	if err := s.tickets.CreateTicket(newTicket); err != nil {
		utils.ErrorLog.Printf("Error in CreateTicket: %v", err)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket created")
}

// handleUpdateTicket handles requests for deleting ticket info
func (s *APIServer) handleDeleteTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("DeleteTicket called")

	vars := mux.Vars(r)
	ticketID := vars["id"]

	if ticketID == "" {
		utils.ErrorLog.Printf("flight id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	if err := s.tickets.DeleteTicket(ticketID); err != nil {
		utils.ErrorLog.Printf("Error in DeleteTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSON(w, http.StatusOK, "Ticket deleted")
}

// Passengers

// handleGetPassengers handles requests for getting list of passengers.
func (s *APIServer) handleGetPassengers(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetPassengers called")

	passengers, err := s.store.GetPassengers()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving passengers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, passengers)
}

// handleGetPassengerByID handles requests for getting passenger by id.
func (s *APIServer) handleGetPassengerByID(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetPassengerByID called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	passenger, err := s.store.GetPassengerByID(passengerID)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving passenger: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, passenger)
}

// handleCreatePassenger handles requests for creating passenger.
func (s *APIServer) handleCreatePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CreatePassenger called")

	createPassengerReq := new(p.CreatePassengerReq)
	if err := json.NewDecoder(r.Body).Decode(createPassengerReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode passenger data")
		return
	}

	newPassenger, err := p.NewPassenger(
		createPassengerReq.FirstName,
		createPassengerReq.LastName,
		createPassengerReq.Email,
		createPassengerReq.Password,
	)

	if err != nil {
		utils.ErrorLog.Printf("Error in CreatePassenger: %v", err)
		return
	}

	utils.InfoLog.Println("new passenger: ", newPassenger, " created")

	if err := s.store.CreatePassenger(newPassenger); err != nil {
		utils.ErrorLog.Printf("Error in CreatePassenger: %v", err)
		return
	}

	WriteJSON(w, http.StatusCreated, "Passenger created")
}

// handleUpdatePassenger handles requests for updating passenger.
func (s *APIServer) handleUpdatePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("UpdatePassenger called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	if passengerID == "" {
		utils.ErrorLog.Printf("Missing required parameters in UpdatePassenger query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)

	createPassengerReq := new(p.CreatePassengerReq)
	if err := json.NewDecoder(r.Body).Decode(createPassengerReq); err != nil {
		utils.ErrorLog.Fatal("Cannot decode passenger data")
		return
	}

	newPassenger, err := p.NewPassenger(
		createPassengerReq.FirstName,
		createPassengerReq.LastName,
		createPassengerReq.Email,
		createPassengerReq.Password,
	)
	if err != nil {
		utils.ErrorLog.Printf("Error in UpdatePassenger: %v", err)
		return
	}
	utils.InfoLog.Println("Passenger: ", newPassenger, " updated")

	if err := s.store.UpdatePassenger(passengerID, newPassenger); err != nil {
		utils.ErrorLog.Printf("Error in UpdatePassenger: %v", err)
		return
	}

	WriteJSON(w, http.StatusOK, "Passenger updated")
}

// handleDeletePassenger handles requests for deleting passenger.
func (s *APIServer) handleDeletePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("DeletePassenger called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	if passengerID == "" {
		utils.ErrorLog.Printf("passenger id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	if err := s.store.DeletePassenger(passengerID); err != nil {
		utils.ErrorLog.Printf("Error in DeletePassenger: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Passenger deleted")
}

// handleGetPassengerTickets handles requests for getting list of tickets for passenger id.
func (s *APIServer) handleGetPassengerTickets(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetPassengerTickets called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	tickets, err := s.tickets.GetPassengerTickets(passengerID)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving tickets for passenger: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, tickets)
}

// // generateID generates unique ID.
// func generateID() string {
// 	return uuid.New().String()
// }

// WriteJSON writes response to JSON
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		utils.ErrorLog.Printf("Error in encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func createJWT(passenger *p.Passenger) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":    15000,
		"passengerNum": passenger.Number,
	}

	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
