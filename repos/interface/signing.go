package repos

import (
	"context"
	"github.com/OfficialEvsty/rt-data/models"
)

// ISigningRepository репозиторий для сохранения информации о подписанных документах
type ISigningRepository interface {
	SaveSigningInfo(ctx context.Context, info models.SigningInfo) error
}
