/*This time the main function only consumes part of data in the Channel
(i.e. the for loop in main() only has 4 steps whereas the for loop in countCat() has 5 steps).
Can you spot the potential problems in this script?

Yes, you are right, the countCat GoRoutine actually has been blocked forever because
 the last "Cat" sent into the channel is not consumed by any other GoRoutines. The reason we do not
 see deadlock here is because the main() GoRoutine exits immediately after printing four "Cat".
 Let's make this example more clear with the help of WaitGroup.
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	go countCat(c)

	for i := 0; i < 4; i++ {
		msg, _ := <-c

		fmt.Println("Message to print ", msg)
	}

}

func countCat(c chan string) {
	for i := 0; i < 5; i++ {
		c <- "Cat"
	}

}
