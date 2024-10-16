package main

import (
	db "flightticketservice/internal/db"
	"flightticketservice/utils"
	"log"
	"net/http"
	"time"

	// "flightticketservice/pkg/api"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type APIServer struct {
	listenAddr string
	listenPort string
	store      db.Storage
}

// NewAPIServer creates API server
func NewAPIServer(listenAddr, listenPort string, store db.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		listenPort: listenPort,
		store:      store,
	}
}

// Run launches http server
func (s *APIServer) Run() {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/api/v1/flights", GetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights/search", GetFlightsByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", GetFlightInfo).Methods("GET")

	r.HandleFunc("/api/v1/passengers", GetPassengers).Methods("GET")
	r.HandleFunc("/api/v1/passengers/{id}", GetPassengerByID).Methods("GET")
	r.HandleFunc("/api/v1/passengers/create", s.HandleCreatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/update", UpdatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/delete ", DeletePassenger).Methods("DELETE")

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
