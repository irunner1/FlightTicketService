package flights

import (
	"database/sql"
	"errors"
	"fmt"
)

// FlightService interface for working with flights.
type FlightService interface {
	GetFlights() ([]*Flight, error)
	GetFlightsByParams(params SearchParams) ([]*Flight, error)
	GetFlightByID(flightID string) (*Flight, error)
	CreateFlight(*Flight) error
	UpdateFlight(id string, newFlight *Flight) error
	DeleteFlight(flightID string) error
}

// FlightsStore structure implements interface FlightService.
type FlightsStore struct {
	db *sql.DB
}

// NewFlightsStore initializes a new PostgresStore with a shared database connection.
func NewFlightsStore(db *sql.DB) *FlightsStore {
	return &FlightsStore{db: db}
}

// Init initializes db with data
func (fs *FlightsStore) Init() error {
	return fs.CreateFlightsTable()
}

// CreateFlightsTable creates flights table in db
func (fs *FlightsStore) CreateFlightsTable() error {
	query := `create table if not exists flights (
		ID serial primary key,
		airline varchar(30),
		origin  varchar(30),
		destination varchar(30),
		departure timestamp,
		arrival timestamp,
		price real
	)`

	_, err := fs.db.Exec(query)
	return err
}

// CreateFlight creates flight in table
// @Summary Creates flight
// @Description Creates new flight profile
// @Tags flights
// @Accept json
// @Produce json
// @Param flight body CreateFlightReq true "Flight data"
// @Success 200 "Flight created"
// @Failure 400 "Invalid flight data"
// @Router /api/v1/flights/create [post]
func (fs *FlightsStore) CreateFlight(fl *Flight) error {
	query := `insert into flights
	(airline, origin, destination, departure, arrival, price)
	values ($1, $2, $3, $4, $5, $6)`

	resp, err := fs.db.Query(
		query,
		fl.Airline,
		fl.Origin,
		fl.Destination,
		fl.Departure,
		fl.Arrival,
		fl.Price)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

// UpdateFlight updates pasanger by id
// @Summary Update flight data
// @Description Update an existing flight's details.
// @Tags flights
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the flight"
// @Param flight body Flight true "Flight data"
// @Success 200 "Flight updated"
// @Failure 404 "Flight not found"
// @Router /api/v1/flights/{id}/update [post]
func (fs *FlightsStore) UpdateFlight(id string, newFlight *Flight) error {

	if newFlight == nil {
		return errors.New("update request is nil")
	}

	query := `UPDATE flights SET
	airline = $1, origin = $2, destination = $3, departure = $4, arrival = $5, price = $6
	WHERE id = $7`

	_, err := fs.db.Query(
		query,
		newFlight.Airline,
		newFlight.Origin,
		newFlight.Destination,
		newFlight.Departure,
		newFlight.Arrival,
		newFlight.Price,
		id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteFlight deletes flight from db
// @Summary Delete flight
// @Description Delete a flight by unique identifier
// @Tags flights
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the flight"
// @Success 200 "Passenger deleted"
// @Failure 404 "Passenger not found"
// @Router /api/v1/flights/{id}/delete [delete]
func (fs *FlightsStore) DeleteFlight(id string) error {
	_, err := fs.db.Query("delete from flights where id = $1", id)
	return err
}

// GetFlights returns all flights
// @Summary Get list of flights
// @Description get flights
// @Tags flights
// @Accept  json
// @Produce  json
// @Success 200 {array} Flight
// @Router /api/v1/flights [get]
func (fs *FlightsStore) GetFlights() ([]*Flight, error) {
	rows, err := fs.db.Query("select * from flights")
	if err != nil {
		return nil, err
	}

	passengers := []*Flight{}
	for rows.Next() {
		passenger, err := scanFlight(rows)

		if err != nil {
			return nil, err
		}

		passengers = append(passengers, passenger)
	}

	return passengers, nil
}

// GetFlightsByParams returns list of flights queries by params
// @Summary Search flights by parameters
// @Description Retrieves a list of flights filtered by the provided search parameters.
// @Tags flights
// @Accept json
// @Produce json
// @Param origin query string false "Origin location of the flight"
// @Param destination query string false "Destination location of the flight"
// @Param departure query string false "Departure date and time"
// @Param arrival query string false "Arrival date and time"
// @Success 200 {array} Flight
// @Failure 404 "No flights found matching the search criteria"
// @Router /api/v1/flights/search [get]
func (fs *FlightsStore) GetFlightsByParams(params SearchParams) ([]*Flight, error) {
	rows, err := fs.db.Query("select * from flights")
	if err != nil {
		return nil, err
	}

	flights := []*Flight{}
	for rows.Next() {
		flight, err := scanFlight(rows)

		if err != nil {
			return nil, err
		}

		if params.Origin != "" && flight.Origin != params.Origin {
			continue
		}
		if params.Destination != "" && flight.Destination != params.Destination {
			continue
		}
		if !params.Departure.IsZero() && !flight.Departure.Equal(params.Departure) {
			continue
		}
		if !params.Arrival.IsZero() && !flight.Arrival.Equal(params.Arrival) {
			continue
		}
		flights = append(flights, flight)
	}

	if len(flights) == 0 {
		return nil, errors.New("no flights found matching the search criteria")
	}

	return flights, nil
}

// GetFlightByID returns flight by id
// @Summary Get flight by ID
// @Description Gets flight details for a specific flight ID.
// @Tags flights
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the flight"
// @Success 200 {object} Flight
// @Failure 404 "Flight not found"
// @Router /api/v1/flights/{id} [get]
func (fs *FlightsStore) GetFlightByID(flightID string) (*Flight, error) {
	rows, err := fs.db.Query("select * from flights where id = $1", flightID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanFlight(rows)
	}

	return nil, errors.New("flight not found")
}

func scanFlight(rows *sql.Rows) (*Flight, error) {
	flight := new(Flight)
	err := rows.Scan(
		&flight.ID,
		&flight.Airline,
		&flight.Origin,
		&flight.Destination,
		&flight.Departure,
		&flight.Arrival,
		&flight.Price)

	return flight, err
}
