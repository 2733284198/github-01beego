package main

import (
	"context"
	"fmt"
	"log"

	//"net/http"
	//"app02/service"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	myhttp "github.com/micro/go-plugins/client/http"

	//"github.com/hashicorp/consul/command/services/register"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)

	// 原始调用2
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)

	callAPI2(mySelector)
}

func callAPI2(s selector.Selector) {

	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)

	req := myClient.NewRequest("prodservice", "/prod", map[string]string{})
	//req := myClient.NewRequest("prodservice", "/prod", map[string]interface{} )
	//req := myClient.NewRequest("prodservice", "/prod", service.NewProd(5))

	var rsp map[string]interface{}

	err := myClient.Call(context.Background(), req, &rsp)

	if err != nil {
		log.Fatal(err)
	}

	//log.Fatal(rsp["data"])
	//fmt.Println(rsp["status"])
	fmt.Println(rsp)
}
