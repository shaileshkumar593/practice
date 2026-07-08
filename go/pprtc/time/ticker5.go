package main


import(
	"fmt"
	"time"
)

// production pattern
func main(){
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-timeout:
			fmt.Println("Exactly 5 seconds reached")
			return
		}
	}
}