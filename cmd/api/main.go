package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"flightticketservice/pkg/flights"
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
	flightSer := flights.NewFlightService()
	searchParams := flights.SearchParams{
		Origin:      "Moscow",
		Destination: "Istanbul",
		Departure:   time.Date(2024, time.April, 15, 10, 0, 0, 0, time.UTC),
		Arrival:     time.Date(2024, time.April, 15, 13, 30, 0, 0, time.UTC),
	}
	flights, err := flightSer.GetFlights(searchParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightInfo handles requests for getting flight info
func GetFlightInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flightID := vars["id"]

	flightSer := flights.NewFlightService()
	flight, err := flightSer.GetFlightByID(flightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flight)
}

// BookTicket handles requests for booking flight
func BookTicket(w http.ResponseWriter, r *http.Request) {

}

// CheckInOnline handles requests for online registration
func CheckInOnline(w http.ResponseWriter, r *http.Request) {

}
