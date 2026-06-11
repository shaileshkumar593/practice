// Properly finish the time.Ticker goroutine
/*
When you run the time.Ticker loop in a goroutine, and at some point, you want to stop processing, you need to perform two steps:

stop the Ticker
ensure that the goroutine in which the Ticker was operating terminated correctly
To stop the Ticker, you need to run the Stop() method. However, the Stop() does not close the C channel.
It only prevents new ticks from being sent. Thus, it does not automatically exit the loop and the goroutine.
To close the goroutine, you need an additional signal that, when read, will cause it to close. So to achieve all this,
we will use a for loop with a select statement. The select blocks and waits until one of its cases can run. In our situation,
it should wait either for the value of the new tick from the C channel or for the signal to terminate the goroutine.

Look at the main() function in the example. After 5 seconds of goroutine processing, we stop the Ticker using the Stop()
method and send a new signal to the done channel. As a result, the Ticker stops receiving new ticks in the C channel,
and the select statement reads the value from the done, which causes the goroutine to return. In the first
5 seconds of the app, there are only messages in the ticker.C channel, so the select always chooses the case with a read from this channel.*/

package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Println("Hey!")
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- struct{}{}
	log.Println("Done")
}
