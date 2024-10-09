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
	log.Printf("loaded env `host: %s, port: %s`", host, port)

	r := mux.NewRouter()

	// r.HandleFunc("/api/openapi.json", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "openapi.json")
	// }).Methods("GET")
	// r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("path/to/swaggerui/"))))

	r.HandleFunc("/api/v1/flights", GetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights_by_params", GetFlightsByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", GetFlightInfo).Methods("GET")
	r.HandleFunc("/api/v1/tickets", BookTicket).Methods("POST")
	r.HandleFunc("/api/v1/checkin", CheckInOnline).Methods("POST")

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
	flights, err := flightSer.GetFlights()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flights)
}

// GetFlightsByParams handles requests for getting list of flights by parameters.
// GET /flights?origin=Moscow&destination=Istanbul&departure=2024-04-15T10:00:00Z&arrival=2024-04-15T13:30:00Z
// GET /flights?origin=New%20York&destination=London&departure=2023-04-15T08:00:00Z&arrival=2023-04-15T20:00:00Z
func GetFlightsByParams(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	origin := query.Get("origin")
	destination := query.Get("destination")

	departure, err := time.Parse(time.RFC3339, query.Get("departure"))
	if err != nil {
		http.Error(w, "Invalid departure time format. Use RFC3339.", http.StatusBadRequest)
		return
	}
	arrival, err := time.Parse(time.RFC3339, query.Get("arrival"))
	if err != nil {
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
