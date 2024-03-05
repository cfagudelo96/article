package restaurant

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cfagudelo96/article/core/restaurant/store"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
)

var _ restaurantv1.RestaurantServiceServer = (*Handler)(nil)

type Handler struct {
	validator *protovalidate.Validator
	store     *store.Store
}

func NewHandler(s *store.Store) (*Handler, error) {
	v, err := protovalidate.New(
		protovalidate.WithMessages(
			&restaurantv1.CreateRestaurantRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing handler: %w", err)
	}

	return &Handler{
		validator: v,
		store:     s,
	}, nil
}

func (h *Handler) CreateRestaurant(
	ctx context.Context, req *restaurantv1.CreateRestaurantRequest,
) (*restaurantv1.CreateRestaurantResponse, error) {
	return nil, nil
}
