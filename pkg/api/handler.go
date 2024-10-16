package api

import (
	"encoding/json"
	"flightticketservice/pkg/booking"
	"flightticketservice/pkg/flights"
	"flightticketservice/pkg/passenger"
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

// GetPassengers handles requests for getting list of passengers.
func GetPassengers(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetPassengers called")
	PassengerSer := passenger.NewPassengerService()
	passengers, err := PassengerSer.GetPassengers()

	if err != nil {
		utils.ErrorLog.Printf("Error receiving passengers: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, passengers)
}

// GetPassengerByID handles requests for getting passenger by id.
func GetPassengerByID(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("GetPassengerByID called")

	vars := mux.Vars(r)
	passengerID := vars["id"]

	PassengerSer := passenger.NewPassengerService()
	passenger, err := PassengerSer.GetPassengerByID(passengerID)

	if err != nil {
		utils.ErrorLog.Printf("Error receiving flight: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	WriteJSON(w, http.StatusOK, passenger)
}

// CreatePassenger handles requests for creating passenger.
func CreatePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("CreatePassenger called")

	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)
	userName := queryParams.Get("name")
	userSurname := queryParams.Get("surname")
	email := queryParams.Get("email")
	password := queryParams.Get("password")

	passengerID := generateID()

	if userName == "" || userSurname == "" || email == "" || password == "" {
		utils.ErrorLog.Printf("Missing required parameters in CreatePassenger query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	newPassenger := passenger.Passenger{
		ID:        passengerID,
		FirstName: userName,
		LastName:  userSurname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	utils.InfoLog.Println("new ticked: ", newPassenger, " created")

	passengerSer := passenger.NewPassengerService()
	err := passengerSer.CreatePassenger(&newPassenger)

	if err != nil {
		utils.ErrorLog.Printf("Error in CreatePassenger: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusCreated, "Passenger created")
}

// UpdatePassenger handles requests for updating passenger.
func UpdatePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("UpdatePassenger called")

	vars := mux.Vars(r)
	passengerID := vars["passengerID"]
	queryParams := r.URL.Query()
	utils.InfoLog.Println(queryParams)
	name := queryParams.Get("name")
	surname := queryParams.Get("surname")
	email := queryParams.Get("email")
	password := queryParams.Get("password")

	newPassenger := passenger.Passenger{
		ID:        generateID(),
		FirstName: name,
		LastName:  surname,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().Add(-24 * time.Hour),
	}

	if passengerID == "" {
		utils.ErrorLog.Printf("Missing required parameters in UpdatePassenger query")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	passengerSer := passenger.NewPassengerService()
	err := passengerSer.UpdatePassenger(passengerID, &newPassenger)

	if err != nil {
		utils.ErrorLog.Printf("Error in UpdatePassenger: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteJSON(w, http.StatusOK, "Passenger updated")
}

// DeletePassenger handles requests for deleting passenger.
func DeletePassenger(w http.ResponseWriter, r *http.Request) {
	utils.InfoLog.Println("DeletePassenger called")

	vars := mux.Vars(r)
	passengerID := vars["passengerID"]

	if passengerID == "" {
		utils.ErrorLog.Printf("passenger id is empty")
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	passengerSer := passenger.NewPassengerService()
	err := passengerSer.DeletePassenger(passengerID)

	if err != nil {
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
