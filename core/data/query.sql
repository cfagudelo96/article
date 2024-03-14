-- name: CreateRestaurant :exec
insert into restaurant (
    id, name, schedule, created_at, updated_at
) values ($1, $2, $3, $4, $5);
