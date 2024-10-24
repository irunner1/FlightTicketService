package main

import (
	"os"

	"github.com/joho/godotenv"

	"flightticketservice/pkg/booking"
	db "flightticketservice/pkg/database"
	"flightticketservice/pkg/flights"
	"flightticketservice/pkg/passenger"

	"flightticketservice/utils"

	_ "flightticketservice/docs"
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

	store, err := db.ConnectDB(os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	if err != nil {
		utils.ErrorLog.Fatal(err)
	}

	passengerStore := passenger.NewPostgresStore(store)
	if err := passengerStore.Init(); err != nil {
		utils.ErrorLog.Fatal(err)
	}

	flightsStore := flights.NewFlightsStore(store)
	if err := flightsStore.Init(); err != nil {
		utils.ErrorLog.Fatal(err)
	}

	ticketStore := booking.NewBookingStore(store)
	if err := ticketStore.Init(); err != nil {
		utils.ErrorLog.Fatal(err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	utils.InfoLog.Printf("loaded env {'host': %s, 'port': %s}", host, port)

	server := NewAPIServer(host, port, passengerStore, flightsStore, ticketStore)
	server.Run()
}
