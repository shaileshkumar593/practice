package main


import (
	"context"
	"fmt"
	"time"
)

func PrintNumbers(ctx context.Context){
	var i int = 0
	for {
		select{
		case <-ctx.Done():
			fmt.Println("Go routine")
			return
		default:
			i++
			print(i)
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