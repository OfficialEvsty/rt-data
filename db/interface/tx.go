package db

import (
	"context"
	"database/sql"
)

// ITxExecutor интерфейс бд, выполняет транзакционные запросы
type ITxExecutor interface {
	WithTx(ctx context.Context, fn func(ctx context.Context, tx *sql.Tx) error) error
}
