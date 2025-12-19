package repos

import (
	"context"
	"database/sql"
	db "github.com/OfficialEvsty/rt-data/db/interface"
	"github.com/OfficialEvsty/rt-data/models"
	repos "github.com/OfficialEvsty/rt-data/repos/interface"
	"github.com/google/uuid"
	"log"
)

type OutboxRepository struct {
	exec db.ISqlExecutor
}

func NewOutboxRepository(exec db.ISqlExecutor) *OutboxRepository {
	return &OutboxRepository{exec}
}

func (r *OutboxRepository) AddEvent(ctx context.Context, dto models.Event) error {
	query := `INSERT INTO outbox (id, type, payload, source) VALUES ($1, $2, $3, $4)`
	_, err := r.exec.ExecContext(ctx, query, dto.ID, dto.Type, dto.Payload, dto.Source)
	return err
}
func (r *OutboxRepository) Get(ctx context.Context, eventID uuid.UUID) (*models.Event, error) {
	query := `SELECT id, type, payload, created_at, published_at, source FROM outbox WHERE id = $1`
	var dto models.Event
	err := r.exec.QueryRowContext(
		ctx,
		query,
		eventID,
	).Scan(
		&dto.ID,
		&dto.Type,
		&dto.Payload,
		&dto.CreatedAt,
		&dto.PublishedAt,
		&dto.Source,
	)
	return &dto, err
}

func (r *OutboxRepository) GetUnpublish(ctx context.Context, maxEntries int) (events []*models.Event, err error) {
	query := `SELECT id, type, payload, created_at, published_at, source
			  FROM outbox
			  WHERE published_at IS NULL
			  ORDER BY created_at
			  FOR UPDATE SKIP LOCKED
			  LIMIT $1`
	rows, err := r.exec.QueryContext(ctx, query, maxEntries)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}()
	for rows.Next() {
		var dto models.Event
		err = rows.Scan(
			&dto.ID,
			&dto.Type,
			&dto.Payload,
			&dto.CreatedAt,
			&dto.PublishedAt,
			&dto.Source,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &dto)
	}
	return events, nil
}
func (r *OutboxRepository) WithTx(tx *sql.Tx) repos.IOutboxRepository {
	return &OutboxRepository{tx}
}
