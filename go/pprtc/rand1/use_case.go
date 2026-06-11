package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
	Load balancing (random backend selection)

	Use case: API Gateway, Service Discovery
	 Why: Avoids hotspotting, simple & fast
*/

func randomLoadBalancing() {
	backends := []string{
		"10.0.0.1",
		"10.0.0.2",
		"10.0.0.3",
	}

	selected := backends[rand.Intn(len(backends))]
	fmt.Println("Selected backend:", selected)
}

/*
	Random retry backoff (jitter)

	Use case: Kafka consumers, K8s controllers
	Why: Prevents thundering herd
*/

func randomRetryBackoff() {
	base := 500 * time.Millisecond
	jitter := time.Duration(rand.Intn(300)) * time.Millisecond

	fmt.Println("Retrying after:", base+jitter)
	time.Sleep(base + jitter)
}

/*
	Shuffling (random order without repetition)

	Use case: Job schedulers, recommendation engines
	Guarantee: No duplication

*/

func shufflingExample() {
	arr := []int{1, 2, 3, 4, 5}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	fmt.Println("Shuffled:", arr)
}

/*
Feature rollout / A-B testing

Use case: Canary deployments

Why: Gradual rollout with controlled risk
*/
func featureRollout() {
	if rand.Intn(100) < 10 { // 10%
		fmt.Println("New feature ENABLED")
	} else {
		fmt.Println("Old feature")
	}
}

/*
Random ID generation (non-secure)
Never use for auth / tokens

	Use crypto/rand instead
*/
func randomIDGeneration() {
	id := rand.Int63()
	fmt.Println("Temporary ID:", id)
}

// Simulations & probabilistic systems
func probabilisticSimulation() {
	if rand.Float64() < 0.7 {
		fmt.Println("Event occurred (70%)")
	} else {
		fmt.Println("Event did not occur")
	}
}

/*
 Use case: Monte Carlo, traffic modeling
 Why: Real-world randomness
*/

// Random leader election (simple systems)
func randomLeaderElection() {
	nodes := []string{"node-1", "node-2", "node-3"}

	leader := nodes[rand.Intn(len(nodes))]
	fmt.Println("Elected leader:", leader)
}

/*
 Not for strong consistency
 Use Raft / etcd / Zookeeper
*/

// Testing & fuzzing
func testingAndFuzzing() {
	for i := 0; i < 5; i++ {
		fmt.Println("Random input:", rand.Intn(1000))
	}
}

//Use case: Chaos testing, fault injection

// Cache eviction (random policy)
func randomCacheEviction() {
	cache := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	keys := make([]string, 0, len(cache))
	for k := range cache {
		keys = append(keys, k)
	}

	evict := keys[rand.Intn(len(keys))]
	delete(cache, evict)

	fmt.Println("Evicted key:", evict)
	fmt.Println("Cache now:", cache)
}

/*
	Why: Low-overhead eviction
	 Better than LRU in very high throughput systems
*/

// Mock data generation
func mockDataGeneration() {
	age := rand.Intn(60) + 18
	score := rand.Float64() * 100

	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
}

// Use case: Load testing, UI demos

// Concurrency-safe random (advanced, production)
func concurrentSafeRandom() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Random:", r.Intn(100))
		}()
	}
	wg.Wait()
}

/*
Why: Avoid contention on global rand
*/

func main() {
	randomLoadBalancing()
	randomRetryBackoff()
	shufflingExample()
	featureRollout()
	randomIDGeneration()
	probabilisticSimulation()
	randomLeaderElection()
	testingAndFuzzing()
	randomCacheEviction()
	mockDataGeneration()
}
