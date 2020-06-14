package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	// here are some values passed to channel
	c <- 4
	c <- 10
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case other := <-c:
			fmt.Println(other, "some")
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
