package store

import (
	"context"
	"fmt"

	"github.com/cfagudelo96/article/core/data/sqlc"
	"github.com/cfagudelo96/article/core/restaurant"
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
	params := sqlc.CreateRestaurantParams{
		ID:        pgtype.UUID{Bytes: r.ID, Valid: true},
		Name:      r.Name,
		CreatedAt: pgtype.Timestamptz{Time: r.CreatedAt, Valid: true},
		UpdatedAt: pgtype.Timestamptz{Time: r.UpdatedAt, Valid: true},
	}
	if err := querier.CreateRestaurant(ctx, params); err != nil {
		return fmt.Errorf("inserting in db: %w", err)
	}
	return nil
}