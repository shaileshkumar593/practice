package service

import (
	"errors"
	"sync"

	"api3/internal/model"
)

type UserService struct {
	mu     sync.RWMutex
	users  map[int]model.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (s *UserService) GetUsers() []model.User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]model.User, 0, len(s.users))

	for _, user := range s.users {
		result = append(result, user)
	}

	return result
}

func (s *UserService) CreateUser(name, website string) model.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := model.User{
		ID:      s.nextID,
		Name:    name,
		Website: website,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user
}

func (s *UserService) UpdateUser(id int, name string) (model.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[id]
	if !ok {
		return model.User{}, errors.New("user not found")
	}

	user.Name = name
	s.users[id] = user

	return user, nil
}
