package main

import (
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()

	greeter := proto.NewGreeterService("greeter.service", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Micro中国"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
