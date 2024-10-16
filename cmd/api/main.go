package main

import (
	"os"

	"github.com/joho/godotenv"

	db "flightticketservice/pkg/passenger"
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

	store, err := db.NewPostgresStore(os.Getenv("DB_USER"), os.Getenv("DB_PASS"))

	if err != nil {
		utils.ErrorLog.Fatal(err)
	}

	if err := store.Init(); err != nil {
		utils.ErrorLog.Fatal(err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	utils.InfoLog.Printf("loaded env {'host': %s, 'port': %s}", host, port)

	server := NewAPIServer(host, port, store)
	server.Run()
}
