package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	r := mux.NewRouter()

	r.HandleFunc("/api/flights", GetFlights).Methods("GET")
	r.HandleFunc("/api/flights/{id}", GetFlightInfo).Methods("GET")
	r.HandleFunc("/api/tickets", BookTicket).Methods("POST")
	r.HandleFunc("/api/checkin", CheckInOnline).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         host + ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// GetFlights handles requests for getting list of flights.
func GetFlights(w http.ResponseWriter, r *http.Request) {

}

// GetFlightInfo handles requests for getting flight info
func GetFlightInfo(w http.ResponseWriter, r *http.Request) {

}

// BookTicket handles requests for booking flight
func BookTicket(w http.ResponseWriter, r *http.Request) {

}

// CheckInOnline handles requests for online registration
func CheckInOnline(w http.ResponseWriter, r *http.Request) {

}
