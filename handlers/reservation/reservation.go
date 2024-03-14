package reservation

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"github.com/cfagudelo96/article/handlers"
	reservationv1 "github.com/cfagudelo96/article/proto/reservation/v1"
)

var _ reservationv1.ReservationServiceServer = (*Handler)(nil)

type Handler struct {
	*handlers.ValidatorHandler
}

func NewHandler() (*Handler, error) {
	vh, err := handlers.NewValidatorHandler(
		protovalidate.WithMessages(
			&reservationv1.CreateReservationRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("initializing validator handler: %w", err)
	}
	return &Handler{
		ValidatorHandler: vh,
	}, nil
}

func (h *Handler) CreateReservation(
	ctx context.Context, req *reservationv1.CreateReservationRequest,
) (*reservationv1.CreateReservationResponse, error) {
	return nil, nil
}
