package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	//go fmt.Println(<-ch1)

	go func() {
		fmt.Println(<-ch1)
	}()

	ch1 <- 5

	//fmt.Println(<-ch1)

	time.Sleep(1 * time.Second)
}
