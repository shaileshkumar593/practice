package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
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

func workerss(ctx context.Context, c string, wg sync.WaitGroup, stop context.CancelFunc) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}

	}()
	url := c
	fmt.Println(url)
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
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		stop()
		//fmt.Println(ctx.Err())
	}
	wg.Done()

}

func main() {
	results.ContentLength = make(map[string]int)

	const shortDuration = 8 * time.Millisecond
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	for _, val := range webPages {
		wg.Add(1)
		go workerss(ctx, val, wg, stop)

	}

	wg.Wait()
	fmt.Println("*********map **********************")
	for key, val := range results.ContentLength {
		fmt.Println(key, " : ", val)
	}
	fmt.Println("Total length", results.TotalContentLength)

	select {
	case <-ctx.Done():
		stop()
		fmt.Println("Took too long!")

	}
}
