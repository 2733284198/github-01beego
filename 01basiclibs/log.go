package main

import (
	"log"
	"os"
)

func main() {
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)

	logFile, _ := os.OpenFile("./01basiclibs/logtest.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	log.SetOutput(logFile)

	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)

	//  log
	log.Println("这是我的日志1")

	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")

}
