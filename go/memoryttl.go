package main



import (

	"fmt"

	"sync"

	"time"

)



type Value struct {

	Data      any

	ExpiresAt time.Time

}



type TSafeMemory struct {

	mu   sync.RWMutex

	data map[string]Value

}



func NewTSafeMemory() *TSafeMemory {

	return &TSafeMemory{

		data: make(map[string]Value),

	}

}



// Store key with custom TTL

func (m *TSafeMemory) Set(key string, value any, ttl time.Duration) {

	if key == "" {

		return

	}



	m.mu.Lock()

	defer m.mu.Unlock()



	m.data[key] = Value{

		Data:      value,

		ExpiresAt: time.Now().Add(ttl),

	}

}



// Read value

func (m *TSafeMemory) Get(key string) (any, bool) {



	m.mu.RLock()

	v, ok := m.data[key]

	m.mu.RUnlock()



	if !ok {

		return nil, false

	}



	// TTL expired

	if time.Now().After(v.ExpiresAt) {



		m.mu.Lock()

		delete(m.data, key)

		m.mu.Unlock()



		return nil, false

	}



	return v.Data, true

}



// Delete key

func (m *TSafeMemory) Delete(key string) {



	m.mu.Lock()

	defer m.mu.Unlock()



	delete(m.data, key)

}



// Background cleanup

func (m *TSafeMemory) StartCleanup(interval time.Duration) {



	ticker := time.NewTicker(interval)



	go func() {

		defer ticker.Stop()



		for range ticker.C {



			now := time.Now()



			m.mu.Lock()



			for k, v := range m.data {



				if now.After(v.ExpiresAt) {

					delete(m.data, k)

				}

			}



			m.mu.Unlock()

		}

	}()

}



func main() {



	cache := NewTSafeMemory()



	// Cleanup every 5 seconds

	cache.StartCleanup(5 * time.Second)



	// Store value for 10 seconds

	cache.Set("user1", "Shailesh", 10*time.Second)



	val, ok := cache.Get("user1")



	fmt.Println(val, ok)



	time.Sleep(12 * time.Second)



	val, ok = cache.Get("user1")



	fmt.Println(val, ok)

}