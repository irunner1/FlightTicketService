package passenger

import (
	"errors"
	"time"
)

// PassengerService interface implements methods for managing passengers.
type PassengerService interface {
	GetPassengers() ([]Passenger, error)
	GetPassengerByID(passengerID string) (Passenger, error)
	CreatePassenger(passenger *Passenger) error
	UpdatePassenger(passengerID string, passenger *Passenger) error
	DeletePassenger(passengerID string) error
}

// passengerServiceImpl structure implements interface PassengerService.
type passengerServiceImpl struct {
	passengers []Passenger
}

// NewPassengerService creates PassengerService with predefined passengers.
func NewPassengerService() PassengerService {
	samplePassengers := []Passenger{
		{
			ID:        "passenger1",
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "test123123",
			CreatedAt: time.Now().Add(-24 * time.Hour),
		},
		{
			ID:        "passenger2",
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Password:  "password123",
			CreatedAt: time.Now().Add(-48 * time.Hour),
		},
	}
	return &passengerServiceImpl{passengers: samplePassengers}
}

// @Summary Get list of passengers
// @Description get passengers
// @Tags passengers
// @Accept  json
// @Produce  json
// @Success 200 {array} Passenger
// @Router /api/v1/passengers [get]
func (s *passengerServiceImpl) GetPassengers() ([]Passenger, error) {
	return s.passengers, nil
}

// @Summary Get passenger by ID
// @Description Gets passenger details for a specific passenger ID.
// @Tags passengers
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the passenger"
// @Success 200 {object} Passenger
// @Failure 404 "Missing required parameters"
// @Router /api/v1/passengers/{id} [get]
func (s *passengerServiceImpl) GetPassengerByID(passengerID string) (Passenger, error) {
	for _, passenger := range s.passengers {
		if passenger.ID == passengerID {
			return passenger, nil
		}
	}
	return Passenger{}, errors.New("passenger not found")
}

// @Summary Creates passenger
// @Description Creates new passenger profile
// @Tags passengers
// @Accept json
// @Produce json
// @Param passenger body Passenger true "Passenger data"
// @Success 200 "Passenger created"
// @Failure 400 "Invalid passenger data"
// @Router /api/v1/passengers/create [post]
func (s *passengerServiceImpl) CreatePassenger(passenger *Passenger) error {
	if passenger == nil {
		return errors.New("passenger cannot be nil")
	}
	if passenger.ID == "" {
		return errors.New("passenger ID cannot be empty")
	}
	for _, p := range s.passengers {
		if p.ID == passenger.ID {
			return errors.New("passenger with this ID already exists")
		}
	}
	passenger.CreatedAt = time.Now()
	s.passengers = append(s.passengers, *passenger)
	return nil
}

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
func (s *passengerServiceImpl) UpdatePassenger(passengerID string, updatedPassenger *Passenger) error {
	if passengerID == "" || updatedPassenger == nil {
		return errors.New("passenger ID cannot be empty and updated passenger cannot be nil")
	}
	for i, p := range s.passengers {
		if p.ID == passengerID {
			s.passengers[i].FirstName = updatedPassenger.FirstName
			s.passengers[i].LastName = updatedPassenger.LastName
			s.passengers[i].Email = updatedPassenger.Email
			s.passengers[i].Password = updatedPassenger.Password
			return nil
		}
	}
	return errors.New("passenger not found")
}

// @Summary Delete passenger
// @Description Delete a passenger by their unique identifier
// @Tags passengers
// @Accept json
// @Produce json
// @Param id path string true "Unique identifier of the passenger"
// @Success 200 "Passenger deleted"
// @Failure 404 "Passenger not found"
// @Router /api/v1/passengers/{id}/delete [delete]
func (s *passengerServiceImpl) DeletePassenger(passengerID string) error {
	if passengerID == "" {
		return errors.New("passenger ID cannot be empty")
	}
	for i, p := range s.passengers {
		if p.ID == passengerID {
			s.passengers = append(s.passengers[:i], s.passengers[i+1:]...)
			return nil
		}
	}
	return errors.New("passenger not found")
}
