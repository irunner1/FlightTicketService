package main

import (
	"encoding/json"
	t "flightticketservice/pkg/booking"
	f "flightticketservice/pkg/flights"
	p "flightticketservice/pkg/passenger"
	"flightticketservice/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// APIServer collects service settings and storage
type APIServer struct {
	listenAddr string
	listenPort string
	store      p.Storage
	flights    f.FlightService
	tickets    t.BookingService
}

// NewAPIServer creates API server
func NewAPIServer(
	listenAddr,
	listenPort string,
	store p.Storage,
	flightsStore f.FlightService,
	ticketStore t.BookingService,
) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		listenPort: listenPort,
		store:      store,
		flights:    flightsStore,
		tickets:    ticketStore,
	}
}

// Run launches http server
func (s *APIServer) Run() {
	r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/api/v1/login", s.handleLogin).Methods("POST")

	r.HandleFunc("/api/v1/flights", s.handleGetFlights).Methods("GET")
	r.HandleFunc("/api/v1/flights/search", s.handleGetFlightByParams).Methods("GET")
	r.HandleFunc("/api/v1/flights/{id}", s.handleGetFlightByID).Methods("GET")
	r.HandleFunc("/api/v1/flights/create", s.handleCreateFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{id}/update", s.handleUpdateFlight).Methods("POST")
	r.HandleFunc("/api/v1/flights/{id}/delete", s.handleDeleteFlight).Methods("DELETE")
	
	r.HandleFunc("/api/v1/passengers", s.handleGetPassengers).Methods("GET")
	r.HandleFunc("/api/v1/passengers/{id}", withJWTAuth(s.handleGetPassengerByID, s.store)).Methods("GET")
	r.HandleFunc("/api/v1/passengers/create", s.handleCreatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/update", s.handleUpdatePassenger).Methods("POST")
	r.HandleFunc("/api/v1/passengers/{id}/delete ", s.handleDeletePassenger).Methods("DELETE")
	
	r.HandleFunc("/api/v1/tickets", s.handleGetTickets).Methods("GET")
	r.HandleFunc("/api/v1/tickets/{id}", s.handleGetTicketByID).Methods("GET")
	r.HandleFunc("/api/v1/tickets/passenger/{id}", s.handleGetPassengerTickets).Methods("GET")

	r.HandleFunc("/api/v1/tickets/book", s.handleBookTicket).Methods("POST")
	r.HandleFunc("/api/v1/checkin", s.handleCheckInOnline).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{id}/change", s.handleChangeTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{id}/cancel", s.handleCancelTicket).Methods("POST")

	r.HandleFunc("/api/v1/tickets/create", s.handleCreateTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{id}/update", s.handleUpdateTicket).Methods("POST")
	r.HandleFunc("/api/v1/tickets/{id}/delete", s.handleDeleteTicket).Methods("DELETE")

	utils.InfoLog.Println("JSON API server running on port: ", s.listenAddr)

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

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) {
	var req p.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	pass, err := s.store.GetPassengerByEmail(req.Email)
	if err != nil {
		utils.ErrorLog.Println("not found")
		return
	}

	if !pass.ValidPassword(req.Password) {
		utils.ErrorLog.Println("not authenticated")
		return
	}

	token, err := createJWT(pass)
	fmt.Println(token)
	if err != nil {
		return
	}

	resp := p.LoginResponse{
		Token: token,
		Email: pass.Email,
	}

	WriteJSON(w, http.StatusOK, resp)
}

func withJWTAuth(handlerFunc http.HandlerFunc, s p.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.InfoLog.Println("calling JWT auth middleware")

		tokenString := r.Header.Get("Authorization")

		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}
		passengerID := mux.Vars(r)["id"]
		passenger, err := s.GetPassengerByID(passengerID)
		if err != nil {
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)

		if passenger.Number != int64(claims["passengerNum"].(float64)) {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusForbidden, APIError{Error: "permission denied"})
}

// APIError creates error
type APIError struct {
	Error string `json:"error"`
}
