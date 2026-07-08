package main


import (
	"context"
	"fmt"
	"time"
)

func PrintNumbers(ctx context.Context){
	var i int = 0

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Go routine")
			return
		default:
			<-ticker.C
			i++
			println(i)
		}
	}
}

func main(){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go PrintNumbers(ctx)

	<-ctx.Done()

	fmt.Println("over")
}