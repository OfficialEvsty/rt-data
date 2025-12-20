package models

import (
	"github.com/google/uuid"
	"time"
)

type SigningInfo struct {
	ID       uuid.UUID `json:"id"`
	Author   string    `json:"author"`
	Signing  string    `json:"signing"`
	SignedAt time.Time `json:"signed_at"`
}
