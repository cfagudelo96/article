package restaurant

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cfagudelo96/article/core/restaurant"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ restaurantv1.RestaurantServiceServer = (*Handler)(nil)

type Handler struct {
	validator  *protovalidate.Validator
	restaurant *restaurant.Core
}

func NewHandler(c *restaurant.Core) (*Handler, error) {
	v, err := protovalidate.New(
		protovalidate.WithMessages(
			&restaurantv1.CreateRestaurantRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing protovalidate: %w", err)
	}
	return &Handler{
		validator:  v,
		restaurant: c,
	}, nil
}

func (h *Handler) CreateRestaurant(
	ctx context.Context, pbreq *restaurantv1.CreateRestaurantRequest,
) (*restaurantv1.CreateRestaurantResponse, error) {
	if err := h.validator.Validate(pbreq); err != nil {
		return nil, fmt.Errorf("validating request: %w", err)
	}
	req := toCreateRestaurantRequest(pbreq)
	r, err := h.restaurant.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("creating restaurant: %w", err)
	}
	return &restaurantv1.CreateRestaurantResponse{
		Restaurant: toPBRestaurant(r),
	}, nil
}

func toPBRestaurant(r restaurant.Restaurant) *restaurantv1.Restaurant {
	return &restaurantv1.Restaurant{
		Id:        r.ID.String(),
		Name:      r.Name,
		CreatedAt: timestamppb.New(r.CreatedAt),
		UpdatedAt: timestamppb.New(r.UpdatedAt),
	}
}

func toCreateRestaurantRequest(req *restaurantv1.CreateRestaurantRequest) restaurant.CreateRequest {
	return restaurant.CreateRequest{
		Name: req.GetName(),
	}
}
