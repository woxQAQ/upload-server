package controller

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/ratelimit"

	"github.com/woxQAQ/upload-server/internal/types"
)

type Sinker interface {
	Import(ctx context.Context, db, table string) error
}

type mysqlSinker struct{}

// Import implements Sinker.
func (m mysqlSinker) Import(ctx context.Context, db string, table string) error {
	panic("unimplemented")
}

type pgSinker struct {
	connPool   *pgxpool.Pool
	rpsLimiter ratelimit.Limiter
	buffer     *DtQueue[[]string]
}

// Import implements Sinker.
func (p *pgSinker) Import(ctx context.Context, db string, table string) error {
	qb := queryBulder{
		strings.Builder{},
		db, table, "", false,
	}
	prev := p.rpsLimiter.Take()
	prevIndex := 0
	index := 0
	for {
		select {
		case <-ctx.Done():
			slog.Error(ctx.Err().Error())
			return ctx.Err()
		default:
			now := p.rpsLimiter.Take()
			row, err := p.buffer.Pop()
			if err != nil {
				return err
			}
			qb.PushRow(row)
			index++
			if now.Sub(prev) >= time.Second {
				if index-prevIndex > 0 {
					slog.Info("begin to import",
						slog.Int("start index", prevIndex),
						slog.Int("end index", index),
					)
					prevIndex = index
					query := qb.String()
					_, err = p.connPool.Exec(ctx, query)
					if err != nil {
						slog.Error("error when import",
							slog.Int("start index", prevIndex),
							slog.Int("end index", index),
							slog.String("error", err.Error()),
						)
					}
					qb.Clear()
				}
			}
		}
	}
}

func newSinker(ctx context.Context, engineType string, dsn string, buffer *DtQueue[[]string]) Sinker {
	rpsLimiter := ratelimit.New(1024)
	switch engineType {
	case types.Mysql:
		return mysqlSinker{}
	case types.Postgres:
		pool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			panic(err)
		}
		return &pgSinker{
			connPool:   pool,
			rpsLimiter: rpsLimiter,
			buffer:     buffer,
		}
	}
	return nil
}
