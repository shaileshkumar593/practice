package main 


import(
	"fmt"
	//"time"

	"sync"
)

// 
func consumer(c chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for val := range c{
		//p := <-c
		fmt.Println(val)
	}
	
}

func producer(c chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for i:= 0; i < 5; i++{
		c <-i
	}

	
}

func main(){
	c := make(chan int)
	var wg,wg1  sync.WaitGroup

	for i:= 0; i < 3; i++{
		wg.Add(1)
		go producer(c, &wg)

	}
	for i := 0; i<3; i++{
		wg1.Add(1)
		go consumer(c, &wg1)
	}
	go func(){
		wg.Wait()
		close(c)
	}()
	

	wg1.Wait()  //---- 
	//time.Sleep(2 *time.Second)


}