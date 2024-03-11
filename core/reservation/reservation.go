package reservation

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrNoMoreAllowed = errors.New("no more reservations allowed")

type NoSpotsAvailable struct {
	RestaurantID uuid.UUID
	// TODO: Other available spots
}

type Reservation struct {
	ID           uuid.UUID
	RestaurantID uuid.UUID
	Client       Client
	Guests       int
	Time         time.Time
	CreatedAt    time.Time
}

type Client struct {
	Name  string
	Email string
}

type CreateReservationRequest struct {
	RestaurantID uuid.UUID
	Client       Client
	Guests       int
	Time         time.Time
}

type Core struct {
}
