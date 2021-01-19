package main

import (
	"fmt"
	"sync"
)

var a int
var b chan int
var wg sync.WaitGroup

func nobufChannel() {

	b = make(chan int, 1)
	//b = make(chan int, 2)

	b <- 10
	fmt.Println("10发到通道b", b)
	// 一个通道，没有取出，会有问题。
	//b <- 20

	x := <-b
	fmt.Println("x取到值", x)

}

func nobufchannel() {

	//fmt.Println(b)
	b = make(chan int, 1)
	//b = make(chan int, 2)
	//b <- 10

	wg.Add(1)
	go func() {
		defer wg.Done()

		x := <-b
		// 1. fmt.println：2个谁先跑完。先后顺序不一定的，
		fmt.Println("后台取到了通道b", x)
	}()

	b <- 10
	// 1. fmt.println：2个谁先跑完
	fmt.Println("10发到后面通道b")
	//b = make(chan int, 16)

	fmt.Println(b)
	wg.Wait()
}

func main() {
	//nobuf()
	nobufChannel()
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
