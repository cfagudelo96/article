package handlers

import (
	"errors"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type ValidatorHandler struct {
	validator *protovalidate.Validator
}

func NewValidatorHandler(options ...protovalidate.ValidatorOption) (*ValidatorHandler, error) {
	v, err := protovalidate.New(options...)
	if err != nil {
		return nil, fmt.Errorf("initializing protovalidate: %w", err)
	}
	return &ValidatorHandler{
		validator: v,
	}, nil
}

func (h *ValidatorHandler) Validate(msg proto.Message) error {
	if err := h.validator.Validate(msg); err != nil {
		var vErr *protovalidate.ValidationError
		if errors.As(err, &vErr) {
			return validationErrorToStatusError(vErr)
		}
		return fmt.Errorf("unexpected proto validation error: %w", err)
	}
	return nil
}

func validationErrorToStatusError(vErr *protovalidate.ValidationError) error {
	if vErr == nil {
		return nil
	}
	vs := make([]*errdetails.BadRequest_FieldViolation, len(vErr.Violations))
	for i, v := range vErr.Violations {
		vs[i] = &errdetails.BadRequest_FieldViolation{
			Field:       v.GetFieldPath(),
			Description: v.GetMessage(),
		}
	}
	s := status.New(codes.InvalidArgument, "Invalid values provided")
	d := &errdetails.BadRequest{FieldViolations: vs}
	s, err := s.WithDetails(d)
	if err != nil {
		return fmt.Errorf("appending details to status: %w", err)
	}
	return s.Err()
}
