package main

import (
	"fmt"

	"github.com/salviati/go-qt5/qt5"
)

func main() {
	fmt.Println(qt5.Version())

	/*qt5.Main(func() {
		w := qt5.NewWidget()
		w.SetWindowTitle(qt5.Version())
		w.SetSizev(300, 200)
		defer w.Close()
		w.Show()
		qt5.Run()
	})*/
}
