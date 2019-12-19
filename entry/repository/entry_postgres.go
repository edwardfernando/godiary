package repository

import (
	"github.com/edwardfernando/godiary/entry"
	"github.com/jinzhu/gorm"
)

type entryPostgresRepository struct {
	db *gorm.DB
}

// NewEntryRepository initialises the repository with Postgres
func NewEntryRepository(db *gorm.DB) entry.Repository {
	return &entryPostgresRepository{
		db: db,
	}
}
