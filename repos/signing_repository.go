package repos

import (
	"context"
	db "github.com/OfficialEvsty/rt-data/db/interface"
	"github.com/OfficialEvsty/rt-data/models"
)

type SigningRepository struct {
	exec db.ISqlExecutor
}

func NewSigningRepository(exec db.ISqlExecutor) *SigningRepository {
	return &SigningRepository{exec}
}

func (r *SigningRepository) SaveSigningInfo(ctx context.Context, info models.SigningInfo) error {
	query := `INSERT INTO signing_info (id, author, signing, signed_at) VALUES ($1, $2, $3, $4)`
	_, err := r.exec.ExecContext(ctx, query,
		info.ID,
		info.Author,
		info.Signing,
		info.SignedAt,
	)
	return err
}
