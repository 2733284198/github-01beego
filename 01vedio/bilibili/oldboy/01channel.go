package main

import (
	"fmt"
	"sync"
)

//var a []int
var b chan int
var wg sync.WaitGroup

func main() {
	//fmt.Println(b)

	//b = make(chan int)
	b = make(chan int, 2)
	//b <- 10
	wg.Add(1)
	go func() {
		defer wg.Done()

		x := <-b
		fmt.Println("后台取到了通道b", x)
	}()

	b <- 10
	fmt.Println("10发到后面通道b")
	//b = make(chan int, 16)

	fmt.Println(b)
	wg.Wait()
	//fmt.Println(b)
}

func goprint() {

	outstr := `
		b = make(chan int)

		后台取到了通道b 10
		10发到后面通道b
		0xc000060060
	`

	/*outstr := `
	b = make(chan int, 2)

	10发到后面通道b
	0xc0000b8000
	后台取到了通道b 10

	10发到后面通道b
	0xc000052070

	`*/

	/*outstr := `
		后台取到了通道b 10
		10发到后面通道b
		0xc000060060
		panic: sync: negative WaitGroup counter
	`*/

	fmt.Println(outstr)
}
