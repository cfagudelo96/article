package restaurant

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cfagudelo96/article/core/restaurant"
	"github.com/cfagudelo96/article/handlers"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
)

var _ restaurantv1.RestaurantServiceServer = (*Handler)(nil)

type Handler struct {
	*handlers.ValidatorHandler
	restaurant *restaurant.Core
}

func NewHandler(c *restaurant.Core) (*Handler, error) {
	vh, err := handlers.NewValidatorHandler(
		protovalidate.WithMessages(
			&restaurantv1.CreateRestaurantRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing validator handler: %w", err)
	}
	return &Handler{
		ValidatorHandler: vh,
		restaurant:       c,
	}, nil
}

func (h *Handler) CreateRestaurant(
	ctx context.Context, pbreq *restaurantv1.CreateRestaurantRequest,
) (*restaurantv1.CreateRestaurantResponse, error) {
	if err := h.Validate(pbreq); err != nil {
		return nil, fmt.Errorf("validating request: %w", err)
	}
	req := toCreateRestaurantRequest(pbreq)
	r, err := h.restaurant.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("creating restaurant: %w", err)
	}
	return &restaurantv1.CreateRestaurantResponse{
		Restaurant: toProtoRestaurant(r),
	}, nil
}

func (h *Handler) CreateRestaurantPointer(
	ctx context.Context, pbreq *restaurantv1.CreateRestaurantRequest,
) (*restaurantv1.CreateRestaurantResponse, error) {
	if err := h.Validate(pbreq); err != nil {
		return nil, fmt.Errorf("validating request: %w", err)
	}
	r, err := h.restaurant.CreatePointer(ctx, pbreq)
	if err != nil {
		return nil, fmt.Errorf("creating restaurant: %w", err)
	}
	return &restaurantv1.CreateRestaurantResponse{
		Restaurant: r,
	}, nil
}
