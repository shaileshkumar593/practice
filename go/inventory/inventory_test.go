package inventory

import (
	"sync"
	"testing"
)

func TestReserveSuccess(t *testing.T) {
	inv := NewInventory(map[string]int{
		"itemA": 10,
	})

	ok := inv.Reserve("itemA", 3)
	if !ok {
		t.Fatalf("expected reservation to succeed")
	}
}

func TestReserveFailsWhenInsufficientStock(t *testing.T) {
	inv := NewInventory(map[string]int{
		"itemA": 5,
	})

	ok := inv.Reserve("itemA", 10)
	if ok {
		t.Fatalf("expected reservation to fail due to insufficient stock")
	}
}

func TestReleaseRestoresStock(t *testing.T) {
	inv := NewInventory(map[string]int{
		"itemA": 5,
	})

	if !inv.Reserve("itemA", 3) {
		t.Fatalf("expected initial reserve to succeed")
	}

	inv.Release("itemA", 3)

	if !inv.Reserve("itemA", 5) {
		t.Fatalf("expected reserve after release to succeed")
	}
}

func TestConcurrentReservations(t *testing.T) {
	inv := NewInventory(map[string]int{
		"itemA": 10,
	})

	var wg sync.WaitGroup
	successCount := 0
	var mu sync.Mutex

	workers := 20
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			if inv.Reserve("itemA", 1) {
				mu.Lock()
				successCount++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	if successCount != 10 {
		t.Fatalf("expected exactly 10 successful reservations, got %d", successCount)
	}
}

func TestInvalidQuantities(t *testing.T) {
	inv := NewInventory(map[string]int{
		"itemA": 5,
	})

	if inv.Reserve("itemA", 0) {
		t.Fatalf("expected reserve with zero quantity to fail")
	}

	if inv.Reserve("itemA", -1) {
		t.Fatalf("expected reserve with negative quantity to fail")
	}

	inv.Release("itemA", -5)
	inv.Release("itemA", 0)
}
