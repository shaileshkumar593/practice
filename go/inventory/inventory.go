package inventory

import (
	"sync"
)

type Inventory struct {
	inventory map[string]int
	mu        sync.Mutex
}

func NewInventory(initialStock map[string]int) *Inventory {
	// Copy map to avoid external mutation
	invMap := make(map[string]int)
	for k, v := range initialStock {
		invMap[k] = v
	}

	return &Inventory{
		inventory: invMap,
	}
}

// Reserve attempts to reserve qty units of itemID.
// Returns true if reservation succeeds, false otherwise.
func (i *Inventory) Reserve(itemID string, qty int) bool {
	// ❗ Invalid quantity check
	if qty <= 0 {
		return false
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	// ❗ Check if item exists
	current, exists := i.inventory[itemID]
	if !exists {
		return false
	}

	// ❗ Check stock
	if current < qty {
		return false
	}

	// Reserve
	i.inventory[itemID] = current - qty
	return true
}

// Release returns qty units of itemID back to inventory.
func (i *Inventory) Release(itemID string, qty int) {
	// ❗ Ignore invalid quantity (as per test)
	if qty <= 0 {
		return
	}

	i.mu.Lock()
	defer i.mu.Unlock()

	// If item doesn't exist, initialize it
	i.inventory[itemID] += qty
}
