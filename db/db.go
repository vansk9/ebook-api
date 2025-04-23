package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func ConnectDB() error {
	var err error
	dbUrl := os.Getenv("postgresql://postgres:miftah87@localhost:5432/ebook-api")

	Conn, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return err
	}
	return Conn.Ping(context.Background())
}
