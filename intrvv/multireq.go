package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
)

var requestCounter int64

// ----------------------------------------------------
// Main
// ----------------------------------------------------

func main() {

	// Equivalent to:
	//
	// http.HandleFunc("/get/", handler)
	//
	// We are manually creating a route table.
	router := NewRouter()

	router.HandleFunc("/get/", handlerss)

	// Equivalent to:
	//
	// http.ListenAndServe(":8080", nil)
	//
	// Internally we first create a TCP listener.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	log.Println("Server listening on :8080")

	// ------------------------------------------------
	// THIS IS THE IMPORTANT LOOP
	// ------------------------------------------------

	for {

		// Wait for a TCP client.
		conn, err := listener.Accept()

		if err != nil {
			log.Println("Accept error:", err)
			continue
		}

		log.Println(
			"New connection:",
			conn.RemoteAddr(),
		)

		// Handle this connection concurrently.
		go handleConnection(conn, router)
	}
}

// ----------------------------------------------------
// Handler
// ----------------------------------------------------

func handlerss(w *ResponseWriter, r *Request) {

	count := atomic.AddInt64(
		&requestCounter,
		1,
	)

	log.Println(
		"Path:",
		r.URL.Path,
		"Count:",
		count,
	)

	w.WriteHeader(200)

	w.Write([]byte("Hello\n"))

	w.Write(
		[]byte(
			fmt.Sprintf(
				"Request count is %d\n",
				count,
			),
		),
	)
}

// ----------------------------------------------------
// Request
// ----------------------------------------------------

type Request struct {
	Method string

	URL URL

	Headers map[string]string
}

// ----------------------------------------------------
// URL
// ----------------------------------------------------

type URL struct {
	Path string
}

// ----------------------------------------------------
// ResponseWriter
// ----------------------------------------------------

type ResponseWriter struct {
	conn net.Conn

	headerWritten bool

	mu sync.Mutex
}

// WriteHeader sends the HTTP status line.
func (w *ResponseWriter) WriteHeader(statusCode int) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if w.headerWritten {
		return
	}

	statusText := "OK"

	if statusCode == 404 {
		statusText = "Not Found"
	}

	response := fmt.Sprintf(
		"HTTP/1.1 %d %s\r\n",
		statusCode,
		statusText,
	)

	response += "Content-Type: text/plain\r\n"
	response += "Connection: close\r\n"
	response += "\r\n"

	_, err := w.conn.Write(
		[]byte(response),
	)

	if err != nil {
		log.Println("Write header error:", err)
		return
	}

	w.headerWritten = true
}

// Write writes response body.
func (w *ResponseWriter) Write(data []byte) {

	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.headerWritten {

		w.mu.Unlock()

		w.WriteHeader(200)

		w.mu.Lock()
	}

	_, err := w.conn.Write(data)

	if err != nil {
		log.Println("Write error:", err)
	}
}

// ----------------------------------------------------
// Router
// ----------------------------------------------------

type Router struct {
	routes []Route
}

type Route struct {
	prefix  string
	handler func(*ResponseWriter, *Request)
}

// NewRouter creates router.
func NewRouter() *Router {

	return &Router{
		routes: make([]Route, 0),
	}
}

// HandleFunc registers a route.
func (r *Router) HandleFunc(
	prefix string,
	handler func(*ResponseWriter, *Request),
) {

	r.routes = append(
		r.routes,
		Route{
			prefix:  prefix,
			handler: handler,
		},
	)
}

// ServeHTTP performs routing.
func (r *Router) ServeHTTP(
	w *ResponseWriter,
	req *Request,
) {

	for _, route := range r.routes {

		if strings.HasPrefix(
			req.URL.Path,
			route.prefix,
		) {

			route.handler(w, req)

			return
		}
	}

	// No route found.
	w.WriteHeader(404)

	w.Write(
		[]byte("404 Not Found\n"),
	)
}

// ----------------------------------------------------
// Connection Handling
// ----------------------------------------------------

func handleConnection(
	conn net.Conn,
	router *Router,
) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	// ------------------------------------------------
	// Request loop
	// ------------------------------------------------
	//
	// A single TCP connection can potentially carry
	// multiple HTTP requests.
	//
	// For simplicity this example closes the
	// connection after one request.
	// ------------------------------------------------

	for {

		request, err := readRequest(reader)

		if err != nil {

			if err.Error() != "EOF" {
				log.Println(
					"Read request error:",
					err,
				)
			}

			return
		}

		log.Println(
			"HTTP request:",
			request.Method,
			request.URL.Path,
		)

		responseWriter := &ResponseWriter{
			conn: conn,
		}

		// Route request.
		router.ServeHTTP(
			responseWriter,
			request,
		)

		// This example uses:
		//
		// Connection: close
		//
		// Therefore we close after one request.
		return
	}
}

// ----------------------------------------------------
// HTTP Request Parser
// ----------------------------------------------------

func readRequest(
	reader *bufio.Reader,
) (*Request, error) {

	// Read request line.
	//
	// Example:
	//
	// GET /get/user HTTP/1.1
	//

	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)

	parts := strings.Split(
		line,
		" ",
	)

	if len(parts) < 3 {

		return nil, fmt.Errorf(
			"invalid request line",
		)
	}

	method := parts[0]

	path := parts[1]

	headers := make(
		map[string]string,
	)

	// Read headers.
	for {

		line, err := reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		line = strings.TrimRight(
			line,
			"\r\n",
		)

		// Empty line means headers are finished.
		if line == "" {
			break
		}

		parts := strings.SplitN(
			line,
			":",
			2,
		)

		if len(parts) == 2 {

			key := strings.TrimSpace(
				parts[0],
			)

			value := strings.TrimSpace(
				parts[1],
			)

			headers[key] = value
		}
	}

	return &Request{
		Method: method,

		URL: URL{
			Path: path,
		},

		Headers: headers,
	}, nil
}
