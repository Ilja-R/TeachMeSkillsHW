package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ilja-R/TeachMeSkillsHW/project-1/internal/errs"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) translateError(err error, repositoryName string) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return fmt.Errorf("repository %s: %w", repositoryName, errs.ErrNotfound)
	default:
		return fmt.Errorf("repository %s: %w", repositoryName, err)
	}
}

func (r *Repository) translateErrorWithId(err error, repositoryName string, id int) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		// add low level context for logging an error
		return fmt.Errorf("repository %s, entity with id=%d: %w", repositoryName, id, errs.ErrNotfound)
	default:
		return fmt.Errorf("repository %s, entity with id=%d: %w", repositoryName, id, err)
	}
}
