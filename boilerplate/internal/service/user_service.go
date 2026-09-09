package service

import (
	"boilerplate/internal/model"
	"context"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInvalidUser = errors.New("invalid user")
)

func (s *service) GetUser(ctx context.Context, id int64) (model.User, error) {
	if id <= 0 {
		return model.User{}, ErrInvalidUser
	}
	return s.repo.GetByID(ctx, id)
}

func (s *service) CreateUser(ctx context.Context, name, email string) (model.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" || !strings.Contains(email, "@") {
		return model.User{}, ErrInvalidUser
	}

	user, err := s.repo.Create(ctx, model.User{Name: name, Email: email})
	if err != nil {
		return model.User{}, fmt.Errorf("create user service: %w", err)
	}

	return user, nil
}
