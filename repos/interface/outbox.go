package repos

import (
	"context"
	"database/sql"
	"github.com/OfficialEvsty/rt-data/models"
	"github.com/google/uuid"
)

// IOutboxRepository repo interface for event handling
type IOutboxRepository interface {
	AddEvent(context.Context, models.Event) error
	Get(context.Context, uuid.UUID) (*models.Event, error)
	WithTx(*sql.Tx) IOutboxRepository
}
