package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f1(i int) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	fmt.Println(i, " hello")
}

func main() {
	for i := 0; i < 10; i++ {
		go f1(i)
	}

	time.Sleep(time.Second)
}
