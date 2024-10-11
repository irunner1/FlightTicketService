package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"flightticketservice/pkg/api"
	"flightticketservice/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		utils.InfoLog.Print("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	utils.InfoLog.Printf("loaded env {'host': %s, 'port': %s}", host, port)

	r := mux.NewRouter()

	// r.HandleFunc("/api/openapi.json", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "openapi.json")
	// }).Methods("GET")
	// r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("path/to/swaggerui/"))))

	r.HandleFunc("/api/v1/flights", api.GetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights_by_params", api.GetFlightsByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", api.GetFlightInfo).Methods("GET")

	r.HandleFunc("/api/v1/book_ticket", api.BookTicket).Methods("POST")
	r.HandleFunc("/api/v1/checkin", api.CheckInOnline).Methods("POST")
	r.HandleFunc("/api/v1/change_ticket", api.ChangeTicket).Methods("POST")
	r.HandleFunc("/api/v1/cancel_ticket", api.CancelTicket).Methods("POST")

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
