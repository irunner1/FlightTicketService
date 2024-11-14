package database

import (
	"database/sql"
	"flightticketservice/utils"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB initializes and returns a database connection.
func ConnectDB(host, user, password, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", host, user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		utils.ErrorLog.Fatal("Failed to open a DB connection: ", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		utils.ErrorLog.Fatal("Failed to ping DB: ", err)
		return nil, err
	}

	return db, nil
}
