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
	tickets []Ticket
}

// NewBookingService creates BookingService.
func NewBookingService() BookingService {
	sampleTickets := []Ticket{
		{
			ID:             "ticket1",
			FlightID:       "flight1",
			PassengerID:    "passenger1",
			BookingTime:    time.Now().Add(-48 * time.Hour),
			DepartureTime:  time.Now().Add(24 * time.Hour),
			ArrivalTime:    time.Now().Add(30 * time.Hour),
			Status:         "booked",
			SeatNumber:     "12A",
			AdditionalInfo: "Vegetarian meal",
		},
		{
			ID:             "ticket2",
			FlightID:       "flight2",
			PassengerID:    "passenger2",
			BookingTime:    time.Now().Add(-24 * time.Hour),
			DepartureTime:  time.Now().Add(48 * time.Hour),
			ArrivalTime:    time.Now().Add(54 * time.Hour),
			Status:         "confirmed",
			SeatNumber:     "1C",
			AdditionalInfo: "Window seat",
		},
	}
	return &bookingServiceImpl{tickets: sampleTickets}
}

// BookTicket books the ticket.
func (s *bookingServiceImpl) BookTicket(ticket *Ticket) error {
	if ticket == nil {
		return errors.New("ticket cannot be nil")
	}

	if ticket.ID == "" {
		return errors.New("ticket ID cannot be empty")
	}

	for _, t := range s.tickets {
		if t.ID == ticket.ID {
			return errors.New("ticket with this ID already exists")
		}
	}

	ticket.Status = "booked"
	ticket.BookingTime = time.Now()
	s.tickets = append(s.tickets, *ticket)
	return nil
}

// CancelTicket cancel ticket for flight.
func (s *bookingServiceImpl) CancelTicket(ticketID string) error {
	if ticketID == "" {
		return errors.New("ticket ID cannot be empty")
	}
	for i, t := range s.tickets {
		if t.ID == ticketID {
			if t.Status == "cancelled" {
				return errors.New("ticket is already cancelled")
			}
			s.tickets[i].Status = "cancelled"
			return nil
		}
	}
	return errors.New("ticket not found")
}

// ChangeFlight change flights.
func (s *bookingServiceImpl) ChangeFlight(ticketID string, newFlightID string) error {
	if ticketID == "" || newFlightID == "" {
		return errors.New("ticket ID and new flight ID cannot be empty")
	}
	for i, t := range s.tickets {
		if t.ID == ticketID {
			if t.Status == "cancelled" {
				return errors.New("cannot change flight for a cancelled ticket")
			}
			s.tickets[i].FlightID = newFlightID
			s.tickets[i].Status = "confirmed"
			return nil
		}
	}
	return errors.New("ticket not found")
}
