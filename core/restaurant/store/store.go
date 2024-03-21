package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cfagudelo96/article/core/data/sqlc"
	"github.com/cfagudelo96/article/core/restaurant"
	restaurantv1 "github.com/cfagudelo96/article/proto/restaurant/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Store struct {
	connection *pgx.Conn
}

func NewStore(connection *pgx.Conn) *Store {
	return &Store{
		connection: connection,
	}
}

func (s *Store) CreateRestaurant(ctx context.Context, r restaurant.Restaurant) error {
	querier := sqlc.New(s.connection)
	scheduleAsJSON, err := json.Marshal(r.WeeklySchedule)
	if err != nil {
		return fmt.Errorf("marshalling schedule as JSON: %w", err)
	}
	params := sqlc.CreateRestaurantParams{
		ID:        pgtype.UUID{Bytes: r.ID, Valid: true},
		Name:      r.Name,
		Schedule:  scheduleAsJSON,
		CreatedAt: pgtype.Timestamptz{Time: r.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: r.UpdatedAt, Valid: true},
	}
	if err = querier.CreateRestaurant(ctx, params); err != nil {
		return fmt.Errorf("inserting in db: %w", err)
	}
	return nil
}

func (s *Store) CreateRestaurantPointer(ctx context.Context, r *restaurantv1.Restaurant) error {
	querier := sqlc.New(s.connection)
	scheduleAsJSON, err := json.Marshal(r.GetWeeklySchedule())
	if err != nil {
		return fmt.Errorf("marshalling schedule as JSON: %w", err)
	}
	bytes := []byte(r.GetId())
	params := sqlc.CreateRestaurantParams{
		ID:        pgtype.UUID{Bytes: ([16]byte)(bytes), Valid: true},
		Name:      r.GetName(),
		Schedule:  scheduleAsJSON,
		CreatedAt: pgtype.Timestamptz{Time: r.GetCreatedAt().AsTime(), Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: r.GetUpdatedAt().AsTime(), Valid: true},
	}
	if err = querier.CreateRestaurant(ctx, params); err != nil {
		return fmt.Errorf("inserting in db: %w", err)
	}
	return nil
}
