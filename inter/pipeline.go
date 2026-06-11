package main

import "fmt"

func generate(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {

	ch := generate(1, 2, 3, 4, 5)

	result := square(ch)

	for v := range result {
		fmt.Println(v)
	}
}

/*

but in function call func Worker(ch1, ch2 <-chan int, out chan<- int)  works fine. Can you expalin diffrences
func Worker(ch1, ch2 <-chan int) <-chan int


Channel Passed As Parameter
func Worker(ch1, ch2 <-chan int, out chan<- int)
Meaning:

ch1 -> receive only
ch2 -> receive only
out -> send only
Inside Worker:

out <- val
✅ Allowed

Outside Worker:

out := make(chan int)

go Worker(ch1, ch2, out)

for v := range out {
	fmt.Println(v)
}
Here:

out
is still:

chan int
(bidirectional)

Only inside the Worker function is it treated as send-only.

Visualization

main

out (chan int)
   |
   |
   +----> Worker sees it as chan<- int
   |
   +----> main sees it as chan int
Channel direction restriction is local to the function parameter.

Version 2: Channel Returned
func Worker(ch1, ch2 <-chan int) <-chan int
Now Worker creates the channel:

out := make(chan int)

return out
Caller gets:

out := Worker(ch1, ch2)
Type of out:

<-chan int
(receive-only)

Allowed:

v := <-out
Not allowed:

out <- 10
Compile error.

Key Difference
Parameter Direction
func Worker(out chan<- int)
restricts Worker.

Caller: full access
Worker: send only
Return Direction
func Worker() <-chan int
restricts Caller.

Worker: full access
Caller: receive only
Example
Parameter Style
func send(ch chan<- int) {
	ch <- 100
}

func main() {

	ch := make(chan int)

	go send(ch)

	fmt.Println(<-ch)
}
Output:

100
Return Style
func generate() <-chan int {

	ch := make(chan int)

	go func() {
		ch <- 100
		close(ch)
	}()

	return ch
}

func main() {

	ch := generate()

	fmt.Println(<-ch)
}
Output:

100
Why Pipelines Usually Return <-chan
Go pipeline pattern:

func square(in <-chan int) <-chan int
Example:

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for v := range in {
			out <- v * v
		}
	}()

	return out
}
Usage:

nums := generate()
sq := square(nums)

for v := range sq {
	fmt.Println(v)
}
This prevents callers from accidentally:

sq <- 10
close(sq)
which would break the pipeline.

Interview Answer
func Worker(ch1, ch2 <-chan int, out chan<- int)
Caller creates the channel.

Worker only gets send permission.

Caller can still send, receive, and close the channel.

func Worker(ch1, ch2 <-chan int) <-chan int
Worker creates the channel.

Caller only gets receive permission.

Caller cannot send or close it.

A useful rule:

Parameter direction restricts the callee.
Return direction restricts the caller.
That's the core difference.


*/
