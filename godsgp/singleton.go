package main

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			AppName: "Payment Service",
		}
	})

	return instance
}

func main() {
	c1 := GetConfig()
	c2 := GetConfig()

	fmt.Println(c1 == c2) // true
}

/*

	Singleton ensures that a resource has one shared instance throughout the application.

Typical examples:

Configuration

Logger

Metrics registry

Connection manager

Application-wide cache


Why sync.Once?
A naive Singleton can have a race condition:

if instance == nil {
	instance = &Config{}
}
Multiple goroutines can execute this simultaneously.

sync.Once guarantees initialization happens exactly once.


"The naive Singleton has a race condition because the check instance == nil and the assignment
 instance = &Config{} are not atomic as a combined operation. Two goroutines can both see nil
 and create separate instances. We use sync.Once or proper synchronization to make initialization thread-safe."

*/
