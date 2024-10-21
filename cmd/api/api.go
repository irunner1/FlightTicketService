package main

import (
	f "flightticketservice/pkg/flights"
	p "flightticketservice/pkg/passenger"
	"flightticketservice/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// APIServer collects service settings and storage
type APIServer struct {
	listenAddr string
	listenPort string
	store      p.Storage
	flights    f.FlightService
}

// NewAPIServer creates API server
func NewAPIServer(listenAddr, listenPort string, store p.Storage, flightsStore f.FlightService) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		listenPort: listenPort,
		store:      store,
		flights:    flightsStore,
	}
}

// Run launches http server
func (s *APIServer) Run() {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/api/v1/flights", s.handleGetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights/search", s.handleGetFlightByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", s.handleGetFlightByID).Methods("GET")
	r.HandleFunc("/api/v1/flights/create", s.handleCreateFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{id}/update", s.handleUpdateFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{id}/delete", s.handleDeleteFlight).Methods("DELETE")

	r.HandleFunc("/api/v1/passengers", s.handleGetPassengers).Methods("GET")
	r.HandleFunc("/api/v1/passengers/{id}", s.handleGetPassengerByID).Methods("GET")
	r.HandleFunc("/api/v1/passengers/create", s.handleCreatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/update", s.handleUpdatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/delete ", s.handleDeletePassenger).Methods("DELETE")

	r.HandleFunc("/api/v1/tickets", GetTickets).Methods("GET")
	r.HandleFunc("/api/v1/tickets/{id}", GetTicketInfo).Methods("GET")

	r.HandleFunc("/api/v1/tickets/book", BookTicket).Methods("POST")
	r.HandleFunc("/api/v1/checkin", CheckInOnline).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{ticketID}/change", ChangeTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{ticketID}/cancel", CancelTicket).Methods("POST")

	log.Println("JSON API server running on port: ", s.listenAddr)

	srv := &http.Server{
		Handler:      r,
		Addr:         s.listenAddr + ":" + s.listenPort,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	utils.InfoLog.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		utils.ErrorLog.Fatal("Server failed to start:", err)
	}
}
