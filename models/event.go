package models

import (
	"encoding/json"
	"github.com/google/uuid"
	"time"
)

type EventType string

const (
	// SignRequestedEvent порождающий цепочку последующих событий подписания документа
	SignRequestedEvent        EventType = "signrequested"
	DocumentGeneratedEvent    EventType = "documentgenerated"
	ServiceSignRequestedEvent EventType = "servicesignrequested"
)

// Event описывает тип, время и параметры определенного события
// Является основной движущей частью event-driven подхода
type Event struct {
	ID          uuid.UUID       `json:"id"`
	Type        EventType       `json:"type"`
	Payload     json.RawMessage `json:"payload"`
	CreatedAt   time.Time       `json:"created_at"`
	PublishedAt *time.Time      `json:"published_at"`
	Source      *uuid.UUID      `json:"source"`
}
