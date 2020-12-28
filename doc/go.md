
# 1 php-二位数组

[url] https://www.cnblogs.com/dingxiaoyue/p/4926793.html

```api

package main

import (
	"fmt"
)

func main() {
	/**
	 * 关键点：interface{} 可以代表任意类型
	 * 原理知识点:interface{} 就是一个空接口，所有类型都实现了这个接口，所以它可以代表所有类型
	 */
	var tags map[string]interface{}
	tags = make(map[string]interface{})

	var tagsLocal map[string]string
	tagsLocal = make(map[string]string)
	var tagsTest map[string]string
	tagsTest = make(map[string]string)
	var tagsProduction map[string]string
	tagsProduction = make(map[string]string)

	tagsLocal["dev.www.9178.us"] = "dev.www.9178.us"
	tagsLocal["dev.static.9178.us"] = "dev.static.9178.us"

	tagsTest["dev.www.9178.us"] = "www.ninja911.com"
	tagsTest["dev.static.9178.us"] = "static.ninja911.com"

	tagsProduction["dev.www.9178.us"] = "ipx.www.ninja911.com"
	tagsProduction["dev.static.9178.us"] = "ipx.static.ninja911.com"

	tags["local"] = tagsLocal
	tags["test"] = tagsLocal
	tags["production"] = tagsLocal

	fmt.Println(tags)
}


```
[url] http://blog.ninja911.com/blog-show-blog_id-76.html
