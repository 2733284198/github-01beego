package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f1(i int) {
	defer wg.Done()

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
	fmt.Println(i, " hello")
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}

	//time.Sleep(time.Second)
	wg.Wait()
}
