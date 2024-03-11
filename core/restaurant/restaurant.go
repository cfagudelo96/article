package restaurant

import (
	"context"
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("restaurant not found")

type Storer interface {
	CreateRestaurant(ctx context.Context, r Restaurant) error
}

type Core struct {
	storer Storer
}

func NewCore(s Storer) *Core {
	return &Core{
		storer: s,
	}
}

func (c *Core) Create(ctx context.Context, req CreateRequest) (Restaurant, error) {
	r := req.ToRestaurant()
	if err := c.storer.CreateRestaurant(ctx, r); err != nil {
		return Restaurant{}, fmt.Errorf("creating in store: %w", err)
	}
	return r, nil
}
