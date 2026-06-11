/*In Go, there is no yield keyword like in Python or C#. Go uses its native concurrency primitives—goroutines and channels—to achieve the same functionality as generators and iterators in other languages. 
This approach allows a function to produce a sequence of values asynchronously, and the caller can consume them as they become available. This is particularly useful for processing large data sets without loading everything into memory at once. 
How to use channels to simulate yield
A common pattern for creating a generator in Go involves these steps: 
Define a function that returns a channel.
Inside this function, start a goroutine.
In the goroutine, a loop sends values to the channel.
After the loop finishes, the goroutine closes the channel to signal that no more values will be sent.
The main function then uses a for...range loop to read values from the channel until it is closed. 
Example: A number generator
This code demonstrates a generator that produces a sequence of numbers from a specified range. 
go*/


package main

import (
	"fmt"
)

// The generator function returns a channel.
// The go routine sends values to this channel.
func generateNumbers(from, to int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // Always close the channel to signal the end of the sequence.
		for i := from; i <= to; i++ {
			ch <- i // Send a value into the channel.
		}
	}()
	return ch
}

func main() {
	// The caller uses a for...range loop to receive values
	// from the channel until it is closed.
	for num := range generateNumbers(1, 5) {
		fmt.Printf("Received: %d\n", num)
	}
}
/*Use code with caution.

Output:
sh
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Use code with caution.*/


/*Advantages of the Go approach
Lazy evaluation: The generator produces values only as they are needed by the consumer, which is ideal for large or infinite data streams.
Clear separation of concerns: The code that produces the data (the generator goroutine) is cleanly separated from the code that consumes it (the for...range loop).
Built-in back-pressure: If the consumer slows down, the channel blocks the generator goroutine from sending more values. This automatically regulates the flow of data.
Concurrency: The generation of data can happen concurrently with its consumption, potentially leveraging multiple CPU cores. 

go*/

