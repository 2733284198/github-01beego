# 百度搜索

 - golang regexp 正则表达式
 - regexp.MatchString()
 - reg := regexp.MustCompile(`[a-z]+`)
   	fmt.Printf("%q\n", reg.FindAllString(text, -1))
   	
# 读取html内容

```golang

> https://pkg.go.dev/net/http

resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...

```

 
 
