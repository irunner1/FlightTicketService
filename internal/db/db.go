package db

import (
	"database/sql"
	"flightticketservice/pkg/passenger"
	"flightticketservice/utils"

	_ "github.com/lib/pq"
)

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

func (s *PostgresStore) Init() error {
	return s.CreatePassengerTable()
}

func (s *PostgresStore) CreatePassengerTable() error {
	query := `create table if not exists passengers (
		ID serial primary key,
		FirstName varchar(20),
		LastName  varchar(20),    
		Email     varchar(20),    
		Password  varchar(20),    
		CreatedAt timestamp,
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreatePassenger(*passenger.Passenger) error { return nil }

func (s *PostgresStore) UpdatePassenger(*passenger.Passenger) error { return nil }

func (s *PostgresStore) DeletePassenger(id string) error { return nil }

func (s *PostgresStore) GetPassengerByID(id string) (*passenger.Passenger, error) { return nil, nil }
