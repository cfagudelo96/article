package restaurant

import (
	"time"

	"github.com/google/uuid"
)

type Restaurant struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateRequest struct {
	Name string
}

func (r CreateRequest) ToRestaurant() Restaurant {
	now := time.Now()
	return Restaurant{
		ID:        uuid.New(),
		Name:      r.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
