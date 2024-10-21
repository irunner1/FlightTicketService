package booking

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// BookingService interface inmplements methods for booking.
type BookingService interface {
	GetTickets() ([]*Ticket, error)
	GetTicketByID(ticketID string) (*Ticket, error)
	BookTicket(ticket *Ticket) error
	CancelTicket(ticketID string) error
	ChangeFlight(ticketID string, newFlightID string) error
	CreateTicket(newTicket *Ticket) error
	UpdateTicket(id string, newTicket *Ticket) error
	DeleteTicket(ticketID string) error
}

// BookingStore structure implements interface FlightService.
type BookingStore struct {
	db *sql.DB
}

// NewBookingStore initializes a new PostgresStore with a shared database connection.
func NewBookingStore(db *sql.DB) *BookingStore {
	return &BookingStore{db: db}
}

// Init initializes db with data
func (bs *BookingStore) Init() error {
	return bs.CreateFlightsTable()
}

// CreateFlightsTable creates flights table in db
func (bs *BookingStore) CreateFlightsTable() error {
	query := `create table if not exists booking_flights (
		ID serial primary key,
		flight_id varchar(10),
		passenger_id varchar(10),
		booking_time timestamp,
		departure_time timestamp,
		arrival_time timestamp,
		status varchar(30),
		seat_number varchar(30),
		additional_info varchar(100)
	)`

	_, err := bs.db.Exec(query)
	return err
}

