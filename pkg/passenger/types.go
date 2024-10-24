package passenger

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

// LoginRequest stores information for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"-"`
}

// LoginResponse for response after login
type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// Passenger stores information about a user.
type Passenger struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Number    int64     `json:"number"`
}

// CreatePassengerReq collects info about passenger for request.
type CreatePassengerReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

// ValidPassword check if enctypted password is valid
func (p *Passenger) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(pw)) == nil
}

// NewPassenger creates new passenger by passed params
func NewPassenger(firstName, lastName, email, password string) (*Passenger, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Passenger{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(encpw),
		CreatedAt: time.Now().UTC(),
		Number:    int64(rand.Intn(1000000)),
	}, nil
}
