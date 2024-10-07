package booking

import (
	"errors"
	"time"
)

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

// BookingService interface inmplements methods for booking.
type BookingService interface {
	BookTicket(ticket *Ticket) error
	CancelTicket(ticketID string) error
	ChangeFlight(ticketID string, newFlightID string) error
}

// bookingServiceImpl structure implements interface BookingService.
type bookingServiceImpl struct {
}

// NewBookingService creates BookingService.
func NewBookingService() BookingService {
	return &bookingServiceImpl{}
}

// BookTicket реализует бизнес-логику бронирования билета.
func (s *bookingServiceImpl) BookTicket(ticket *Ticket) error {
	return nil
}

// CancelTicket cancel ticket for flight.
func (s *bookingServiceImpl) CancelTicket(ticketID string) error {
	if ticketID == "" {
		return errors.New("ticket ID cannot be empty")
	}
	return nil
}

// ChangeFlight change flights.
func (s *bookingServiceImpl) ChangeFlight(ticketID string, newFlightID string) error {
	if ticketID == "" || newFlightID == "" {
		return errors.New("ticket ID and new flight ID cannot be empty")
	}
	return nil
}
