package caichao

import (
	"fmt"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			//fmt.Println("hello - ", j)
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond * 1)
	//t.Fail()
}
