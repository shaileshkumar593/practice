package service

import (
	"boilerplate/internal/model"
	"boilerplate/internal/repository"
	"context"
)

type IService interface {
	GetUser(ctx context.Context, id int64) (model.User, error)
	CreateUser(ctx context.Context, name, email string) (model.User, error)
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) *service {
	return &service{repo: repo}
}
