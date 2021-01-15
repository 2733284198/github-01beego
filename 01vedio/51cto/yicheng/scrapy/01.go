package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	learn1()
	printbaidu()
	getFile()
}

func getFile() {
	resp, err := http.Get("http://localhost:9091/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func printbaidu() {
	//response, err :=http.Get("http://www.baidu.com")
	//if err != nil {
	//	fmt.Println("错误:", err)
	//}

	//baidubody := response.Body
	//baidufile, _ := ioutil.ReadAll(response.Body)

	//reg := regexp.MustCompile(`baidu`)
	//fmt.Printf("%q\n", reg.FindAllString(baidufile, -1))

	//fmt.Println(string(baidufile))
}

func learn1() {
	text := `Hello 世界！123 Go.`

	// 查找连续的小写字母
	reg := regexp.MustCompile(`[a-z]+`)
	fmt.Printf("%q\n", reg.FindAllString(text, -1))
}
