package booking

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTicket(t *testing.T) {
	departureTime := time.Date(2023, 4, 15, 8, 0, 0, 0, time.UTC)
	arrivalTime := time.Date(2023, 4, 15, 20, 0, 0, 0, time.UTC)
	flightID := "flight1"
	passengerID := "passenger1"
	seatNumber := "12A"
	status := "V"
	additionalInfo := "booked"

	ticket := CreateNewTicket(
		flightID,
		passengerID,
		status,
		seatNumber,
		additionalInfo,
		departureTime,
		arrivalTime,
	)

	assert.Equal(t, flightID, ticket.FlightID)
	assert.Equal(t, passengerID, ticket.PassengerID)
	assert.Equal(t, seatNumber, ticket.SeatNumber)
	assert.Equal(t, status, ticket.Status)
	assert.Equal(t, additionalInfo, ticket.AdditionalInfo)
	assert.Equal(t, departureTime, ticket.DepartureTime)
	assert.Equal(t, arrivalTime, ticket.ArrivalTime)
}
