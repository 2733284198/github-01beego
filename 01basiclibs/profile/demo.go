package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	//"github.com/EDDYCJY/go-pprof-example/data"
)

var datas []string
var data []string

func Add(str string) string {
	data := []byte(str)
	sData := string(data)
	datas = append(datas, sData)

	return sData
}

func main() {
	go func() {
		for {
			//log.Println(data.Add("https://github.com/EDDYCJY"))
			log.Println(("https://www.baidu.com/"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
