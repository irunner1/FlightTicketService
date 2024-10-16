package passenger

import "time"

// Passenger stores information about a user.
type Passenger struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// CreatePassengerRequest collects info about passenger for request.
type CreatePassengerRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

// CreateNewPassenger creates new passenger by passed params
func CreateNewPassenger(firstName, lastName, email, password string) *Passenger {
	return &Passenger{
		ID:        "FL001",
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}
