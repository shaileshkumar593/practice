package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gammazero/workerpool"
)

var wps *workerpool.WorkerPool

func allocate() {
	wps = workerpool.New(10)
}
func main() {
	start := time.Now()
	allocate()
	//fmt.Println(wps)
	wps = workerpool.New(30)
	//fmt.Println(*wps, wps)
	//fmt.Println(wps, *wps)
	// wps = workerpool.New(3)
	// fmt.Println(wps)
	fmt.Println("oh ok")
	wp1 := workerpool.New(0)
	wp := workerpool.New(2)
	if reflect.TypeOf(wp1) == reflect.TypeOf(wp) {
		fmt.Println("Hello")
		// fmt.Println(wp.Size(), wp1.Size())
		// fmt.Println(&wp1)
		// fmt.Println(&wp)

	}
	requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}

	for _, r := range requests {
		r := r
		wp.Submit(func() {
			fmt.Println("Handling request:", r)
		})
	}
	fmt.Println("Before 2sec ", wps.Size(), wp1.Size(), wp.Size())
	wp.StopWait()
	wp.Stop()
	wps.Stop()
	wp1.Stop()

	extime := time.Since(start)
	time.Sleep(2 * time.Second)
	time.Sleep(3 * time.Second)
	fmt.Println("after 2sec", wps.Size(), wp1.Size(), wp.Size())
	fmt.Println(wp.Stopped())
	fmt.Println(wps.Stopped())
	fmt.Println(wp1.Stopped())
	fmt.Println(extime)

	fmt.Println(&wp, &wps, &wp1)
	fmt.Println(*wp, *wps, *wp1)
	for _, r := range requests {
		r := r
		wp.Submit(func() {
			fmt.Println("Handling request:", r)
		})
	}
}
