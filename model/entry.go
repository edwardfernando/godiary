package model

import (
	"time"

	"github.com/google/uuid"
)

// Entry represents diary entry
type Entry struct {
	ID        uuid.UUID `db:"id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
