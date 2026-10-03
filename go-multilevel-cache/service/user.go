package service

import (
	"context"
	"encoding/json"
	"fmt"

	"go-multilevel-cache/cache"
	"go-multilevel-cache/repository"
)

type UserService struct {
	repo  *repository.UserRepository
	cache *cache.MultiLevelCache
}

func NewUserService(
	repo *repository.UserRepository,
	cache *cache.MultiLevelCache,
) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (repository.User, string, error) {

	key := fmt.Sprintf("user:%d", id)

	value, source, err := s.cache.Get(
		ctx,
		key,
		func(ctx context.Context) ([]byte, error) {
			user, err := s.repo.GetByID(ctx, id)
			if err != nil {
				return nil, err
			}

			return json.Marshal(user)
		},
	)

	if err != nil {
		return repository.User{}, "", err
	}

	var user repository.User
	if err := json.Unmarshal(value, &user); err != nil {
		return repository.User{}, "", err
	}

	return user, source, nil
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	user repository.User,
) error {

	// Database/source of truth first.
	if err := s.repo.Update(ctx, user); err != nil {
		return err
	}

	// Then invalidate L1 + L2.
	key := fmt.Sprintf("user:%d", user.ID)
	return s.cache.Delete(ctx, key)
}
