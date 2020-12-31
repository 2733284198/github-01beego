/*package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	_"runtime/pprof"

	//"runtime/pprof"
	//"runtime/pprof"
	_ "net/http/pprof"
	//_ "runtime/pprof"
)

func main() {
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	http.ListenAndServe("localhost:6060", nil)

	//go func() {
	//	http.ListenAndServe("localhost:6060", nil)
	//}()

	logFile, _ := os.OpenFile("./01basiclibs/logtest.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)

	//  log
	log.Println("这是我的日志1")

	//log.Fatalln("这是一条会触发fatal的日志。")
	//log.Panicln("这是一条会触发panic的日志。")

	//pprof.StartCPUProfile(logFile)
	//defer pprof.StopCPUProfile()


}*/

package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	fmt.Fprintf(w, "Hello World")

	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal(err)
	}
}

/*

1. 运行
http://127.0.0.1:6060/debug/pprof/

2. 保存》profile后，
 go tool pprof log log.profile

3.

*/
