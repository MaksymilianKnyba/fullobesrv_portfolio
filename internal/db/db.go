package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Conn *pgx.Conn
}

func New(dsn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &DB{Conn: conn}, nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.Conn.Ping(ctx)
}