package restaurant

import (
	"context"
	"errors"
	"fmt"
	"strings"

	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
)

var ErrNotFound = errors.New("restaurant not found")

//go:generate mockgen -source=restaurant.go -destination=mocks/mocks.go Storer

type Storer interface {
	CreateRestaurant(ctx context.Context, r Restaurant) error
	CreateRestaurantPointer(ctx context.Context, r *restaurantv1.Restaurant) error
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
	r = mod1(r)
	r = mod2(r)
	if err := c.storer.CreateRestaurant(ctx, r); err != nil {
		return Restaurant{}, fmt.Errorf("creating in store: %w", err)
	}
	return r, nil
}

func mod1(r Restaurant) Restaurant {
	r.Name = strings.ToUpper(r.Name)
	return r
}

func mod2(r Restaurant) Restaurant {
	r.Name += "!"
	return r
}

func (c *Core) CreatePointer(
	ctx context.Context, req *restaurantv1.CreateRestaurantRequest,
) (*restaurantv1.Restaurant, error) {
	r := createRestaurantRequestToRestaurant(req)
	mod1p(r)
	mod2p(r)
	if err := c.storer.CreateRestaurantPointer(ctx, r); err != nil {
		return nil, fmt.Errorf("creating in store: %w", err)
	}
	return r, nil
}

func mod1p(r *restaurantv1.Restaurant) {
	r.Name = strings.ToUpper(r.GetName())
}

func mod2p(r *restaurantv1.Restaurant) {
	r.Name = r.GetName() + "!"
}
