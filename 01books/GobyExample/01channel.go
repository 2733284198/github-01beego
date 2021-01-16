package main

import "fmt"

func main() {
	//messages := make(chan string)
	messages := make(chan string, 2)

	messages <- "buffer1"
	messages <- "buffer2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
