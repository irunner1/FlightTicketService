package main

import (
	"encoding/json"
	"flightticketservice/pkg/booking"
	"flightticketservice/pkg/flights"
	p "flightticketservice/pkg/passenger"
	"net/http"
	"time"

	"flightticketservice/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// GetFlights handles requests for getting list of flights.
func GetFlights(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetFlights called")
	flightSer := flights.NewFlightService()
	flights, err := flightSer.GetFlights()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, flights)
}

// GetFlightsByParams handles requests for getting list of flights by parameters.
func GetFlightsByParams(w http.ResponseWriter, r *http.Request) {
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

	searchParams := flights.SearchParams{
		Origin:      origin,
		Destination: destination,
		Departure:   departure,
		Arrival:     arrival,
	}

	flightSer := flights.NewFlightService()
	flights, err := flightSer.GetFlightsByParams(searchParams)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, flights)
}

// GetFlightInfo handles requests for getting flight info.
func GetFlightInfo(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetFlightInfo called")

	vars := mux.Vars(r)
	flightID := vars["id"]

	flightSer := flights.NewFlightService()
	flight, err := flightSer.GetFlightByID(flightID)
	if err != nil {
		utils.ErrorLog.Printf("Error receiving flight: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, flight)
}

// BookTicket handles requests for booking flight.
func BookTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("BookTicket called")

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)
	passengerID := queryParams.Get("passengerID")
	seatNumber := queryParams.Get("seatNumber")
	additionalInfo := queryParams.Get("additionalInfo")
	flightID := queryParams.Get("flightID")

	ticketID := generateID()

	if passengerID == "" || seatNumber == "" || flightID == "" {
		utils.ErrorLog.Printf("Missing required parameters in BookTicket query")
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
	utils.InfoLog.Println("new ticked: ", newTicket, " created")

	bookingSer := booking.NewBookingService()
	err := bookingSer.BookTicket(&newTicket)

	if err != nil {
		utils.ErrorLog.Printf("Error in BookTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket booked")
}

// CheckInOnline handles requests for online registration.
func CheckInOnline(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CheckInOnline called")

	WriteJSON(w, http.StatusOK, "")
}

// ChangeTicket handles requests for changing tickets.
func ChangeTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("ChangeTicket called")

	vars := mux.Vars(r)
	ticketID := vars["ticketID"]
	flightID := r.URL.Query().Get("flightID")

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)

	if ticketID == "" {
		utils.ErrorLog.Printf("Missing required parameters in ChangeTicket query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	bookingSer := booking.NewBookingService()
	err := bookingSer.ChangeFlight(ticketID, flightID)

	if err != nil {
		utils.ErrorLog.Printf("Error in ChangeTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Flight id changed")
}

// CancelTicket handles requests for ticket cancellation.
func CancelTicket(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CancelTicket called")

	vars := mux.Vars(r)
	ticketID := vars["ticketID"]

	if ticketID == "" {
		utils.ErrorLog.Printf("ticket id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	bookingSer := booking.NewBookingService()
	err := bookingSer.CancelTicket(ticketID)

	if err != nil {
		utils.ErrorLog.Printf("Error in CancelTicket: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket Cancelled")
}

// GetTickets handles requests for getting list of tickets.
func GetTickets(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetTickets called")

	ticketSer := booking.NewBookingService()
	tickets, err := ticketSer.GetTickets()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flights: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, tickets)
}

// GetTicketInfo handles requests for getting ticket info.
func GetTicketInfo(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetTicketInfo called")

	vars := mux.Vars(r)
	ticketID := vars["id"]

	ticketSer := booking.NewBookingService()
	tickets, err := ticketSer.GetTicketByID(ticketID)
	if err != nil {
		utils.ErrorLog.Printf("Error receiving tickets: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, tickets)
}

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
		utils.ErrorLog.Printf("Error receiving flight: %v", err)
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

	newPassenger := p.NewPassenger(
		createPassengerReq.FirstName,
		createPassengerReq.LastName,
		createPassengerReq.Email,
		createPassengerReq.Password,
	)

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

	newPassenger := p.NewPassenger(
		createPassengerReq.FirstName,
		createPassengerReq.LastName,
		createPassengerReq.Email,
		createPassengerReq.Password,
	)
	utils.InfoLog.Println("new passenger: ", newPassenger, " created")

	if err := s.store.UpdatePassenger(passengerID, newPassenger); err != nil {
		utils.ErrorLog.Printf("Error in CreatePassenger: %v", err)
		return
	}

	WriteJSON(w, http.StatusOK, "Passenger updated")
}

// handleDeletePassenger handles requests for deleting passenger.
func (s *APIServer) handleDeletePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("DeletePassenger called")

	vars := mux.Vars(r)
	passengerID := vars["passengerID"]

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

// generateID generates unique ID.
func generateID() string {
	id := uuid.New()
	return id.String()
}

// WriteJSON writes response to JSON
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
