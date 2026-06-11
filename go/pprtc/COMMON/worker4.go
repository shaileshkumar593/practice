package main

/*import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var (
	webPages = []string{
		"yahoo.com",
		"google.com",
		"bing.com",
		"amazon.com",
		"github.com",
		"gitlab.com",
	}

	results struct {
		// put here content length of each page
		ContentLength map[string]int

		// accumulate here the content length of all pages
		TotalContentLength int
	}
)

func workers(ctx context.Context, id int, wg *sync.WaitGroup, jobs <-chan int, out chan<- int, errors chan<- error, wp chan string) {
	defer func(context.Context) {
		r := recover()
		if r != nil {
			fmt.Println(r)
			select {
			case <-ctx.Done():
				fmt.Println("Timeout")
				return
			default:
				value := ctx.Value("key")
				fmt.Println(value)
			}
		}
	}(ctx)
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)

		fmt.Println("URL  :", <-wp)
		url := <-wp
		res, err := http.Get(url)
		if err != nil {
			// handle error
			fmt.Sprintln("error received from http request", err)
			results.ContentLength[url] = -1
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			// handle error
			fmt.Sprintln("error during parsing response ", err)
			results.ContentLength[url] = -1
		}
		results.ContentLength[url] = len(body)

	}
	wg.Done()
}

func main() {
	jobs := make(chan int, len(webPages))
	out := make(chan int, len(webPages))
	errors := make(chan error, len(webPages))
	wp := make(chan string, len(webPages))

	results.ContentLength = make(map[string]int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		go workers(ctx, w, &wg, jobs, out, errors, wp)
	}

	for j := 1; j <= len(webPages); j++ {
		jobs <- j
		wp <- webPages[j-1]
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()

	select {

	case <-ctx.Done():
		fmt.Println("Took too long!")

	}
}
*/
