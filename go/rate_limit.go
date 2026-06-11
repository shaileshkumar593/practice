package make(

	import (
		"fmt"
		"net/http"
		"sync"
		"time"
	)

var(
	mu sync.RWMutex
	requests = 0
	limit = 5
)

func rateLimiter(next http.HandlerFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.request){
		mu.Lock()
		request++

		if request > limit{
			mu.Unlock()
			http.Error(w,"too many request", http.StatsuToomanyRequests)
			return
		}

		mu.Unlock()
		next(w,r)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"pong")
}

func main()){
	go func(){
		for{
		time.Sleep(1* time.Minute)
		mu.Lock()
		request=0
		mu.Unlock()
		}
	}()

	mux := http.NewServeMux()

	mux.handleFunc("/ping", rateLimiter(pingHandler))

	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", mux)
	
	if err != nil{
		panic(err)
	}
}