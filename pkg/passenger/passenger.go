package passenger

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// Storage collects methods for postgres
type Storage interface {
	CreatePassenger(*Passenger) error
	GetPassengers() ([]*Passenger, error)
	GetPassengerByID(passengerID string) (*Passenger, error)
	UpdatePassenger(passengerID string, passenger *Passenger) error
	DeletePassenger(passengerID string) error
}

// PostgresStore stores db pointer
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore initializes a new PostgresStore with a shared database connection.
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

// Init initializes db with data
func (ps *PostgresStore) Init() error {
	return ps.CreatePassengerTable()
}

// CreatePassengerTable creates passenger table in db
func (ps *PostgresStore) CreatePassengerTable() error {
	query := `create table if not exists passengers (
		ID serial primary key,
		first_name varchar(30),
		last_name  varchar(30),    
		email     varchar(30),    
		password  varchar(30),    
		created_at timestamp
	)`

	_, err := ps.db.Exec(query)
	return err
}

// CreatePassenger creates passenger in table
// @Summary Creates passenger
// @Description Creates new passenger profile
// @Tags passengers
// @Accept json
// @Produce json
// @Param passenger body Passenger true "Passenger data"
// @Success 200 "Passenger created"
// @Failure 400 "Invalid passenger data"
// @Router /api/v1/passengers/create [post]
func (ps *PostgresStore) CreatePassenger(pass *Passenger) error {
	query := `insert into passengers
	(first_name, last_name, email, password, created_at) 
	values ($1, $2, $3, $4, $5)`

	resp, err := ps.db.Query(
		query,
		pass.FirstName,
		pass.LastName,
		pass.Email,
		pass.Password,
		pass.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

// UpdatePassenger updates pasanger by id
// @Summary Update passenger data
// @Description Update an existing passenger's details.
// @Tags passengers
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the passenger"
// @Param name query string true "User name"
// @Param surname query string true "User surname"
// @Param email query string true "User email"
// @Param password query string true "User password"
// @Success 200 "Passenger updated"
// @Failure 404 "Passenger not found"
// @Router /api/v1/passengers/{id}/update [post]
func (ps *PostgresStore) UpdatePassenger(id string, newPassenger *Passenger) error {
	if newPassenger == nil {
		return errors.New("update request is nil")
	}

	query := "UPDATE passengers SET first_name = $1, last_name = $2, email = $3 WHERE id = $4"

	_, err := ps.db.Query(
		query,
		newPassenger.FirstName,
		newPassenger.LastName,
		newPassenger.Email,
		id)

	if err != nil {
		return err
	}

	return nil
}

// DeletePassenger deletes pasanger from db
// @Summary Delete passenger
// @Description Delete a passenger by their unique identifier
// @Tags passengers
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the passenger"
// @Success 200 "Passenger deleted"
// @Failure 404 "Passenger not found"
// @Router /api/v1/passengers/{id}/delete [delete]
func (ps *PostgresStore) DeletePassenger(id string) error {
	_, err := ps.db.Query("delete from account where id = $1", id)
	return err
}

// GetPassengerByID returns passenger by id
// @Summary Get passenger by ID
// @Description Gets passenger details for a specific passenger ID.
// @Tags passengers
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the passenger"
// @Success 200 {object} Passenger
// @Failure 404 "Missing required parameters"
// @Router /api/v1/passengers/{id} [get]
func (ps *PostgresStore) GetPassengerByID(id string) (*Passenger, error) {
	rows, err := ps.db.Query("select * from passengers where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanPassenger(rows)
	}

	return nil, errors.New("passenger not found")
}

// GetPassengers return list of passengers
// @Summary Get list of passengers
// @Description get passengers
// @Tags passengers
// @Accept  json
// @Produce  json
// @Success 200 {array} Passenger
// @Router /api/v1/passengers [get]
func (ps *PostgresStore) GetPassengers() ([]*Passenger, error) {
	rows, err := ps.db.Query("select * from passengers")
	if err != nil {
		return nil, err
	}

	passengers := []*Passenger{}
	for rows.Next() {
		passenger, err := scanPassenger(rows)

		if err != nil {
			return nil, err
		}

		passengers = append(passengers, passenger)
	}

	return passengers, nil
}

func scanPassenger(rows *sql.Rows) (*Passenger, error) {
	passenger := new(Passenger)
	err := rows.Scan(
		&passenger.ID,
		&passenger.FirstName,
		&passenger.LastName,
		&passenger.Email,
		&passenger.Password,
		&passenger.CreatedAt)

	return passenger, err
}
