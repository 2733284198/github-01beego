package main

import (
	//"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	//"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	//myhttp "github.com/micro/go-plugins/client/http"

	//"github.com/hashicorp/consul/command/services/register"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// 原始调用2
	//mySelector := selector.NewSelector(
	//		selector.Registry(consulReg),
	//		selector.SetStrategy(selector.RoundRobin),
	//	)

	//callAPI2(mySelector)

	// 原始调用1
	for {

		getService, err := consulReg.GetService("prodservice")
		if err != nil {
			fmt.Println("===>err: ", err)
		}

		// 服务：随机，轮询
		//next := selector.Random(getService)
		next := selector.RoundRobin(getService)

		node, err := next()

		// 错误处理

		// CallAPI
		callRes, err := callAPI(node.Address, "/prod", "GET")

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(callRes)

		//fmt.Println("===>nodeid: ", node.Id , " nodeAddr: ", node.Address)
		//fmt.Println("===> nodeAddr: ", node.Address, "nodeid: ", node.Id, " nodeName: ", node.Metadata)
		fmt.Println("===> nodeAddr: ", node.Address)

		time.Sleep(time.Second * 1)

	}
}

func callAPI(addr string, path string, method string) (string, error) {
	//http.n
	req, err := http.NewRequest(method, "http://"+addr+path, nil)

	clientres := http.DefaultClient
	res, err := clientres.Do(req)

	if err != nil {
		return "-->err:", err
	}

	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)

	return string(buf), nil
}

/*func callAPI2(s selector.Selector)  {

	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
		)

	req := myClient.NewRequest("prodservice", "/prod", map[string]string{})

	var rsp map[string]interface{}

	err := myClient.Call(context.Background(), req, &rsp )

	if err != nil {
		log.Fatal(err)
	}

	//log.Fatal(rsp["data"])
	fmt.Println(rsp)
}*/
