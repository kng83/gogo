package main

import "fmt"

func sum(s []int, channel chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	channel <- sum // send sum to c
}

func main() {
	s := []int{70, 25, 80, -9, 4, 10, 20, 23, -300, 57, 23, -39, 23, 20, 3, -232, 32, 32, 3, 23, 23, 24, 23, 23, 23, 23, 20, 33, -100, 50}

	channel := make(chan int)
	//channel2 := make(chan int)
	go sum(s[:len(s)/2], channel)
	go sum(s[len(s)/2:], channel)
	x, y := <-channel, <-channel // receive from c (stos kanalow najpierw ostatni a na koncu pierwszy)

	fmt.Println(x, y, x+y)
}
