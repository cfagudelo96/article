// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRestaurant = `-- name: CreateRestaurant :exec
insert into restaurant (
    id, name, created_at, updated_at
) values ($1, $2, $3, $4)
`

type CreateRestaurantParams struct {
	ID        pgtype.UUID
	Name      string
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

func (q *Queries) CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) error {
	_, err := q.db.Exec(ctx, createRestaurant,
		arg.ID,
		arg.Name,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}
