package flights

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFlight(t *testing.T) {
	airline := "Aeroflot"
	origin := "Москва"
	destination := "Париж"
	departure := time.Date(2024, 3, 16, 10, 0, 0, 0, time.UTC)
	arrival := time.Date(2024, 3, 16, 12, 0, 0, 0, time.UTC)
	price := 1000.0

	flight := NewFlight(airline, origin, destination, departure, arrival, price)

	assert.Equal(t, airline, flight.Airline)
	assert.Equal(t, origin, flight.Origin)
	assert.Equal(t, destination, flight.Destination)
	assert.Equal(t, departure, flight.Departure)
	assert.Equal(t, arrival, flight.Arrival)
	assert.Equal(t, price, flight.Price)
}
