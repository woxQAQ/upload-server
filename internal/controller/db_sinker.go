package controller

import (
	"context"

	"github.com/woxQAQ/upload-server/internal/types"
)

type Sinker interface {
	Import(ctx context.Context, db, table string, rows [][]string) error
}

type mysqlSinker struct{}

// Import implements Sinker.
func (m mysqlSinker) Import(ctx context.Context, db string, table string, rows [][]string) error {
	panic("unimplemented")
}

type pgSinker struct{}

// Import implements Sinker.
func (p pgSinker) Import(ctx context.Context, db string, table string, rows [][]string) error {
	panic("unimplemented")
}

func newSinker(engineType string) Sinker {
	switch engineType {
	case types.Mysql:
		return mysqlSinker{}
	case types.Postgres:
		return pgSinker{}
	}
	return nil
}
