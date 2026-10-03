package repository

import (
	"context"
	"sync"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository struct {
	mu    sync.RWMutex
	users map[int]User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: map[int]User{
			101: {
				ID:    101,
				Name:  "Alice",
				Email: "alice@example.com",
			},
			102: {
				ID:    102,
				Name:  "Bob",
				Email: "bob@example.com",
			},
		},
	}
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (User, error) {
	// Simulates database latency.
	select {
	case <-time.After(200 * time.Millisecond):
	case <-ctx.Done():
		return User{}, ctx.Err()
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		// Keep the example simple.
		return User{
			ID:    id,
			Name:  "Generated User",
			Email: "user@example.com",
		}, nil
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID] = user
	return nil
}
