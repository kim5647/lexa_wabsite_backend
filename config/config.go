package config

import (
	"context"
	db "lexa_wabsite_backend/db/sqlc"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Config() {
	connStr := "postgres://postgres:3006@localhost:3006/lexa_group?sslmode=disable"

	pool, err := pgxpool.New(context.Background(), connStr)
	qer := db.New(pool)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	print("Result: %v", qer)
	defer pool.Close()
}
