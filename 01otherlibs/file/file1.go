package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./01otherlibs/file/file1.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
	}

	//file.WriteString("ab\n")
	//file.WriteString("ab2\n")

	defer file.Close()
}