// CreateTicket creates ticket in table
// @Summary Creates ticket
// @Description Creates new ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param flight body CreateTicketReq true "Ticket data"
// @Success 200 "Ticket created"
// @Failure 400 "Invalid ticket data"
// @Router /api/v1/tickets/create [post]
func (bs *BookingStore) CreateTicket(ticket *Ticket) error {
	query := `insert into booking_flights (
	flight_id,
	passenger_id,
	booking_time,
	departure_time,
	arrival_time,
	status,
	seat_number,
	additional_info
	)
	values ($1, $2, $3, $4, $5, $6, $7, $8)`

	resp, err := bs.db.Query(
		query,
		ticket.FlightID,
		ticket.PassengerID,
		ticket.BookingTime,
		ticket.DepartureTime,
		ticket.ArrivalTime,
		ticket.Status,
		ticket.SeatNumber,
		ticket.AdditionalInfo)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

// BookTicket books a new ticket
// @Summary Book a new ticket
// @Description Creates a new ticket booking for a flight.
// @Tags booking
// @Accept json
// @Produce json
// @Param ticket body CreateTicketReq true "Ticket data"
// @Success 200 "Ticket successfully booked"
// @Failure 400 "Invalid ticket data"
// @Router /api/v1/tickets/book [post]
func (bs *BookingStore) BookTicket(ticket *Ticket) error {
	if ticket == nil {
		return errors.New("ticket cannot be nil")
	}

	if ticket.ID == "" {
		return errors.New("ticket ID cannot be empty")
	}

	rows, err := bs.db.Query("select * from booking_flights where id = $1", ticket.ID)

	if rows != nil && err == nil {
		return errors.New("ticket with this ID already exists")
	}

	ticket.Status = "booked"
	ticket.BookingTime = time.Now()

	query := `insert into booking_flights
	(flight_id, passenger_id, booking_time, departure_time, arrival_time, status, seat_number, additional_info)
	values ($1, $2, $3, $4, $5, $6, $7, $8)`

	resp, err := bs.db.Query(
		query,
		ticket.FlightID,
		ticket.PassengerID,
		ticket.BookingTime,
		ticket.DepartureTime,
		ticket.ArrivalTime,
		ticket.Status,
		ticket.SeatNumber,
		ticket.AdditionalInfo)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

// CancelTicket cancels an existing ticket
// @Summary Cancel an existing ticket
// @Description Cancels an existing ticket using the ticket ID.
// @Tags booking
// @Accept json
// @Produce json
// @Param ticketID path string true "The ID of the ticket to cancel"
// @Success 200 "Ticket successfully cancelled"
// @Failure 404 "Ticket not found"
// @Router /api/v1/tickets/{ticketID}/cancel [post]
func (bs *BookingStore) CancelTicket(ticketID string) error {
	if ticketID == "" {
		return errors.New("ticket ID cannot be empty")
	}

	query := `update booking_flights set status = 'cancelled' where id = $1 and status != 'cancelled'`
	res, err := bs.db.Exec(query, ticketID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found or already cancelled")
	}

	return nil
}

// ChangeFlight changes the flight associated with a ticket
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
func (bs *BookingStore) ChangeFlight(ticketID string, newFlightID string) error {
	if ticketID == "" || newFlightID == "" {
		return errors.New("ticket ID and new flight ID cannot be empty")
	}

	query := `update booking_flights set flight_id = $1, status = 'confirmed' where id = $2 and status != 'cancelled'`
	res, err := bs.db.Exec(query, newFlightID, ticketID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found or cannot change flight for a cancelled ticket")
	}

	return nil
}

// GetTicketByID returns ticket details for a specific ticket ID
// @Summary Get ticket by ID
// @Description Returns ticket details for a specific ticket ID.
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the ticket"
// @Success 200 {object} Ticket
// @Failure 404 "ticket not found"
// @Router /api/v1/tickets/{id} [get]
func (bs *BookingStore) GetTicketByID(ticketID string) (*Ticket, error) {
	query := `select * from booking_flights where id = $1`
	row := bs.db.QueryRow(query, ticketID)

	ticket := &Ticket{}
	err := row.Scan(
		&ticket.ID,
		&ticket.FlightID,
		&ticket.PassengerID,
		&ticket.BookingTime,
		&ticket.DepartureTime,
		&ticket.ArrivalTime,
		&ticket.Status,
		&ticket.SeatNumber,
		&ticket.AdditionalInfo,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ticket not found")
		}
		return nil, err
	}

	return ticket, nil
}

// GetTickets returns a list of all tickets
// @Summary Get list of tickets
// @Description get tickets
// @Tags tickets
// @Accept  json
// @Produce  json
// @Success 200 {array} Ticket
// @Router /api/v1/tickets [get]
func (bs *BookingStore) GetTickets() ([]*Ticket, error) {
	rows, err := bs.db.Query("select * from booking_flights")
	if err != nil {
		return nil, err
	}

	tickets := []*Ticket{}
	for rows.Next() {
		ticket := &Ticket{}
		err := rows.Scan(
			&ticket.ID,
			&ticket.FlightID,
			&ticket.PassengerID,
			&ticket.BookingTime,
			&ticket.DepartureTime,
			&ticket.ArrivalTime,
			&ticket.Status,
			&ticket.SeatNumber,
			&ticket.AdditionalInfo,
		)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

// UpdateTicket updates the details of an existing ticket
// @Summary Update ticket details
// @Description Updates the details of a ticket using the ticket ID.
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the ticket"
// @Param ticket body Ticket true "Updated ticket data"
// @Success 200 "Ticket successfully updated"
// @Failure 400 "Invalid ticket data"
// @Failure 404 "Ticket not found"
// @Router /api/v1/tickets/{id}/update [post]
func (bs *BookingStore) UpdateTicket(id string, newTicket *Ticket) error {
	if newTicket == nil {
		return errors.New("ticket cannot be nil")
	}

	query := `UPDATE booking_flights SET
	flight_id = $1,
	passenger_id = $2,
	booking_time = $3,
	departure_time = $4,
	arrival_time = $5,
	status = $6,
	seat_number = $7,
	additional_info = $8
	WHERE id = $9`

	res, err := bs.db.Exec(
		query,
		newTicket.FlightID,
		newTicket.PassengerID,
		newTicket.BookingTime,
		newTicket.DepartureTime,
		newTicket.ArrivalTime,
		newTicket.Status,
		newTicket.SeatNumber,
		newTicket.AdditionalInfo,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return nil
}

// DeleteTicket deletes a ticket from the database
// @Summary Delete a ticket
// @Description Deletes a ticket using the ticket ID.
// @Tags tickets
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the ticket"
// @Success 200 "Ticket successfully deleted"
// @Failure 404 "Ticket not found"
// @Router /api/v1/tickets/{id}/delete [delete]
func (bs *BookingStore) DeleteTicket(ticketID string) error {
	query := `delete from booking_flights where id = $1`

	res, err := bs.db.Exec(query, ticketID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}

	return nil
}
