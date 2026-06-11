package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	wg := sync.WaitGroup{}

	c := make(chan int)

	for i := 0; i < 4; i++{
		wg.Add(1)
		go Worker(c,&wg, i)
	}


	for j:=0; j<20; j++{
		c <- j
	}

	close(c)
	wg.Wait()


}

func Worker(c chan int,wg *sync.WaitGroup, id int ){
	defer wg.Done()
	fmt.Println("WorkerId"," ----- ","Value")
	for val := range c{
		time.Sleep(2*time.Second)
		fmt.Println(id,"----------", val)
	}
}