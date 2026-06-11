package store

import (
	"errors"
	"sync"
	"time	uuid"
)

// Order model used internally
type Order struct {
	ID        string
	Items     []string
	Price     float64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var (
	ErrNotFound = errors.New("order not found")
)

// InMemoryStore stores orders in a slice protected by RWMutex.
// This is intentionally simple for demonstration/testing.
type InMemoryStore struct {
	mu     sync.RWMutex
	orders []Order
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{orders: make([]Order, 0, 16)}
}

func (s *InMemoryStore) Create(items []string, price float64) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := uuid.New().String()
	now := time.Now().UTC()
	ord := Order{
		ID:        id,
		Items:     items,
		Price:     price,
		Status:    "created",
		CreatedAt: now,
		UpdatedAt: now,
	}
	s.orders = append(s.orders, ord)
	return ord, nil
}

func (s *InMemoryStore) Get(id string) (Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, o := range s.orders {
		if o.ID == id {
			return o, nil
		}
	}
	return Order{}, ErrNotFound
}

func (s *InMemoryStore) Update(id string, items []string, price float64, status string) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range s.orders {
		if s.orders[i].ID == id {
			s.orders[i].Items = items
			s.orders[i].Price = price
			s.orders[i].Status = status
			s.orders[i].UpdatedAt = time.Now().UTC()
			return s.orders[i], nil
		}
	}
	return Order{}, ErrNotFound
}
