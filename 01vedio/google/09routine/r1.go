package main

import (
	"fmt"
	"time"
)

func main2() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println(i, "hello go")
			}
		}()

	}
	//time.Sleep(time.Millisecond)
	//Example01()
}

func Example01() {
	/*
		10 hello go
		10 hello go
		3 hello go
		10 hello go
		10 hello go
		10 hello go
		9 hello go
		10 hello go
		10 hello go
		10 hello go

	*/
}

func main() {
	//f, _ := os.Create("t3.out")
	//_ = trace.Start(f)

	//trace.Start(os.Stderr)
	//defer trace.Stop()

	for i := 0; i < 100000; i++ {
		//func(i int) {
		go func(i int) {
			for {
				fmt.Println(i, "hello go")
			}
		}(i)
	}
	time.Sleep(time.Millisecond)

	//trace.Stop()
}

func Example02() {
	/*
			0 hello go
		2 hello go
		4 hello go
		9 hello go
		5 hello go
		6 hello go
		3 hello go
		1 hello go
		8 hello go
		7 hello go
	*/
}

func TestMain() {
	fmt.Println("TestMain")
}
