package database

import "github.com/jackc/pgx/v5/pgxpool"

func NewQueries(pool *pgxpool.Pool) *Queries {
	return &Queries{db: pool}
}
