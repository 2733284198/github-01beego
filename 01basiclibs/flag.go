package main

import (
	"flag"
	"fmt"
)

//os.Args demo
func main() {
	//os.Args是一个[]string
	//if len(os.Args) > 0 {
	//	for index, arg := range os.Args {
	//		fmt.Printf("args[%d]=%v\n", index, arg)
	//	}
	//}

	//for index,arg := range os.Args {
	//	fmt.Printf("%d %v", index, arg)
	//}

	var name string

	flag.StringVar(&name, "name", "张三", "姓名")

	flag.Parse()
	//fmt.Println(name, age, married, delay)
	fmt.Println(name)
	//返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	//返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
