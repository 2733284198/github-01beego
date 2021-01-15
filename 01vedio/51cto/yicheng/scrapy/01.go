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

	// 读取全部数据，到 []bytes，转换为 string(bytes)
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	// 查找字符串
	html := string(body)

	//reg := regexp.MustCompile(`qrcode`)

	//compilestr := `[a-z]+`
	//compilestr := `\Qqrcode\E`
	//compilestr := `(?i:^qr).*Go`
	//compilestr := `^qr*`

	//reg := regexp.MustCompile(compilestr)

	//reg := regexp.MustCompile(`\Qqrcode\E`)
	//reg = regexp.MustCompile(`(?i:^hello).*Go`)
	// 查找 Hello 或 Go，替换为 Hellooo、Gooo
	//reg := regexp.MustCompile(`\w+`)

	/*
		<ul>
		        <li>qrcode1</li>
		        <li>qrcode2</li>
		        <li>qrcode3</li>
		        <li>qrcode4</li>
		        <li>qrcode5</li>

		        <li>qrcode6a</li>
		    </ul>

	*/
	//reg := regexp.MustCompile(`qr[a-z]+[1-9][a-z]*`)
	//reg := regexp.MustCompile(`qr[a-z]+[1-9][a-z]*`)
	//reg := regexp.MustCompile(`[qr]*[a-z]+[1-9][a-z]*`)
	// 手机号码
	//reg := regexp.MustCompile(`1[3-9]\d{9}`)

	/*
		<ul>
		        <li> --- 10 位</li>
		        <li>13104884551</li>
		        <li>18104884552</li>

		        <li> --- 11 位</li>
		        <li>131048845532</li>
		    </ul>
	*/

	// 12的，手机号，没有处理
	reg := regexp.MustCompile(`1[3456789]\d{9}`)
	//reg := regexp.MustCompile(`qr[a-z]+[1-9][a-z]*`)

	//reg = regexp.MustCompile(`\Qqrcode\E`)
	//reg := regexp.MustCompile(`小米`)

	fmt.Printf("%q\n", reg.FindAllString(html, -1))
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
