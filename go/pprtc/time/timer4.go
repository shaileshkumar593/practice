package main 


import(
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	for {

		select {

		case t := <-ticker.C:
			fmt.Println("Tick:", t.Format("15:04:05"))

		case <-timer.C:
			fmt.Println("5 seconds completed")
			return
		}
	}
}