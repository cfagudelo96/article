package reservation

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	reservationv1 "github.com/cfagudelo96/article/proto/reservation/v1"
)

var _ reservationv1.ReservationServiceServer = (*Handler)(nil)

type Handler struct {
	validator *protovalidate.Validator
}

func NewHandler() (*Handler, error) {
	v, err := protovalidate.New(
		protovalidate.WithMessages(
			&reservationv1.CreateReservationRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing protovalidate: %w", err)
	}
	return &Handler{
		validator: v,
	}, nil
}

func (h *Handler) CreateReservation(
	ctx context.Context, req *reservationv1.CreateReservationRequest,
) (*reservationv1.CreateReservationResponse, error) {
	return nil, nil
}
