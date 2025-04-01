package passenger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)


func TestNewPassenger(t *testing.T) {
	firstName := "John"
	lastName := "Doe"
	email := "john.doe@example.com"
	password := "securepassword"

	passenger, err := NewPassenger(firstName, lastName, email, password)

	assert.NoError(t, err, "Expected no error from NewPassenger")
	assert.NotNil(t, passenger, "Expected passenger to be not nil")

	assert.Equal(t, firstName, passenger.FirstName)
	assert.Equal(t, lastName, passenger.LastName)
	assert.Equal(t, email, passenger.Email)

	err = bcrypt.CompareHashAndPassword([]byte(passenger.Password), []byte(password))
	assert.NoError(t, err, "Expected the password to match after hashing")

	assert.WithinDuration(t, time.Now().UTC(), passenger.CreatedAt, time.Second, "Expected CreatedAt to be recent")

	assert.GreaterOrEqual(t, passenger.Number, int64(0), "Expected Number to be non-negative")
	assert.Less(t, passenger.Number, int64(1000000), "Expected Number to be less than 1000000")
}