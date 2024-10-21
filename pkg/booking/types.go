package booking

import "time"

// Ticket collects info about ticket.
type Ticket struct {
	ID             string    `json:"id"`
	FlightID       string    `json:"flight_id"`
	PassengerID    string    `json:"passenger_id"`
	BookingTime    time.Time `json:"booking_time"`
	DepartureTime  time.Time `json:"departure_time"`
	ArrivalTime    time.Time `json:"arrival_time"`
	Status         string    `json:"status"` // "booked", "cancelled", "confirmed"
	SeatNumber     string    `json:"seat_number"`
	AdditionalInfo string    `json:"additional_info"`
}

// CreateTicketReq collects info about ticket for request.
type CreateTicketReq struct {
	FlightID       string `json:"flight_id"`
	PassengerID    string `json:"passenger_id"`
	AdditionalInfo string `json:"additional_info"`
}

// CreateNewTicket creates new ticket by passed params
func CreateNewTicket(flightID, passengerID, additionalInfo string) *Ticket {
	return &Ticket{
		FlightID:       flightID,
		PassengerID:    passengerID,
		BookingTime:    time.Now().UTC(),
		DepartureTime:  time.Now().Add(24 * time.Hour),
		ArrivalTime:    time.Now().Add(30 * time.Hour),
		Status:         "booked",
		SeatNumber:     "12A",
		AdditionalInfo: additionalInfo,
	}
}
