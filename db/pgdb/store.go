// Querie me se
package pgdb

import "github.com/jackc/pgx/v5/pgxpool"

type Store interface {
	Querier
}

type SQLStrore struct {
	Querier
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) Store {
	return &SQLStrore{
		Querier: New(db),
		db:      db,
	}
}
