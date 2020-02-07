package repository

import (
	"github.com/edwardfernando/godiary/entry"
	"github.com/edwardfernando/godiary/model"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

func (r *entryPostgresRepository) CreateEntry(entry model.Entry) error {
	if err := r.db.Create(&entry).Error; err != nil {
		return errors.Wrapf(err, "entry repository: failed to create entry: %+v", entry)
	}

	return nil
}
