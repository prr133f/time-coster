package domain

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/rs/zerolog/log"
)

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pgOnce     sync.Once
	pgInstance *Postgres
)

func NewPG(ctx context.Context, DSN string) (*Postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, DSN)
		if err != nil {
			log.Fatal().Err(err)
		}

		pgInstance = &Postgres{DB: db}
	})

	return pgInstance, nil
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}
