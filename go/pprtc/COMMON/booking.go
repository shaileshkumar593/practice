
import (
	"fmt"
	"sync"
)

type SeatBooking struct {
	Seat  int
	Mutex sync.Mutex
	Uname chan string
}

func (t SeatBooking) SeatBook(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer t.Mutex.Unlock()
	t.Mutex.Lock()
	if t.Seat <= 0 {
		fmt.Println("No seats available")
	} else {
		t.Seat = t.Seat - 1
		t.Uname <- s
	}

}

func (t SeatBooking) PrintInvoice(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Seat booked for User ", <-t.Uname)
}

func main() {
	wgrp := sync.WaitGroup{}

	m := sync.Mutex{}

	c := make(chan string)

	t := SeatBooking{4, m, c}

	name := []string{"Ramesh", "Ganesh", "Suresh"}

	for i := 0; i < 3; i++ {
		wgrp.Add(1)
		go t.SeatBook(name[i], &wgrp)

	}
	wgrp.Add(1)
	go t.PrintInvoice(&wgrp)

	wgrp.Wait()

}