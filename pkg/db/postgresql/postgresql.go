package postgresql

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/kuzuru/codesync/pkg/logger"
	"github.com/kuzuru/codesync/utils"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxConnectionAttempts int) (pool *pgxpool.Pool, err error) {
	DSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	log := logger.New()

	defer log.Sync()

	// Trying to connect to postgres
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, DSN)
		if err != nil {
			return err
		}

		return nil
	}, maxConnectionAttempts, 5*time.Second)

	if err != nil {
		log.Error("Failed to connect to postgres")
	}

	return pool, nil
}
