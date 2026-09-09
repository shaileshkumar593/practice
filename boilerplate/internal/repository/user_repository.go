package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"boilerplate/internal/model"
)

var ErrUserNotFound = errors.New("user not found")

func (r *Repository) GetByID(ctx context.Context, id int64) (model.User, error) {
	var u model.User

	err := r.db.QueryRowContext(ctx,
		`SELECT id, name, email, created_at FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return model.User{}, ErrUserNotFound
	}
	if err != nil {
		return model.User{}, fmt.Errorf("get user: %w", err)
	}

	return u, nil
}

func (r *Repository) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.QueryRowContext(ctx,
		`INSERT INTO users (name, email) VALUES ($1, $2)
		 RETURNING id, name, email, created_at`,
		user.Name, user.Email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		return model.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
