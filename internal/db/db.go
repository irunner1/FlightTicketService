package db

import (
	"database/sql"
	"flightticketservice/pkg/passenger"
	"flightticketservice/utils"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreatePassenger(*passenger.Passenger) error
}

// PostgresStore stores db pointer
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates connection woth postgres
func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=dbroot sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		utils.ErrorLog.Fatal(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		utils.ErrorLog.Fatal(err)
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

// Init initializes db with data
func (s *PostgresStore) Init() error {
	return s.CreatePassengerTable()
}

// CreatePassengerTable creates passenger table in db
func (s *PostgresStore) CreatePassengerTable() error {
	query := `create table if not exists passengers (
		ID serial primary key,
		first_name varchar(30),
		last_name  varchar(30),    
		email     varchar(30),    
		password  varchar(30),    
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

// CreatePassenger creates passenger in table
func (s *PostgresStore) CreatePassenger(pass *passenger.Passenger) error {
	query := `insert into passengers
	(first_name, last_name, email, password, created_at) 
	values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Query(
		query,
		pass.FirstName,
		pass.LastName,
		pass.Email,
		pass.Password,
		pass.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)
	return nil
}

func (s *PostgresStore) UpdatePassenger(*passenger.Passenger) error { return nil }

func (s *PostgresStore) DeletePassenger(id string) error { return nil }

func (s *PostgresStore) GetPassengerByID(id string) (*passenger.Passenger, error) { return nil, nil }
