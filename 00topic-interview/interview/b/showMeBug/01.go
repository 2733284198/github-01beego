package main

import "fmt"

func main() {
	var a uint = 1
	var b uint = 2

	fmt.Println("a - b:", a-b)
	//fmt.Println(b - a)

	const MaxUint = ^uint(0)

	fmt.Println("MaxUint:", MaxUint)

	const MaxUint32 = 1<<32 - 1
	fmt.Println("MaxUint:", MaxUint)
}
