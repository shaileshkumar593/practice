package main


func main(){
	ch := make(chan int, 2)

	ch <- 10

	ch <-20

	ch<-40 // cause deadlock

}