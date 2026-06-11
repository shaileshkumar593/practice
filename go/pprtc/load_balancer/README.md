# Go Load Balancer

A simple HTTP reverse-proxy load balancer written in Go with round-robin distribution and health checking.

## Features

- **Round-robin load distribution** across multiple backends
- **Health checking** with configurable intervals
- **Automatic failover** - removes unhealthy backends from rotation
- **Request forwarding** - preserves headers, method, and body
- **Logging** - tracks all requests and backend status changes
- **No external dependencies** - uses only Go standard library

## Project Structure

```
load_balancer/
├── load_balancer.go     # Main load balancer logic
├── backend_example.go   # Example backend server
├── go.mod              # Go module definition
└── README.md           # This file
```

## Quick Start

### Prerequisites

- Go 1.21 or later

### 1. Build the application

```bash
go build -o loadbalancer.exe load_balancer.go
go build -o backend.exe backend_example.go
```

Or compile both at once:

```bash
go build -o loadbalancer.exe load_balancer.go; go build -o backend.exe backend_example.go
```

### 2. Start two backend servers

**Terminal 1:**
```bash
.\backend.exe -port 8001 -id "backend-1"
```

**Terminal 2:**
```bash
.\backend.exe -port 8002 -id "backend-2"
```

### 3. Start the load balancer

**Terminal 3:**
```bash
.\loadbalancer.exe -listen ":8080" -health-path "/health" -health-interval 5s
```

### 4. Test with requests

**Terminal 4:**
```bash
# Test with curl (requires curl installed)
curl http://localhost:8080/

# Or use PowerShell:
(Invoke-WebRequest -Uri http://localhost:8080/).Content

# Multiple requests to see round-robin in action
for ($i=1; $i -le 5; $i++) { (Invoke-WebRequest -Uri http://localhost:8080/).Headers['X-Served-By'] }
```

## Configuration

### Load Balancer Flags

- `-listen` (default: `:8080`) - Address and port to listen on
- `-backends` (required) - Comma-separated backend URLs
- `-health-path` (default: `/health`) - Path for health checks
- `-health-interval` (default: `5s`) - Interval between health checks

### Backend Server Flags

- `-port` (default: `8001`) - Port to listen on
- `-id` (default: `server1`) - Server identifier for logging

## How It Works

### Load Distribution

Requests are distributed using **round-robin** algorithm:
1. Maintain a counter that increments with each request
2. Select backend at `counter % number_of_backends`
3. Skip unhealthy backends
4. If all backends are down, return 503 Service Unavailable

### Health Checking

Background goroutine periodically checks each backend:
1. Makes HTTP GET request to `{backend}/health`
2. Expects 2xx or 3xx status code
3. If health check fails, backend is marked DOWN
4. Logs status changes

### Request Flow

```
Client Request
    ↓
Load Balancer
    ↓ (round-robin select healthy backend)
Backend Server
    ↓
Response forwarded to client
```

## Example Output

**Load Balancer:**
```
2025/12/04 10:15:32 Load Balancer started on :8080
2025/12/04 10:15:32 Configured backends: [http://localhost:8001 http://localhost:8002]
2025/12/04 10:15:32 Health check path: /health, interval: 5s
2025/12/04 10:15:37 Health check: localhost:8001 -> UP
2025/12/04 10:15:37 Health check: localhost:8002 -> UP
2025/12/04 10:15:45 127.0.0.1:54321 GET /
2025/12/04 10:15:46 127.0.0.1:54322 GET /
```

**Backend Server:**
```
2025/12/04 10:15:32 [backend-1] Starting backend server on :8001
2025/12/04 10:15:45 [backend-1] GET / -> 200 (total: 1)
```

## Testing Scenarios

### 1. Verify Round-Robin Distribution

Send multiple requests and check which backend served each:

```bash
for ($i=1; $i -le 10; $i++) { 
    $response = Invoke-WebRequest -Uri http://localhost:8080/
    $response.Headers['X-Served-By']
}
```

Expected: Alternates between `backend-1` and `backend-2`

### 2. Test Failover

Stop one backend server and send requests:

```bash
# Stop backend-1, send requests
for ($i=1; $i -le 5; $i++) { 
    $response = Invoke-WebRequest -Uri http://localhost:8080/
    $response.Headers['X-Served-By']
}
```

Expected: All requests served by `backend-2`

### 3. Health Check Verification

Check logs for health status messages when backends go up/down.

## Code Highlights

### Backend Selection
```go
func (lb *LoadBalancer) getNextBackend() *Backend {
    // Round-robin with atomic counter
    // Skips unhealthy backends
}
```

### Health Checking
```go
func (lb *LoadBalancer) checkHealth(backend *Backend) {
    // GET {backend}/health
    // Mark alive/dead based on response
}
```

### Request Proxying
```go
func (lb *LoadBalancer) ProxyRequest(w http.ResponseWriter, r *http.Request) {
    // Select backend
    // Forward request with proper headers
    // Stream response back to client
}
```

## Performance Notes

- Uses goroutines for concurrent request handling
- Atomic operations for lock-free backend selection
- Connection pooling via http.Client
- Efficient health checking with background goroutine

## Limitations

- Simple round-robin (no weighted distribution)
- No session affinity / sticky sessions
- No request/response modification beyond headers
- Backend URL modification is basic
- No metrics/monitoring endpoint

## Extensions

You could enhance this with:
- Weighted round-robin
- Least-connections algorithm
- Session persistence
- Request rate limiting
- Metrics endpoint
- Configuration reload without restart
- TLS/HTTPS support
- Request/response logging

## License

MIT
