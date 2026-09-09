package repository

import (
	"boilerplate/internal/model"
	"context"
	"database/sql"
)

type IRepository interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Compile-time verification
var _ IRepository = (*Repository)(nil)
