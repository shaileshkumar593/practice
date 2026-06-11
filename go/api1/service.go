package main

import (
	"errors"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserService struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *UserService) GetUsers() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]User, 0, len(s.users))

	for _, u := range s.users {
		users = append(users, u)
	}

	return users
}

func (s *UserService) CreateUser(name string) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:   s.nextID,
		Name: name,
	}

	s.users[user.ID] = user
	s.nextID++

	return user
}

func (s *UserService) UpdateUser(id int, name string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[id]
	if !ok {
		return User{}, errors.New("user not found")
	}

	user.Name = name
	s.users[id] = user

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}
