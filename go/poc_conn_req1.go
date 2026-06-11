
package main 

import (
	"fmt"
	"sync/atomic"
	"net/http"
	"time"
)

var requestCount1 int64

type PingHandler struct{}

func (p PingHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	count := atomic.AddInt64(&requestCount1,1)
	fmt.Printf("Ping request %d\n",count)

	time.Sleep(2*time.Second)

	fmt.Fprint(w,"Pong response for request %d", count)
}

type HelloHandler struct{}
func(h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	count := atomic.AddInt64(&requestCount1,1)
	fmt.Printf("HELLO request %d\n", count)

	fmt.Fprintf(w, "Hello from request %d\n", count)
}

type TimemerHandler struct{}

func (t TimemerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	count := atomic.AddInt64(&requestCount1, 1)
	fmt.Printf("TIME request %d\n", count)

	fmt.Fprintf(w, "Current time: %s (request %d)\n", time.Now().Format(time.RFC3339), count)
}


type StatusHandler struct{}

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count := atomic.AddInt64(&requestCount1, 1)
	fmt.Printf("STATUS request %d\n", count)

	fmt.Fprintf(w, "Server is running (request %d)\n", count)
}

func main(){
	mux  := http.NewServeMux()

	mux.Handle("/ping", PingHandler{})
	mux.Handle("/hello", HelloHandler{})
	mux.Handle("/time", TimemerHandler{})
	mux.Handle("/status", StatusHandler{})

	err := http.ListenAndServe(":8080", mux)
	/*
	for{
		conn := listner.Accept()
		go serveConnection(conn)
	}*/

	if err != nil{
		panic(err)
	}
}