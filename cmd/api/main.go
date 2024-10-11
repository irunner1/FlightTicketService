package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"flightticketservice/pkg/api"
	"flightticketservice/utils"

	_ "flightticketservice/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Flight Ticket Service
// @version 1.0
// @description API Server for booking flight tickets

// @host localhost:8010
// @Basepath /

func main() {
	if err := godotenv.Load(); err != nil {
		utils.InfoLog.Print("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	utils.InfoLog.Printf("loaded env {'host': %s, 'port': %s}", host, port)

	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/api/v1/flights", api.GetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights/search", api.GetFlightsByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", api.GetFlightInfo).Methods("GET")

	r.HandleFunc("/api/v1/tickets", api.GetTickets).Methods("GET")
	r.HandleFunc("/api/v1/tickets/{id}", api.GetTicketInfo).Methods("GET")

	r.HandleFunc("/api/v1/tickets/book", api.BookTicket).Methods("POST")
	r.HandleFunc("/api/v1/checkin", api.CheckInOnline).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{ticketID}/change", api.ChangeTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{ticketID}/cancel", api.CancelTicket).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	utils.InfoLog.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		utils.ErrorLog.Fatal("Server failed to start:", err)
	}
}
