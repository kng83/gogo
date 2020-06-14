package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			//	x, y = y, x+y
			temp := y //to jest to samo co wyzej
			y = x + y
			x = temp
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// ta funkcja odbiera dane z kanalu
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			//fmt.Println(i, "some")
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
