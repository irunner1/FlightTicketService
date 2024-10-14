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
	GetTickets() ([]Ticket, error)
	BookTicket(ticket *Ticket) error
	CancelTicket(ticketID string) error
	ChangeFlight(ticketID string, newFlightID string) error
	GetTicketByID(ticketID string) (Ticket, error)
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

// @Summary Book a new ticket
// @Description Creates a new ticket booking for a flight.
// @Tags booking
// @Accept json
// @Produce json
// @Param ticket body Ticket true "Ticket data"
// @Success 200 "Ticket successfully booked"
// @Failure 400 "Invalid ticket data"
// @Router /api/v1/tickets/book [post]
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

// @Summary Cancel an existing ticket
// @Description Cancels an existing ticket using the ticket ID.
// @Tags booking
// @Accept json
// @Produce json
// @Param ticketID path string true "The ID of the ticket to cancel"
// @Success 200 "Ticket successfully cancelled"
// @Failure 404 "Ticket not found"
// @Router /api/v1/tickets/{ticketID}/cancel [post]
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

// @Summary Change the flight of a ticket
// @Description Changes the flight associated with a ticket to a new flight using the ticket ID and new flight ID.
// @Tags booking
// @Accept json
// @Produce json
// @Param ticketID path string true "The ID of the ticket to update"
// @Param newFlightID query string true "The new flight ID to associate with the ticket"
// @Success 200 "Flight successfully changed for the ticket"
// @Failure 400 "Invalid parameters"
// @Failure 404 "Ticket not found or cannot change flight for a cancelled ticket"
// @Router /api/v1/{ticketID}/change [post]
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

// @Summary Get ticket by ID
// @Description Returns ticket details for a specific ticket ID.
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the ticket"
// @Success 200 {object} Ticket
// @Failure 404 "ticket not found"
// @Router /api/v1/tickets/{id} [get]
func (s *bookingServiceImpl) GetTicketByID(ticketID string) (Ticket, error) {
	for _, ticket := range s.tickets {
		if ticket.ID == ticketID {
			return ticket, nil
		}
	}
	return Ticket{}, errors.New("flight not found")
}

// @Summary Get list of tickets
// @Description get tickets
// @Tags tickets
// @Accept  json
// @Produce  json
// @Success 200 {array} Ticket
// @Router /api/v1/tickets [get]
func (s *bookingServiceImpl) GetTickets() ([]Ticket, error) {
	return s.tickets, nil
}
