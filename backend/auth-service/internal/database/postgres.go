package database

import (
	"context"
	"fmt"

	"github.com/Lemper29/Jarcy/auth-service/internal/database/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDatabase struct {
	Pool    *pgxpool.Pool
	Queries *queries.Queries
}

func NewPostgresDatabase(connStr string) (*PostgresDatabase, error) {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("connection: %w", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	queries := queries.New(dbpool)

	fmt.Println("Start db")
	return &PostgresDatabase{
		Pool:    dbpool,
		Queries: queries,
	}, nil
}
