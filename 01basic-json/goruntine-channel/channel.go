package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		//fmt.Println("Worker %d received %d\n", <- c)
		fmt.Println("Worker received %d\n", <-c)
	}
}

func chanDemo() {
	c := make(chan int)

	go func() {
		for {
			n := <-c

			println(n)
		}
	}()

	//go worker(c)

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()

}
