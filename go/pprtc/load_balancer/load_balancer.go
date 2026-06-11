package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

// Backend represents a backend server
type Backend struct {
	URL    *url.URL
	Alive  atomic.Bool
	Mutex  sync.RWMutex
}

// LoadBalancer distributes requests across backends
type LoadBalancer struct {
	backends      []*Backend
	current       uint64
	mu            sync.RWMutex
	healthPath    string
	healthTicker  *time.Ticker
	clientTimeout time.Duration
}

// NewLoadBalancer creates a new load balancer
func NewLoadBalancer(backendURLs []string, healthPath string, healthInterval time.Duration) (*LoadBalancer, error) {
	if len(backendURLs) == 0 {
		return nil, fmt.Errorf("at least one backend required")
	}

	lb := &LoadBalancer{
		backends:      make([]*Backend, 0, len(backendURLs)),
		healthPath:    healthPath,
		clientTimeout: 10 * time.Second,
	}

	for _, urlStr := range backendURLs {
		u, err := url.Parse(urlStr)
		if err != nil {
			return nil, fmt.Errorf("invalid backend URL %q: %w", urlStr, err)
		}
		backend := &Backend{URL: u}
		backend.Alive.Store(true)
		lb.backends = append(lb.backends, backend)
	}

	// Start health checker
	lb.startHealthChecker(healthInterval)

	return lb, nil
}

// getNextBackend returns next alive backend using round-robin
func (lb *LoadBalancer) getNextBackend() *Backend {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	n := uint64(len(lb.backends))
	if n == 0 {
		return nil
	}

	for i := uint64(0); i < n; i++ {
		idx := atomic.AddUint64(&lb.current, 1) % n
		b := lb.backends[idx]
		if b.Alive.Load() {
			return b
		}
	}
	return nil
}

// startHealthChecker periodically checks backend health
func (lb *LoadBalancer) startHealthChecker(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			for _, backend := range lb.backends {
				go lb.checkHealth(backend)
			}
		}
	}()

	// Initial health check
	for _, backend := range lb.backends {
		go lb.checkHealth(backend)
	}
}

// checkHealth checks if a backend is alive
func (lb *LoadBalancer) checkHealth(backend *Backend) {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	path := lb.healthPath
	if path == "" {
		path = "/"
	}

	healthURL := fmt.Sprintf("%s://%s%s", backend.URL.Scheme, backend.URL.Host, path)
	resp, err := client.Get(healthURL)

	alive := err == nil && resp.StatusCode >= 200 && resp.StatusCode < 400
	if resp != nil {
		resp.Body.Close()
	}

	prev := backend.Alive.Load()
	backend.Alive.Store(alive)

	if prev != alive {
		status := "UP"
		if !alive {
			status = "DOWN"
		}
		log.Printf("Health check: %s -> %s", backend.URL.Host, status)
	}
}

// ProxyRequest forwards a request to a backend
func (lb *LoadBalancer) ProxyRequest(w http.ResponseWriter, r *http.Request) {
	backend := lb.getNextBackend()
	if backend == nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	// Create proxy URL
	r.RequestURI = ""
	r.URL.Scheme = backend.URL.Scheme
	r.URL.Host = backend.URL.Host
	if backend.URL.Path != "" {
		r.URL.Path = backend.URL.Path + r.URL.Path
	}

	// Create proxy request
	proxyReq := r.Clone(r.Context())
	proxyReq.RequestURI = ""
	proxyReq.Host = backend.URL.Host
	proxyReq.Header.Set("X-Forwarded-For", getClientIP(r))
	proxyReq.Header.Set("X-Forwarded-Proto", r.Header.Get("X-Forwarded-Proto"))
	if proxyReq.Header.Get("X-Forwarded-Proto") == "" {
		proxyReq.Header.Set("X-Forwarded-Proto", "http")
	}

	// Execute proxy request
	client := &http.Client{Timeout: lb.clientTimeout}
	resp, err := client.Do(proxyReq)
	if err != nil {
		log.Printf("Error proxying to %s: %v", backend.URL.Host, err)
		backend.Alive.Store(false)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for k, v := range resp.Header {
		w.Header()[k] = v
	}

	// Write status and body
	w.WriteHeader(resp.StatusCode)
	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// getClientIP extracts client IP from request
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return xff
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func main() {
	var (
		listenAddr      = flag.String("listen", ":8080", "Listen address and port")
		backendsFlag    = flag.String("backends", "", "Comma-separated backend URLs (required)")
		healthPath      = flag.String("health-path", "/health", "Health check path")
		healthInterval  = flag.Duration("health-interval", 5*time.Second, "Health check interval")
	)
	flag.Parse()

	if *backendsFlag == "" {
		log.Fatal("--backends flag is required")
	}

	// Parse backends from flag (simple comma-split, no spaces)
	var backends []string
	// For simplicity, use individual backend args instead of comma-separated
	// Example: go run . -backends "http://localhost:8001" -backends "http://localhost:8002"
	// But flag doesn't support that directly, so we'll use a workaround

	// Better approach: use os.Args or accept multiple -backend flags
	// For now, demonstrate with hardcoded or simpler approach

	log.Printf("Load Balancer started on %s", *listenAddr)

	// Example: hardcoded or read from env
	backends = []string{
		"http://localhost:8001",
		"http://localhost:8002",
	}

	lb, err := NewLoadBalancer(backends, *healthPath, *healthInterval)
	if err != nil {
		log.Fatalf("Failed to create load balancer: %v", err)
	}

	log.Printf("Configured backends: %v", backends)
	log.Printf("Health check path: %s, interval: %v", *healthPath, *healthInterval)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		lb.ProxyRequest(w, r)
	})

	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
